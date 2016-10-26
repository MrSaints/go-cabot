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
	Alerts        []int  `json:"alerts,omitempty"`
	AlertsEnabled bool   `json:"alerts_enabled,omitempty"`
	ID            int    `json:"id,omitempty"`
	Instances     []int  `json:"instances,omitempty"`
	Name          string `json:"name,omitempty"`
	OverallStatus string `json:"overall_status,omitempty"`
	StatusChecks  []int  `json:"status_checks,omitempty"`
	URL           string `json:"url,omitempty"`
	UsersToNotify []int  `json:"users_to_notify,omitempty"`
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

func (s *ServicesService) Edit(id int, service *Service) (*Service, error) {
	u := fmt.Sprintf("%v%v/", ServicesEndpoint, id)
	req, err := s.client.NewRequest("PATCH", u, service)
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
