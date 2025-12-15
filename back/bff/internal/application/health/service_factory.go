package health

import "fmt"

func NewService(checker UpstreamChecker) (Service, error) {
	if checker == nil {
		return Service{}, fmt.Errorf("upstream checker is required")
	}

	return Service{checker: checker}, nil
}
