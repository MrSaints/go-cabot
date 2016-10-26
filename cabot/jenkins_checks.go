package cabot

import (
	"fmt"
	"net/http"
)

const (
	JenkinsChecksEndpoint = "/api/jenkins_checks/"
)

type JenkinsChecksService service

type JenkinsCheck struct {
	StatusCheck

	MaxQueuedBuildTime int `json:"max_queued_build_time,omitempty"`
}

func (s *JenkinsChecksService) List() ([]*JenkinsCheck, error) {
	req, err := s.client.NewRequest("GET", JenkinsChecksEndpoint, nil)
	if err != nil {
		return nil, err
	}

	checks := new([]*JenkinsCheck)
	err = s.client.Do(req, checks)
	if err != nil {
		return nil, err
	}

	return *checks, nil
}

func (s *JenkinsChecksService) doSingleJenkinsCheck(req *http.Request) (*JenkinsCheck, error) {
	c := new(JenkinsCheck)
	err := s.client.Do(req, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *JenkinsChecksService) Get(id int) (*JenkinsCheck, error) {
	u := fmt.Sprintf("%v%v/", JenkinsChecksEndpoint, id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	return s.doSingleJenkinsCheck(req)
}

func (s *JenkinsChecksService) Create(check *JenkinsCheck) (*JenkinsCheck, error) {
	req, err := s.client.NewRequest("POST", JenkinsChecksEndpoint, check)
	if err != nil {
		return nil, err
	}
	return s.doSingleJenkinsCheck(req)
}

func (s *JenkinsChecksService) Edit(id int, check *JenkinsCheck) (*JenkinsCheck, error) {
	u := fmt.Sprintf("%v%v/", JenkinsChecksEndpoint, id)
	req, err := s.client.NewRequest("PATCH", u, check)
	if err != nil {
		return nil, err
	}
	return s.doSingleJenkinsCheck(req)
}

func (s *JenkinsChecksService) Delete(id int) error {
	u := fmt.Sprintf("%v%v/", JenkinsChecksEndpoint, id)
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
