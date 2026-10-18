package sign

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"github.com/google/uuid"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/errs"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/request"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	AccessKey             = "AccessKey"
	TIMESTAMP             = "Timestamp"
	TimestampFormat       = "2006-01-02T15:04:05Z"
	Signature             = "Signature"
	SecretKeyPrefix       = "BC_SIGNATURE&"
	SignatureMethod       = "SignatureMethod"
	SignatureMethodValue  = "HmacSHA1"
	SignatureVersion      = "SignatureVersion"
	SignatureVersionValue = "V2.0"
	SignatureNonce        = "SignatureNonce"
	LineSeparator         = "\n"
	ParameterSeparator    = "&"
	QueryStartSymbol      = "?"
	QuerySeparator        = "="
)

const (
	HighMask = 0xf0
	LowMask  = 0x0f
)

var HexCodeTable = []string{
	"0", "1", "2", "3",
	"4", "5", "6", "7",
	"8", "9", "a", "b",
	"c", "d", "e", "f",
}

func Sign(httpRequest *request.HttpRequest, accessKey string, secretKey string) error {
	params := make(map[string]string)
	for key, value := range httpRequest.QueryParams {
		params[key] = value
	}
	params[AccessKey] = accessKey
	now := time.Now()
	params[TIMESTAMP] = now.Format(TimestampFormat)
	params[SignatureMethod] = SignatureMethodValue
	params[SignatureVersion] = SignatureVersionValue
	params[SignatureNonce] = nonce()
	keys := make([]string, len(params))
	index := 0
	for key := range params {
		keys[index] = key
		index++
	}
	sort.Strings(keys)
	builder := strings.Builder{}
	pos := 0
	paramsLen := len(keys)
	for _, key := range keys {
		value := params[key]
		builder.WriteString(utils.PercentEncode(key))
		builder.WriteString(QuerySeparator)
		builder.WriteString(utils.PercentEncode(value))
		if pos != paramsLen-1 {
			builder.WriteString(ParameterSeparator)
			pos++
		}
	}
	canonicalQueryString := builder.String()

	hashString := convertToHexString(sha256Encode(canonicalQueryString))

	unescapedPath, err := url.QueryUnescape(httpRequest.Path)
	if nil != err {
		return errs.NewSignatureError(err.Error(), err)
	}
	builder.Reset()
	builder.WriteString(strings.ToUpper(httpRequest.Method))
	builder.WriteString(LineSeparator)
	builder.WriteString(utils.PercentEncode(unescapedPath))
	builder.WriteString(LineSeparator)
	builder.WriteString(hashString)
	stringToSign := builder.String()

	signature := convertToHexString(hmacSha1(stringToSign, SecretKeyPrefix+secretKey))

	builder.Reset()
	builder.WriteString(unescapedPath)
	builder.WriteString(QueryStartSymbol)
	builder.WriteString(canonicalQueryString)
	builder.WriteString(ParameterSeparator)
	builder.WriteString(Signature)
	builder.WriteString(QuerySeparator)
	builder.WriteString(utils.PercentEncode(signature))
	httpRequest.Path = builder.String()
	return nil
}

func hmacSha1(text string, keyStr string) []byte {
	key := []byte(keyStr)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(text))
	return mac.Sum(nil)
}

func convertToHexString(data []byte) string {
	if data == nil {
		return ""
	}
	builder := strings.Builder{}
	for _, d := range data {
		builder.WriteString(HexCodeTable[(HighMask&d)>>4])
		builder.WriteString(HexCodeTable[LowMask&d])
	}
	return builder.String()
}

func sha256Encode(text string) []byte {
	h := sha256.New()
	h.Write([]byte(text))
	return h.Sum(nil)
}

func nonce() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
