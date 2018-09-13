package slack

import (
	"reflect"
	"testing"
)

func TestFindChecksThePresenceOfTheGivenArgument(t *testing.T) {
	cases := []struct {
		label       string
		n           string
		h           []string
		expectation bool
	}{
		{
			"When the slice contains multiple strings including the given string",
			"Caninos",
			[]string{"Simularbre", "Caninos", "Noctali", "Gobou"},
			true,
		},
		{
			"When the slice contains only the given string",
			"Caninos",
			[]string{"Caninos"},
			true,
		},
		{
			"When the slice contains multiple strings but not the given string",
			"Balignon",
			[]string{"Simularbre", "Caninos", "Noctali", "Gobou"},
			false,
		},
		{
			"When the slice is empty and a string is given",
			"Caninos",
			[]string{},
			false,
		},
		{
			"When the slice contains multiple strings and an empty string is given",
			"",
			[]string{"Simularbre", "Caninos", "Noctali", "Gobou"},
			false,
		},
		{
			"When the slice is empty and an empty string is given",
			"",
			[]string{},
			false,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.label, func(t *testing.T) {
			result := find(testCase.n, testCase.h)
			if result != testCase.expectation {
				t.Errorf("Expected %t, got %t when finding %s in %s", testCase.expectation, result, testCase.n, testCase.h)
			}
		})
	}
}

func TestFilterReturnsANewSliceWithoutTheExcludedVAlues(t *testing.T) {
	cases := []struct {
		in             []string
		excludedValues []string
		expectation    []string
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

		if !reflect.DeepEqual(result, testCase.expectation) {
			t.Errorf("Expected %s, got %s", testCase.expectation, result)
		}
	}
}
