package slack

import (
	"reflect"
	"testing"
)

func TestItAddsANewUserToTheNominees(t *testing.T) {
	cases := []struct {
		userID         string
		subject        nominees
		expectedResult nominees
	}{
		{
			"Bidoof",
			nominees{},
			nominees{"Bidoof"},
		},
		{
			"Bidoof",
			nominees{"Tortank", "Pikachu", "Dracofeu", "Salamèche", "Goélise"},
			nominees{"Pikachu", "Dracofeu", "Salamèche", "Goélise", "Bidoof"},
		},
	}

	for _, testCase := range cases {
		t.Run("", func(t *testing.T) {
			testCase.subject.Push(testCase.userID)
			if !reflect.DeepEqual(testCase.subject, testCase.expectedResult) {
				t.Errorf("Expected %v got %v", testCase.expectedResult, testCase.subject)
			}
		})
	}
}
