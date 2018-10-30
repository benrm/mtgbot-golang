package main

import (
	"testing"
)

func TestAllowedCardRarity(t *testing.T) {
	if !allowedCardRarity("Common") {
		t.Fatal("Command returned false for input 'Common'")
	}
	if !allowedCardRarity("Uncommon") {
		t.Fatal("Command returned false for input 'Uncommon'")
	}
	if !allowedCardRarity("Rare") {
		t.Fatal("Command returned false for input 'Rare'")
	}
	if !allowedCardRarity("Mythic Rare") {
		t.Fatal("Command returned false for input 'Mythic Rare'")
	}
	if !allowedCardRarity("Basic Land") {
		t.Fatal("Command returned false for input 'Basic Land'")
	}
	if allowedCardRarity("Masterpiece") {
		t.Fatal("Command returned true for input 'Masterpiece'")
	}
}
