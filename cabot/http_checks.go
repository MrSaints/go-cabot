package cabot

import (
	"fmt"
	"net/http"
)

const (
	HTTPChecksEndpoint = "/api/http_checks/"
)

type HTTPChecksService service

type HTTPCheck struct {
	StatusCheck

	Endpoint   string `json:"endpoint,omitempty"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	TextMatch  string `json:"text_match,omitempty"`
	StatusCode string `json:"status_code,omitempty"`
	Timeout    int    `json:"timeout,omitempty"`
	VerifySSL  bool   `json:"verify_ssl_certificate,omitempty"`
}

func (s *HTTPChecksService) List() ([]*HTTPCheck, error) {
	req, err := s.client.NewRequest("GET", HTTPChecksEndpoint, nil)
	if err != nil {
		return nil, err
	}

	checks := new([]*HTTPCheck)
	err = s.client.Do(req, checks)
	if err != nil {
		return nil, err
	}

	return *checks, nil
}

func (s *HTTPChecksService) doSingleHTTPCheck(req *http.Request) (*HTTPCheck, error) {
	c := new(HTTPCheck)
	err := s.client.Do(req, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *HTTPChecksService) Get(id int) (*HTTPCheck, error) {
	u := fmt.Sprintf("%v%v/", HTTPChecksEndpoint, id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	return s.doSingleHTTPCheck(req)
}

func (s *HTTPChecksService) Create(check *HTTPCheck) (*HTTPCheck, error) {
	req, err := s.client.NewRequest("POST", HTTPChecksEndpoint, check)
	if err != nil {
		return nil, err
	}
	return s.doSingleHTTPCheck(req)
}

func (s *HTTPChecksService) Update(id int, check *HTTPCheck) (*HTTPCheck, error) {
	u := fmt.Sprintf("%v%v/", HTTPChecksEndpoint, id)
	req, err := s.client.NewRequest("PATCH", u, check)
	if err != nil {
		return nil, err
	}
	return s.doSingleHTTPCheck(req)
}

func (s *HTTPChecksService) Delete(id int) error {
	u := fmt.Sprintf("%v%v/", HTTPChecksEndpoint, id)
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
