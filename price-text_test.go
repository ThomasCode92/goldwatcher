package main

import "testing"

func TestApp_GetPriceTe(t *testing.T) {
	open, _, _ := testApp.GetPriceText()
	if open.Text != "Open: $2361.0100 USD" {
		t.Errorf("Expected 'Open: $2361.0100 USD', got %v", open.Text)
	}
}
