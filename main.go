package main

import (
	"log"
	"net/http"
)

// Reference: https://www.namecheap.com/support/knowledgebase/article.aspx/29/11/how-to-dynamically-update-the-hosts-ip-with-an-http-request/

func main() {
	req, err := http.NewRequest("GET", "https://dynamicdns.park-your-domain.com/update", nil)
	if err != nil {
		log.Fatalln(err)
	}

	q := req.URL.Query()
	q.Add("host", "")     // TODO: to flag
	q.Add("domain", "")   // TODO: to flag
	q.Add("password", "") // TODO: to flag
	req.URL.RawQuery = q.Encode()

	if err := send(req); err != nil {
		log.Fatalln("Error:", err)
	}
}

func send(req *http.Request) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	result, err := parse(resp.Body)
	if err != nil {
		return err
	}

	return result
}
