package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/goccy/go-yaml"
	"github.com/pkg/errors"
)

type Config struct {
	StubsDirectory  string        `short:"d" long:"directory" description:"Path to directory containing yml stub files" required:"true"`
	Port            int           `short:"p" long:"port" description:"Port to serve stubs from" required:"true"`
	CorsAllowOrigin string        `short:"o" long:"cors-allow-origin" description:"Allow CORS from specified origin, this will automatically apply headers 'Access-Control-Allow-Origin', 'Access-Control-Allow-Methods' and 'Access-Control-Expose-Headers'"`
	WriteTimeout    time.Duration `long:"write-timeout" description:"Server write timeout duration"`
	ReadTimeout     time.Duration `long:"read-timeout" description:"Server read timeout duration"`
	IdleTimeout     time.Duration `long:"idle-timeout" description:"Server idle timeout duration"`
	DisableColor    bool          `long:"disable-color" description:"Disable color in console output"`
	Stubs           []Stub
}

func (c *Config) IsCorsEnabled() bool {
	return c.CorsAllowOrigin != ""
}

func (c *Config) LoadStubs() error {
	fmt.Printf("Loading stubs from %s\n", color.BlueString(c.StubsDirectory))

	stubFilePaths, err := walkMatch(c.StubsDirectory, "*.stub.yml")
	if err != nil {
		return errors.Wrapf(err, "failed to read stub directory '%s'", c.StubsDirectory)
	}

	yamlReferenceDirs := yaml.ReferenceDirs(c.StubsDirectory)
	var allStubs []Stub
	for _, filePath := range stubFilePaths {
		fmt.Printf("Reading %s\n", color.BlueString(filePath[strings.LastIndex(filePath, "/")+1:]))
		file, err := ioutil.ReadFile(filePath)
		if err != nil {
			return errors.Wrapf(err, "failed to read stub file '%s'", filePath)
		}

		var stubs []Stub
		err = yaml.NewDecoder(bytes.NewBuffer(file), yamlReferenceDirs).Decode(&stubs)
		if err != nil {
			return errors.Wrapf(err, "failed to unmarshal stub file '%s'", filePath)
		}

		for _, stub := range stubs {
			if err = stub.Validate(); err != nil {
				return errors.Wrapf(err, "failed to validate stub in file '%s'", filePath)
			}
		}

		allStubs = append(allStubs, stubs...)
	}

	c.Stubs = allStubs
	return nil
}

func walkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}
