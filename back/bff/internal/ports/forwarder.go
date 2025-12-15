package ports

import (
	"context"
	"net/http"
)

type Forwarder interface {
	ForwardGet(ctx context.Context, target string, headers http.Header) (*http.Response, error)
	ForwardPost(ctx context.Context, target string, body []byte, headers http.Header) (*http.Response, error)
}
