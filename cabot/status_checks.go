package cabot

const (
	StatusChecksEndpoint = "/api/status_checks/"
)

type StatusChecksService service

type StatusCheck struct {
	Active           bool       `json:"active,omitempty"`
	CalculatedStatus Status     `json:"calculated_status,omitempty"`
	Debounce         int        `json:"debounce,omitempty"`
	Frequency        int        `json:"frequency,omitempty"`
	ID               int        `json:"id,omitempty"`
	Importance       Importance `json:"importance,omitempty"`
	Name             string     `json:"name,omitempty"`
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
