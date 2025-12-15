package status

import "testing"

func TestStatusResponseSerializationFields(t *testing.T) {
	resp := StatusResponse{Service: "bff", Status: "ok", Upstream: &UpstreamStatus{Service: "service", Status: "ok"}}
	if resp.Service != "bff" || resp.Upstream.Service != "service" {
		t.Fatalf("unexpected struct initialization: %+v", resp)
	}
}
