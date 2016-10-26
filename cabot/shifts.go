package cabot

import (
	"time"
)

const (
	ShiftsEndpoint = "/api/shifts/"
)

type ShiftsService service

type Shift struct {
	Deleted bool      `json:"deleted,omitempty"`
	End     time.Time `json:"end,omitempty"`
	ID      int       `json:"id,omitempty"`
	Start   time.Time `json:"start,omitempty"`
	UID     string    `json:"uid,omitempty"`
	User    int       `json:"user,omitempty"`
}

func (s *ShiftsService) List() ([]*Shift, error) {
	req, err := s.client.NewRequest("GET", ShiftsEndpoint, nil)
	if err != nil {
		return nil, err
	}

	shifts := new([]*Shift)
	err = s.client.Do(req, shifts)
	if err != nil {
		return nil, err
	}

	return *shifts, nil
}
