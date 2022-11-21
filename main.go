package main

import (
	"log"
	"time"

	"github.com/spf13/pflag"
)

// Reference: https://www.namecheap.com/support/knowledgebase/article.aspx/29/11/how-to-dynamically-update-the-hosts-ip-with-an-http-request/

func main() {
	hosts := pflag.StringSlice("host", nil, "hosts (subdomains) to be updated")
	domain := pflag.String("domain", "", "domain to be updated")
	password := pflag.String("password", "", "password to be used")
	interval := pflag.Duration("interval", 30*time.Second, "time between subsequent IP update requests")
	pflag.Parse()

	log.Println("starting with interval", *interval)
	for range time.Tick(*interval) {
		for _, host := range *hosts {
			if err := updateIP(host, *domain, *password); err != nil {
				log.Println(err)
			}
		}
	}
}
