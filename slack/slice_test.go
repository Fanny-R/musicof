package slack

import (
	"reflect"
	"testing"
)

func TestFindReturnsTrueWhenSliceContainsGivenString(t *testing.T) {
	cases := []struct {
		n string
		h []string
	}{
		{"Caninos", []string{"Simularbre", "Caninos", "Noctali", "Gobou"}},
		{"Caninos", []string{"Caninos"}},
	}

	for _, testCase := range cases {
		if find(testCase.n, testCase.h) != true {
			t.Errorf("Expected true, got false when finding %s in %s", testCase.n, testCase.h)
		}
	}
}

func TestFindReturnsFalseWhenSliceDoesNotContainGivenString(t *testing.T) {
	cases := []struct {
		n string
		h []string
	}{
		{"Balignon", []string{"Simularbre", "Caninos", "Noctali", "Gobou"}},
		{"Balignon", []string{}},
		{"Balignon", []string{"Caninos"}},
		{"", []string{"Simularbre", "Caninos", "Noctali", "Gobou"}},
	}

	for _, testCase := range cases {
		if find(testCase.n, testCase.h) != false {
			t.Errorf("Expected false, got true when finding %s in %s", testCase.n, testCase.h)
		}
	}
}

func TestFilterReturnsANewSliceWithoutTheExcludedVAlues(t *testing.T) {
	cases := []struct {
		in             []string
		excludedValues []string
		expectedResult []string
	}{
		{
			[]string{"Simularbre", "Caninos", "Noctali", "Gobou"},
			[]string{"Caninos"},
			[]string{"Simularbre", "Noctali", "Gobou"},
		},
		{
			[]string{"Caninos"},
			[]string{"Caninos"},
			[]string{},
		},
		{
			[]string{"Simularbre", "Noctali", "Caninos", "Gobou"},
			[]string{"Caninos", "Noctali"},
			[]string{"Simularbre", "Gobou"},
		},
		{
			[]string{"Simularbre", "Caninos", "Noctali", "Gobou"},
			[]string{"Caninos", "Évoli"},
			[]string{"Simularbre", "Noctali", "Gobou"},
		},
		{
			[]string{"Simularbre", "Noctali", "Gobou"},
			[]string{"Caninos"},
			[]string{"Simularbre", "Noctali", "Gobou"},
		},
		{
			[]string{"Simularbre", "Noctali", "Gobou"},
			[]string{},
			[]string{"Simularbre", "Noctali", "Gobou"},
		},
	}

	for _, testCase := range cases {
		result := filter(testCase.in, testCase.excludedValues...)

		if reflect.DeepEqual(result, testCase.expectedResult) != true {
			t.Errorf("Expected %s, got %s", testCase.expectedResult, result)
		}
	}
}
