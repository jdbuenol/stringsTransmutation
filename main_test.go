package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	got := Reverse("abecedario")
	if got != "oiradeceba" {
		t.Errorf("Reverse(\"abecedario\") = %s; want oiradeceba", got)
	}
	got2 := Reverse("123")
	if got2 != "321" {
		t.Errorf("Reverse(\"123\") = %s; want 321", got2)
	}
	got3 := Reverse("")
	if got3 != "" {
		t.Errorf("Reverse(\"\") = %s; want \"\"", got3)
	}
}

func TestConcat(t *testing.T) {
	got := Concat("abece", "dario")
	if got != "abecedario" {
		t.Errorf("Concat(\"abece\", \"dario\") = %s; want abecedario", got)
	}
	got2 := Concat("123", "456")
	if got2 != "123456" {
		t.Errorf("Concat(\"123\", \"456\") = %s; want 123456", got2)
	}
	got3 := Concat("", "")
	if got3 != "" {
		t.Errorf("Concat(\"\", \"\") = %s; want \"\"", got3)
	}
}
