package cabot

const (
	StatusChecksEndpoint = "/api/status_checks/"
)

type StatusChecksService service

type StatusCheck struct {
	Active           bool       `json:"active"`
	CalculatedStatus Status     `json:"calculated_status"`
	Debounce         int        `json:"debounce,omitempty"`
	Frequency        int        `json:"frequency"`
	ID               int        `json:"id"`
	Importance       Importance `json:"importance"`
	Name             string     `json:"name"`
}

//go:generate jsonenums -type=Importance
type Importance byte

const (
	WARNING Importance = iota
	ERROR
	CRITICAL
)

//go:generate jsonenums -type=Status
type Status byte

const (
	passing Status = iota
	intermittent
	failing
)

func (s *StatusChecksService) List() ([]*StatusCheck, error) {
	req, err := s.client.NewRequest("GET", StatusChecksEndpoint, nil)
	if err != nil {
		return nil, err
	}

	checks := new([]*StatusCheck)
	err = s.client.Do(req, checks)
	if err != nil {
		return nil, err
	}

	return *checks, nil
}
