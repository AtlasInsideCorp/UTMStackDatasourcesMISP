package controller

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// getData allows you to make a get request to the MISP api and save the response
func getData(key, url string, data, body interface{}) error {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	clientHttp := &http.Client{Transport: transCfg}
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error to marshaling body: %v", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJson))
	if err != nil {
		return fmt.Errorf("error to create a request: %v", err)
	}

	req.Header.Add("Authorization", key)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return fmt.Errorf("error doing the request: %v", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("error: statuscode: %d", resp.StatusCode)
	}
	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading the response: %v", err)
	}
	err = json.Unmarshal(bodyResp, data)
	if err != nil {
		return fmt.Errorf("error unmarshal json response: %v", err)
	}
	return nil
}
