package crypter

import (
	"testing"
)

func TestCLI(t *testing.T) {

}

func TestToBase64(t *testing.T) {
	var testString string = "hello world"
	var expectedBase64String string = "aGVsbG8gd29ybGQ="

	result := toBase64([]byte(testString))
	if string(result) != expectedBase64String {
		t.Errorf("toBase64('hello world') got %v, expected %v", result, expectedBase64String)
	}

}

func TestfromBase64(t *testing.T) {
	var testData string = "aGVsbG8gd29ybGQ="
	expectedResult := "hello world"

	result, _ := fromBase64([]byte(testData))

	if string(result) != expectedResult {
		t.Errorf("fromBase64('aGVsbG8gd29ybGQ=') got %v, expected: %v", string(result), expectedResult)
	}
}
