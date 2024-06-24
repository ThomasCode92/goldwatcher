package main

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestGold_GetPrices(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
			Header:     make(http.Header),
		}
	})

	g := Gold{
		Client: client,
		Prices: nil,
	}

	p, err := g.GetPrices()
	if err != nil {
		t.Error(err)
	}

	if p.Price != 2320.79 {
		t.Errorf("Expected 2320.79, got %v", p.Price)
	}
}
