package cabot

import (
	"github.com/pkg/errors"
	"strings"
)

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
//go:generate stringer -type=Importance
type Importance byte

const (
	_ Importance = iota
	WARNING
	ERROR
	CRITICAL
)

func ImportanceStringToConst(s string) (Importance, error) {
	if v, ok := _ImportanceNameToValue[s]; ok {
		return v, nil
	}

	keys := make([]string, 0, len(_ImportanceNameToValue))
	for k := range _ImportanceNameToValue {
		keys = append(keys, k)
	}

	return 0, errors.Errorf(
		"invalid importance: %s (available options: %s)",
		s,
		strings.Join(keys, ", "),
	)
}

//go:generate jsonenums -type=Status
//go:generate stringer -type=Status
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
