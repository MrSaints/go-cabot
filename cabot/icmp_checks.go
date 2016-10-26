package cabot

import (
	"fmt"
	"net/http"
)

const (
	ICMPChecksEndpoint = "/api/icmp_checks/"
)

type ICMPChecksService service

type ICMPCheck struct {
	StatusCheck
}

func (s *ICMPChecksService) List() ([]*ICMPCheck, error) {
	req, err := s.client.NewRequest("GET", ICMPChecksEndpoint, nil)
	if err != nil {
		return nil, err
	}

	checks := new([]*ICMPCheck)
	err = s.client.Do(req, checks)
	if err != nil {
		return nil, err
	}

	return *checks, nil
}

func (s *ICMPChecksService) doSingleICMPCheck(req *http.Request) (*ICMPCheck, error) {
	c := new(ICMPCheck)
	err := s.client.Do(req, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *ICMPChecksService) Get(id int) (*ICMPCheck, error) {
	u := fmt.Sprintf("%v%v/", ICMPChecksEndpoint, id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	return s.doSingleICMPCheck(req)
}

func (s *ICMPChecksService) Create(check *ICMPCheck) (*ICMPCheck, error) {
	req, err := s.client.NewRequest("POST", ICMPChecksEndpoint, check)
	if err != nil {
		return nil, err
	}
	return s.doSingleICMPCheck(req)
}
