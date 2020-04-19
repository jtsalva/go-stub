package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	StubsDirectory  string        `short:"d" long:"directory" description:"Path to directory containing yml stub files" required:"true"`
	Port            int           `short:"p" long:"port" description:"Port to serve stubs from" required:"true"`
	CorsAllowOrigin string        `short:"o" long:"cors-allow-origin" description:"Blanket allow CORS from specified origin"`
	WriteTimeout    time.Duration `long:"write-timeout" description:"Server write timeout"`
	ReadTimeout     time.Duration `long:"read-timeout" description:"Server read timeout"`
	IdleTimeout     time.Duration `long:"idle-timeout" description:"Server idle timeout"`
	Stubs           []Stub
}

func (c *Config) IsCorsEnabled() bool {
	return c.CorsAllowOrigin != ""
}

func (c *Config) LoadStubs() error {
	fmt.Printf("Loading test-stubs from %s\n", c.StubsDirectory)

	stubFilePaths, err := walkMatch(c.StubsDirectory, "*.yml")
	if err != nil {
		return errors.Wrapf(err, "failed to read stub directory: '%s'", c.StubsDirectory)
	}

	var allStubs []Stub
	for _, filePath := range stubFilePaths {
		fmt.Printf("Reading %s\n", filePath[strings.LastIndex(filePath, "/")+1:])
		file, err := ioutil.ReadFile(filePath)
		if err != nil {
			return errors.Wrapf(err, "failed to read stub file '%s': %v", filePath, err.Error())
		}

		var stubs []Stub
		err = yaml.Unmarshal(file, &stubs)
		if err != nil {
			return errors.Wrapf(err, "failed to unmarshal stub file '%s': %v", filePath, err)
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
