package cabot

import (
	"fmt"
	"net/http"
)

const (
	ServicesEndpoint = "/api/services/"
)

type ServicesService service

type Service struct {
	Alerts        []int  `json:"alerts"`
	AlertsEnabled bool   `json:"alerts_enabled"`
	ID            int    `json:"id"`
	Instances     []int  `json:"instances"`
	Name          string `json:"name"`
	OverallStatus string `json:"overall_status"`
	StatusChecks  []int  `json:"status_checks"`
	URL           string `json:"url"`
	UsersToNotify []int  `json:"users_to_notify"`
}

func (s *ServicesService) List() ([]*Service, error) {
	req, err := s.client.NewRequest("GET", ServicesEndpoint, nil)
	if err != nil {
		return nil, err
	}

	services := new([]*Service)
	err = s.client.Do(req, services)
	if err != nil {
		return nil, err
	}

	return *services, nil
}

func (s *ServicesService) doSingleService(req *http.Request) (*Service, error) {
	svc := new(Service)
	err := s.client.Do(req, svc)
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func (s *ServicesService) Get(id int) (*Service, error) {
	u := fmt.Sprintf("%v%v/", ServicesEndpoint, id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	return s.doSingleService(req)
}

func (s *ServicesService) Create(service *Service) (*Service, error) {
	req, err := s.client.NewRequest("POST", ServicesEndpoint, service)
	if err != nil {
		return nil, err
	}
	return s.doSingleService(req)
}

func (s *ServicesService) Delete(id int) error {
	u := fmt.Sprintf("%v%v/", ServicesEndpoint, id)
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
