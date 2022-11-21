package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
)

// Reference: https://www.namecheap.com/support/knowledgebase/article.aspx/29/11/how-to-dynamically-update-the-hosts-ip-with-an-http-request/

func main() {
	hosts := pflag.StringSlice("host", nil, "hosts (subdomains) to be updated")
	domain := pflag.String("domain", "", "domain name to be updated")
	password := pflag.String("password", "", "password to be used")
	interval := pflag.Duration("interval", 30*time.Second, "time between subsequent IP update requests")
	pflag.Parse()

	const timeFormat = "2006-01-02 15:04 -0700"
	zerolog.TimeFieldFormat = timeFormat
	log.Logger = log.Output(zerolog.ConsoleWriter{TimeFormat: timeFormat, Out: os.Stderr})

	if *password == "" {
		log.Fatal().Msg("missing password")
	}
	if *domain == "" {
		log.Fatal().Msg("missing domain name")
	}
	if len(*hosts) == 0 {
		log.Fatal().Msg("missing hosts")
	}

	round := func() {
		for _, host := range *hosts {
			if err := updateIP(host, *domain, *password); err != nil {
				log.Error().Err(err).Send()
			}
		}
	}

	log.Info().Strs("hosts", *hosts).Str("domain", *domain).Dur("interval", *interval).Msg("starting...")
	round()
	for range time.Tick(*interval) {
		round()
	}
}
