# Namecheap DDNS Updater

A simple CLI for updating dynamic DNS at Namecheap.

[Reference](https://www.namecheap.com/support/knowledgebase/article.aspx/29/11/how-to-dynamically-update-the-hosts-ip-with-an-http-request/)

## Usage
Default interval is low because there is no checks for IP change.
In other words, an update request is being made every time.
```
--domain string       domain name to be updated
--host strings        hosts (subdomains) to be updated
--interval duration   time between subsequent IP update requests (default 30m0s)
--password string     password to be used
```

## Tips
If you need several domains pointing to the same IP, then there is no need to make multiple DDNS records.
Use a domain with DDNS record and point the other domains to the DDNS one using CNAME records.

Note: a top level domain cannot have a CNAME record. You might need to make that one DDNS as well.
