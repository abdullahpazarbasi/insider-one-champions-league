package status

type StatusResponse struct {
	Service  string          `json:"service"`
	Status   string          `json:"status"`
	Upstream *UpstreamStatus `json:"upstream,omitempty"`
}

type UpstreamStatus struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}
