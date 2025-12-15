package domain

type StatusResponse struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}
