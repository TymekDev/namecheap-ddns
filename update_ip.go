package main

import (
	"log"
	"net/http"
)

func updateIP(host, domain, password string) error {
	req, err := createRequest(host, domain, password)
	if err != nil {
		return err
	}

	return sendRequest(req)
}

func createRequest(host, domain, password string) (*http.Request, error) {
	req, err := http.NewRequest("GET", "https://dynamicdns.park-your-domain.com/update", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("host", host)
	q.Add("domain", domain)
	q.Add("password", password)
	req.URL.RawQuery = q.Encode()

	return req, nil
}

func sendRequest(req *http.Request) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("ERROR", err)
		}
	}()

	result, err := parse(resp.Body)
	if err != nil {
		return err
	}

	return result
}
