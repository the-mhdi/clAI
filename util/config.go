package util

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Config struct {
	GEMINI_API_KEY   string  `json:"GEMINI_API_KEY"`
	Model            string  `json:"Model"`
	Temperature      float32 `json:"Temperature"`
	TopK             int32   `json:"TopK"`
	TopP             float32 `json:"TopP"`
	MaxOutput        int32   `json:"MaxOutput"`
	ResponseMIMEType string  `json:"ResponseMIMEType"`
}

func Newconfig() *Config {
	c := new(Config)
	c.setDefault()
	return c
}

func (c *Config) SetAPI_KEY(k string) {
	c.GEMINI_API_KEY = k
}

func (c *Config) setDefault() error {
	if c.GEMINI_API_KEY == "" {
		return errors.New("gimme the damn API key")
	}

	c = &Config{
		Model:            "gemini-1.5-pro-002",
		Temperature:      1.5,
		TopK:             40,
		TopP:             0.95,
		MaxOutput:        8192,
		ResponseMIMEType: "text/plain",
	}
	return nil
}
func LoadConfig(path string) (*Config, error) {
	return loadConfig(path)
}

func loadConfig(path string) (*Config, error) {

	dir, err := isDirectory(path)

	if err != nil {
		return nil, err
	}

	if dir {
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}

		defer f.Close()

		buf := make([]byte, 2048)

		_, err = f.Read(buf)

		if err != nil {
			return nil, err
		}
		conf := Newconfig()
		jsonstr := string(buf)
		cleanjson := strings.ReplaceAll(jsonstr, "\x00", "")
		err = json.Unmarshal([]byte(cleanjson), conf)

		if err != nil {
			return nil, err
		}
		return conf, err
	}

	if isURL(path) {
		resp, err := http.Get(path)
		if err != nil {
			return nil, err
		}
		buf := make([]byte, 2048)
		_, err = resp.Body.Read(buf)

		if err != nil {
			return nil, err
		}
		conf := Newconfig()
		jsonstr := string(buf)
		cleanjson := strings.ReplaceAll(jsonstr, "\x00", "")
		err = json.Unmarshal([]byte(cleanjson), conf)

		if err != nil {
			return nil, err
		}

		return conf, err
	}

	return nil, errors.New("everything's wrong")
}

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) { // If the path simply doesn't exist
			return false, nil
		}
		return false, err // Return other potential file system errors for upper layer to handle
	}
	return fileInfo.Mode().IsRegular(), nil
}

func isURL(str string) bool {
	u, err := url.Parse(str)
	if err != nil {
		return false
	}

	// Check if a scheme is present. HTTP and HTTPS are common, but
	// other schemes like FTP, file, or custom schemes are also valid.
	if u.Scheme == "" {
		return false
	}

	// Check for a host (domain or IP address). It should not be empty.
	if u.Host == "" {
		return false
	}

	// Optional further checks.

	return true

}
