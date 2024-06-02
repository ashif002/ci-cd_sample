package main

import "testing"

func TestRunnig(t *testing.T) {
	ecxpect := "Running Successfully."
	if result := Printing(); result != ecxpect {
		t.Error("ecxpect ", ecxpect, "but got ", result)
	}
}

func Printing() string {
	return "Running Successfully."
}
