package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/pkg/errors"
)

type KeyValuePairs map[string]string

type Request struct {
	Methods    []string      `yaml:"methods"`
	Path       string        `yaml:"path"`
	PathPrefix string        `yaml:"path-prefix"`
	Query      KeyValuePairs `yaml:"query"`
	Headers    KeyValuePairs `yaml:"headers"`
}

func (req *Request) Validate() error {
	var err error
	if len(req.Methods) == 0 {
		err = errors.New("missing request methods")
	} else if req.Path == "" && req.PathPrefix == "" {
		err = errors.New("missing request path or path-prefix")
	}
	return err
}

func (req *Request) String() string {
	var route string
	if req.Path != "" {
		route = req.Path
	} else {
		route = req.PathPrefix
	}
	return fmt.Sprintf("%s %s", color.MagentaString(fmt.Sprintf("%v", req.Methods)), color.GreenString(route))
}

type Response struct {
	Headers KeyValuePairs `yaml:"headers"`
	Latency int           `yaml:"latency"`
	File    string        `yaml:"file"`
	Body    string        `yaml:"body"`
	Status  int           `yaml:"status"`
}

func (res *Response) Validate() error {
	var err error
	if res.Status == 0 {
		err = errors.New("missing response status")
	} else if res.File == "" && res.Body == "" {
		err = errors.New("missing response file or body")
	}
	return err
}

func (res *Response) String() string {
	return color.YellowString(fmt.Sprintf("%d", res.Status))
}

type Stub struct {
	Request  Request  `yaml:"request"`
	Response Response `yaml:"response"`
}

func (s *Stub) String() string {
	return fmt.Sprintf("%s %s %s", s.Request.String(), color.CyanString("->"), s.Response.String())
}

func (s *Stub) Validate() error {
	var err error
	err = s.Request.Validate()
	if err != nil {
		return err
	}
	err = s.Response.Validate()
	return err
}
