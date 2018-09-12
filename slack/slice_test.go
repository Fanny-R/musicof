package slack

import (
	"testing"
)

func TestFindReturnsTrueWhenSliceContainsGivenString(t *testing.T) {
	n := "Caninos"
	h := []string{"Simularbre", "Caninos", "Noctali", "Gobou"}

	if find(n, h) != true {
		t.Errorf("Expected true, got false when finding %s in %s", n, h)
	}
}

func TestFindReturnsFalseWhenSliceDoesNotContainGivenString(t *testing.T) {
	n := "Balignon"
	h := []string{"Simularbre", "Caninos", "Noctali", "Gobou"}

	if find(n, h) != false {
		t.Errorf("Expected false, got true when finding %s in %s", n, h)
	}
}
