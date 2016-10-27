package cabot

import (
	"fmt"
	"net/http"
)

const (
	InstancesEndpoint = "/api/instances/"
)

type InstancesService service

type Instance struct {
	Address       string `json:"address,omitempty"`
	Alerts        []int  `json:"alerts,omitempty"`
	AlertsEnabled bool   `json:"alerts_enabled,omitempty"`
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	OverallStatus string `json:"overall_status,omitempty"`
	StatusChecks  []int  `json:"status_checks,omitempty"`
	UsersToNotify []int  `json:"users_to_notify,omitempty"`
}

func (s *InstancesService) List() ([]*Instance, error) {
	req, err := s.client.NewRequest("GET", InstancesEndpoint, nil)
	if err != nil {
		return nil, err
	}

	instances := new([]*Instance)
	err = s.client.Do(req, instances)
	if err != nil {
		return nil, err
	}

	return *instances, nil
}

func (s *InstancesService) doSingleInstance(req *http.Request) (*Instance, error) {
	ins := new(Instance)
	err := s.client.Do(req, ins)
	if err != nil {
		return nil, err
	}
	return ins, nil
}

func (s *InstancesService) Get(id int) (*Instance, error) {
	u := fmt.Sprintf("%v%v/", InstancesEndpoint, id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	return s.doSingleInstance(req)
}

func (s *InstancesService) Create(instance *Instance) (*Instance, error) {
	req, err := s.client.NewRequest("POST", InstancesEndpoint, instance)
	if err != nil {
		return nil, err
	}
	return s.doSingleInstance(req)
}

func (s *InstancesService) Update(id int, instance *Instance) (*Instance, error) {
	u := fmt.Sprintf("%v%v/", InstancesEndpoint, id)
	req, err := s.client.NewRequest("PATCH", u, instance)
	if err != nil {
		return nil, err
	}
	return s.doSingleInstance(req)
}

func (s *InstancesService) Delete(id int) error {
	u := fmt.Sprintf("%v%v/", InstancesEndpoint, id)
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
