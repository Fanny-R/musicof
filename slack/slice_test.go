package slack

import (
	"reflect"
	"testing"
)

func TestContainsChecksThePresenceOfTheGivenArgument(t *testing.T) {
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
			result := contains(testCase.n, testCase.h)
			if result != testCase.expectation {
				t.Errorf("Expected %t, got %t when checking if %s contains %s", testCase.expectation, result, testCase.h, testCase.n)
			}
		})
	}
}

func TestFilterReturnsANewSliceWithoutTheExcludedValues(t *testing.T) {
	cases := []struct {
		label          string
		in             []string
		excludedValues []string
		expectation    []string
	}{
		{
			"When the slice contains multiple string and one contained string is given",
			[]string{"Simularbre", "Caninos", "Noctali", "Gobou"},
			[]string{"Caninos"},
			[]string{"Simularbre", "Noctali", "Gobou"},
		},
		{
			"When the slice contains only the given string",
			[]string{"Caninos"},
			[]string{"Caninos"},
			[]string{},
		},
		{
			"When the slice contains multiple strings and contained strings are given",
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
			"When the slice contains multiple strings and an uncontained string is given",
			[]string{"Simularbre", "Noctali", "Gobou"},
			[]string{"Caninos"},
			[]string{"Simularbre", "Noctali", "Gobou"},
		},
		{
			"When the slice contains multiple strings and no strings are given",
			[]string{"Simularbre", "Noctali", "Gobou"},
			[]string{},
			[]string{"Simularbre", "Noctali", "Gobou"},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.label, func(t *testing.T) {
			result := filter(testCase.in, testCase.excludedValues...)

			if !reflect.DeepEqual(result, testCase.expectation) {
				t.Errorf("Expected %s, got %s", testCase.expectation, result)
			}
		})
	}
}
