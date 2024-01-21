package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 106 {
		t.Errorf("expected deck lenght of 16, but got %v", len(d))
	}
}
