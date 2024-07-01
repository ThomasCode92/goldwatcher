package main

import "testing"

func TestApp_getToolBar(t *testing.T) {
	tb := testApp.getToolBar()

	if len(tb.Items) != 4 {
		t.Errorf("Expected 4 items, got %v", len(tb.Items))
	}
}
