package network

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

func GetExternalIp() (string, error) {
	var err error
	var body []byte
	var resp *http.Response
	resp, err = http.Get("https://myexternalip.com/raw")
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	ip := string(body)
	return ip, nil
}