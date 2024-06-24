package main

import (
	"net/http"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

var jsonToReturn = `
{
  "ts": 1719141806020,
  "tsj": 1719141805343,
  "date": "Jun 23rd 2024, 07:23:25 am NY",
  "items": [
    {
      "curr": "USD",
      "xauPrice": 2320.79,
      "xagPrice": 29.5425,
      "chgXau": -40.22,
      "chgXag": -1.228,
      "pcXau": -1.7035,
      "pcXag": -3.9908,
      "xauClose": 2361.01,
      "xagClose": 30.7705
    }
  ]
}
`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}
