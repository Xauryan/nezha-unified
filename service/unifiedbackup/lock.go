package unifiedbackup

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/gofrs/flock"
)

// LockDataRoot is shared by Dashboard and offline maintenance commands. The
// lock file is never unlinked: removing it would allow two live lock inodes.
func LockDataRoot(rootPath string) (func() error, error) {
	if err := os.MkdirAll(rootPath, 0700); err != nil {
		return nil, err
	}
	if err := rejectSymlinks(rootPath); err != nil {
		return nil, err
	}
	lockPath := filepath.Join(rootPath, ".unified.lock")
	if info, err := os.Lstat(lockPath); err == nil && !info.Mode().IsRegular() {
		return nil, errors.New("data lock must be a regular file")
	} else if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	lock := flock.New(lockPath)
	ok, err := lock.TryLock()
	if err != nil {
		return nil, errors.New("cannot acquire Dashboard data lock")
	}
	if !ok {
		_ = lock.Close()
		return nil, errors.New("Dashboard data is in use; stop Dashboard before offline maintenance")
	}
	return lock.Close, nil
}
