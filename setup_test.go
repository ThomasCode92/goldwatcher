package main

import (
	"bytes"
	"goldwatcher/repository"
	"io"
	"net/http"
	"os"
	"testing"

	"fyne.io/fyne/v2/test"
)

var testApp Config

func TestMain(m *testing.M) {
	a := test.NewApp()

	testApp.App = a
	testApp.MainWindow = a.NewWindow("")
	testApp.DB = repository.NewTestRepository()
	testApp.HTTPClient = client

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

var client = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})
