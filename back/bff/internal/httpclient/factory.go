package httpclient

import (
	"net/http"
	"time"
)

func New(timeout time.Duration) *http.Client {
	effectiveTimeout := timeout
	if effectiveTimeout <= 0 {
		effectiveTimeout = 2 * time.Second
	}

	return &http.Client{Timeout: effectiveTimeout}
}
