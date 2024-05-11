package api

import (
	"crypto/tls"
	"net/http"
	"time"
)

func NewClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // #nosec G402 -- We're okay with insecure internal communication
			},
		},
	}
}
