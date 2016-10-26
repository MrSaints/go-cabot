package cabot

const (
	PluginsEndpoint = "/api/alertplugins/"
)

type PluginsService service

type Plugin struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

func (s *PluginsService) List() ([]*Plugin, error) {
	req, err := s.client.NewRequest("GET", PluginsEndpoint, nil)
	if err != nil {
		return nil, err
	}

	plugins := new([]*Plugin)
	err = s.client.Do(req, plugins)
	if err != nil {
		return nil, err
	}

	return *plugins, nil
}
