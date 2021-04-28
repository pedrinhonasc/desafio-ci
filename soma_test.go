package main

import "testing"

func TestSomaCorrectResult(t *testing.T) {
	expectedResult := 10
	result := soma(5, 5)
	if result != expectedResult {
		t.Errorf("Expected %v but got %v", expectedResult, result)
	}

	expectedResult = 15
	result = soma(5, 10)
	if result != expectedResult {
		t.Errorf("Expected %v but got %v", expectedResult, result)
	}
}

func TestSomaWrongResult(t *testing.T) {
	expectedResult := 10
	result := soma(5, 15)
	if result == expectedResult {
		t.Errorf("Expected %v but got %v", expectedResult, result)
	}
}