package rpc

import (
	"net/http"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
)

type customCodec struct {
}

func newCustomCodec() *customCodec {
	return &customCodec{}
}

type customCodecRequest struct {
	*json2.CodecRequest
}

func (cr *customCodecRequest) Method() (string, error) {
	m, err := cr.CodecRequest.Method()
	if len(m) == 0 || err != nil {
		return m, err
	}

	ar := strings.SplitN(m, "_", 2)
	if len(ar) < 2 || len(ar[0]) == 0 || len(ar[1]) == 0 {
		return m, nil
	}

	svc, method := ar[0], ar[1]
	r, n := utf8.DecodeRuneInString(method)
	if !unicode.IsLower(r) {
		return m, nil
	}

	return svc + "." + string(unicode.ToUpper(r)) + method[n:], nil
}

func (c *customCodec) NewRequest(r *http.Request) rpc.CodecRequest {
	return &customCodecRequest{
		CodecRequest: json2.NewCodec().NewRequest(r).(*json2.CodecRequest),
	}
}
