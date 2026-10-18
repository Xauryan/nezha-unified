package unifiedbackup

import (
	"crypto/cipher"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"io"
)

const streamMagic = "NEZHA-BACKUP-V1\n"
const chunkSize = 64 << 10

// Every fixed-size record is bound to its sequence and the complete header. A
// final authenticated empty record detects truncation, including at boundaries.
type encryptWriter struct {
	w          io.Writer
	aead       cipher.AEAD
	headerHash [32]byte
	sequence   uint64
	buffer     []byte
	closed     bool
}

func newEncryptWriter(w io.Writer, recoveryKey, key []byte, id string) (*encryptWriter, error) {
	header := []byte(streamMagic + id)
	wrapped, err := seal(recoveryKey, key, header)
	if err != nil {
		return nil, err
	}
	header = append(header, wrapped...)
	if _, err := w.Write(header); err != nil {
		return nil, err
	}
	aead, err := newAEAD(key)
	if err != nil {
		return nil, err
	}
	return &encryptWriter{w: w, aead: aead, headerHash: sha256.Sum256(header), buffer: make([]byte, 0, chunkSize)}, nil
}

func (w *encryptWriter) Write(p []byte) (int, error) {
	if w.closed {
		return 0, errors.New("backup encryption stream closed")
	}
	n := 0
	for len(p) > 0 {
		copied := min(len(p), chunkSize-len(w.buffer))
		w.buffer = append(w.buffer, p[:copied]...)
		p, n = p[copied:], n+copied
		if len(w.buffer) == chunkSize {
			if err := w.record(w.buffer, false); err != nil {
				return n, err
			}
			clear(w.buffer)
			w.buffer = w.buffer[:0]
		}
	}
	return n, nil
}
func (w *encryptWriter) record(data []byte, final bool) error {
	nonce, aad := recordContext(w.headerHash, w.sequence, final)
	sealed := w.aead.Seal(nil, nonce, data, aad)
	var prefix [5]byte
	binary.BigEndian.PutUint32(prefix[:4], uint32(len(sealed)))
	if final {
		prefix[4] = 1
	}
	if _, err := w.w.Write(prefix[:]); err != nil {
		return err
	}
	if _, err := w.w.Write(sealed); err != nil {
		return err
	}
	w.sequence++
	return nil
}
func (w *encryptWriter) Close() error {
	if w.closed {
		return errors.New("backup encryption stream already closed")
	}
	w.closed = true
	if len(w.buffer) > 0 {
		if err := w.record(w.buffer, false); err != nil {
			return err
		}
	}
	clear(w.buffer)
	return w.record(nil, true)
}

type decryptReader struct {
	r          io.Reader
	aead       cipher.AEAD
	headerHash [32]byte
	sequence   uint64
	buffer     []byte
	final      bool
}

func newDecryptReader(r io.Reader, recoveryKey []byte) (*decryptReader, []byte, string, error) {
	header := make([]byte, len(streamMagic)+32+12+32+16)
	if _, err := io.ReadFull(r, header); err != nil {
		return nil, nil, "", errors.New("invalid encrypted backup header")
	}
	if string(header[:len(streamMagic)]) != streamMagic {
		return nil, nil, "", errors.New("unsupported backup format")
	}
	prefix := header[:len(streamMagic)+32]
	key, err := unseal(recoveryKey, header[len(prefix):], prefix)
	if err != nil || len(key) != 32 {
		return nil, nil, "", errors.New("cannot authenticate backup recovery key")
	}
	aead, err := newAEAD(key)
	if err != nil {
		clear(key)
		return nil, nil, "", err
	}
	return &decryptReader{r: r, aead: aead, headerHash: sha256.Sum256(header)}, key, string(prefix[len(streamMagic):]), nil
}
func (r *decryptReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	if len(r.buffer) == 0 {
		if r.final {
			return 0, io.EOF
		}
		var prefix [5]byte
		if _, err := io.ReadFull(r.r, prefix[:]); err != nil {
			return 0, errors.New("encrypted backup is truncated")
		}
		size := binary.BigEndian.Uint32(prefix[:4])
		final := prefix[4] == 1
		if prefix[4] > 1 || size < uint32(r.aead.Overhead()) || size > chunkSize+uint32(r.aead.Overhead()) || (final && size != uint32(r.aead.Overhead())) || (!final && size == uint32(r.aead.Overhead())) {
			return 0, errors.New("invalid encrypted backup record")
		}
		sealed := make([]byte, size)
		if _, err := io.ReadFull(r.r, sealed); err != nil {
			return 0, errors.New("encrypted backup is truncated")
		}
		nonce, aad := recordContext(r.headerHash, r.sequence, final)
		data, err := r.aead.Open(nil, nonce, sealed, aad)
		if err != nil {
			return 0, errors.New("backup record authentication failed")
		}
		r.sequence++
		if final {
			var probe [1]byte
			if n, err := r.r.Read(probe[:]); n != 0 || err != io.EOF {
				return 0, errors.New("backup contains unauthenticated trailing data")
			}
			r.final = true
			return 0, io.EOF
		}
		r.buffer = data
	}
	n := copy(p, r.buffer)
	clear(r.buffer[:n])
	r.buffer = r.buffer[n:]
	return n, nil
}

func recordContext(headerHash [32]byte, sequence uint64, final bool) ([]byte, []byte) {
	nonce := make([]byte, 12)
	binary.BigEndian.PutUint64(nonce[4:], sequence)
	aad := make([]byte, 41)
	copy(aad, headerHash[:])
	binary.BigEndian.PutUint64(aad[32:40], sequence)
	if final {
		aad[40] = 1
	}
	return nonce, aad
}
