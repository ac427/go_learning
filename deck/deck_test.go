package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("expected length of 16, but got %v ", len(d))
	}

	if d[len(d)-1] != "Heart of Four" {
		t.Errorf("expected Heart of Four, but got %v", d[15])
	}

	if d[0] != "Diamonds of Ace" {
		t.Errorf("expected Diamonds of Ace, but got %v", d[15])
	}

}

func TestSavetoDeck(t *testing.T) {

	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	_, err := os.ReadFile("_decktesting")
	if err != nil {
		t.Errorf("failed to read file _decktesting %v", err)
	}

	f, _ := os.Stat("_decktesting")
	if f.Mode().Perm() != 0o644 {
		t.Errorf("incorrect permissions %s (0%o), must be 0644 ", f.Mode().Perm(), f.Mode().Perm())
	}

}
