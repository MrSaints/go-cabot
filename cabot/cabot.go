package cabot

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "1.0.0"
	userAgent      = "go-cabot/" + libraryVersion
	defaultBaseURL = "http://localhost:5001/"
)

type Client struct {
	client    *http.Client
	basicAuth *BasicAuth
	BaseURL   *url.URL
	UserAgent string

	common         service
	Plugins        *PluginsService
	Shifts         *ShiftsService
	Services       *ServicesService
	Instances      *InstancesService
	StatusChecks   *StatusChecksService
	ICMPChecks     *ICMPChecksService
	HTTPChecks     *HTTPChecksService
	JenkinsChecks  *JenkinsChecksService
	GraphiteChecks *GraphiteChecksService
}

type BasicAuth struct {
	username, password string
}

type service struct {
	client *Client
}

type Option func(*Client) error

func NewClient(opts ...Option) (*Client, error) {
	baseURL, err := url.Parse(defaultBaseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client:    http.DefaultClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
	}
	c.common.client = c
	c.Plugins = (*PluginsService)(&c.common)
	c.Shifts = (*ShiftsService)(&c.common)
	c.Services = (*ServicesService)(&c.common)
	c.Instances = (*InstancesService)(&c.common)
	c.StatusChecks = (*StatusChecksService)(&c.common)
	c.ICMPChecks = (*ICMPChecksService)(&c.common)
	c.HTTPChecks = (*HTTPChecksService)(&c.common)
	c.JenkinsChecks = (*JenkinsChecksService)(&c.common)
	c.GraphiteChecks = (*GraphiteChecksService)(&c.common)

	for _, opt := range opts {
		err = opt(c)
		if err != nil {
			return nil, errors.Wrap(err, "set option failed")
		}
	}

	return c, nil
}

func WithBasicAuth(u string, p string) Option {
	return func(c *Client) error {
		if c.basicAuth == nil {
			c.basicAuth = new(BasicAuth)
		}
		c.basicAuth.username = u
		c.basicAuth.password = p
		return nil
	}
}

func WithBaseURL(u string) Option {
	return func(c *Client) error {
		baseURL, err := url.Parse(u)
		if err != nil {
			return err
		}
		c.BaseURL = baseURL
		return nil
	}
}

func (c *Client) NewRequest(method, api string, body interface{}) (*http.Request, error) {
	ref, err := url.Parse(api)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(ref)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.basicAuth != nil {
		req.SetBasicAuth(c.basicAuth.username, c.basicAuth.password)
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) error {
	res, err := c.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "api request failed")
	}

	if res.Body != nil {
		defer func() {
			io.CopyN(ioutil.Discard, res.Body, 512)
			res.Body.Close()
		}()
	}

	if s := res.StatusCode; s < 200 || s > 299 {
		raw, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return errors.Errorf("api did not return a http success code: %s ( %s )", res.Status, string(raw))
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, res.Body)
		} else {
			err = json.NewDecoder(res.Body).Decode(v)
			if err == io.EOF {
				err = nil
			}
		}
	}

	return nil
}
