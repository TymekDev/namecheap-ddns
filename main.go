package main

import (
	"log"
	"os"
	"time"

	"github.com/spf13/pflag"
)

// Reference: https://www.namecheap.com/support/knowledgebase/article.aspx/29/11/how-to-dynamically-update-the-hosts-ip-with-an-http-request/

func main() {
	hosts := pflag.StringSlice("host", nil, "hosts (subdomains) to be updated")
	domain := pflag.String("domain", "", "domain name to be updated")
	password := pflag.String("password", "", "password to be used")
	interval := pflag.Duration("interval", 30*time.Second, "time between subsequent IP update requests")
	pflag.Parse()

	ok := true
	if len(*hosts) == 0 {
		ok = false
		log.Println("ERROR", "missing hosts")
	}
	if *domain == "" {
		ok = false
		log.Println("ERROR", "missing domain name")
	}
	if *password == "" {
		ok = false
		log.Println("ERROR", "missing password")
	}
	if !ok {
		os.Exit(1)
	}

	log.Println("INFO", "picked up hosts:", *hosts)
	log.Println("INFO", "picked up domain:", *domain)
	log.Println("INFO", "picked up password")
	log.Println("INFO", "picked up interval:", *interval)

	round := func() {
		for _, host := range *hosts {
			log.Printf("INFO updating (host: %s, domain: %d)\n", host, domain)
			if err := updateIP(host, *domain, *password); err != nil {
				log.Println("ERROR", err)
			}
		}
	}

	log.Println("INFO", "starting...")
	round()
	for range time.Tick(*interval) {
		round()
	}
}
