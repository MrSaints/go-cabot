package cabot

import (
	"fmt"
	"net/http"
)

const (
	GraphiteChecksEndpoint = "/api/graphite_checks/"
)

type GraphiteChecksService service

type GraphiteCheck struct {
	StatusCheck

	Metric             string `json:"metric,omitempty"`
	CheckType          string `json:"check_type,omitempty"`
	Value              string `json:"value,omitempty"`
	ExpectedNumHosts   int    `json:"expected_num_hosts,omitempty"`
	AllowedNumFailures int    `json:"allowed_num_failures,omitempty"`
}

func (s *GraphiteChecksService) List() ([]*GraphiteCheck, error) {
	req, err := s.client.NewRequest("GET", GraphiteChecksEndpoint, nil)
	if err != nil {
		return nil, err
	}

	checks := new([]*GraphiteCheck)
	err = s.client.Do(req, checks)
	if err != nil {
		return nil, err
	}

	return *checks, nil
}

func (s *GraphiteChecksService) doSingleGraphiteCheck(req *http.Request) (*GraphiteCheck, error) {
	c := new(GraphiteCheck)
	err := s.client.Do(req, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *GraphiteChecksService) Get(id int) (*GraphiteCheck, error) {
	u := fmt.Sprintf("%v%v/", GraphiteChecksEndpoint, id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	return s.doSingleGraphiteCheck(req)
}

func (s *GraphiteChecksService) Create(check *GraphiteCheck) (*GraphiteCheck, error) {
	req, err := s.client.NewRequest("POST", GraphiteChecksEndpoint, check)
	if err != nil {
		return nil, err
	}
	return s.doSingleGraphiteCheck(req)
}

func (s *GraphiteChecksService) Update(id int, check *GraphiteCheck) (*GraphiteCheck, error) {
	u := fmt.Sprintf("%v%v/", GraphiteChecksEndpoint, id)
	req, err := s.client.NewRequest("PATCH", u, check)
	if err != nil {
		return nil, err
	}
	return s.doSingleGraphiteCheck(req)
}

func (s *GraphiteChecksService) Delete(id int) error {
	u := fmt.Sprintf("%v%v/", GraphiteChecksEndpoint, id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}
	err = s.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}
