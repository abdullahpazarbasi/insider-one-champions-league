package proxy

import (
	"context"
	"net/http"
)

func (s Service) Get(ctx context.Context, target string, headers http.Header) (*http.Response, error) {
	return s.forwarder.ForwardGet(ctx, target, headers)
}

func (s Service) Post(ctx context.Context, target string, body []byte, headers http.Header) (*http.Response, error) {
	return s.forwarder.ForwardPost(ctx, target, body, headers)
}
