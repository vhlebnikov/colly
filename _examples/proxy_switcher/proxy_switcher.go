package main

import (
	"bytes"
	"log"

	"github.com/vhlebnikov/colly/v2"
	"github.com/vhlebnikov/colly/v2/proxy"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(colly.AllowURLRevisit())

	// Rotate two socks5 proxies
	rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:1337", "socks5://127.0.0.1:1338")
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	// Print the response
	c.OnResponse(func(r *colly.Response) {
		log.Printf("Proxy Address: %s\n", r.Request.ProxyURL)
		log.Printf("%s\n", bytes.Replace(r.Body, []byte("\n"), nil, -1))
	})

	// Print the error in case of proxy problems
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Proxy Address: %s\n", r.Request.ProxyURL)
		log.Printf("Error: %s\n", err)
	})

	// Fetch httpbin.org/ip five times
	for i := 0; i < 5; i++ {
		c.Visit("https://httpbin.org/ip")
	}
}
