package main

import (
	"testing"
)

func TestGold_GetPrices(t *testing.T) {
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
