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
			nominees{list: []string{}, maxLength: 5},
			nominees{list: []string{"Bidoof"}, maxLength: 5},
		},
		{
			"Bidoof",
			nominees{
				list:      []string{"Tortank", "Pikachu", "Dracofeu", "Salamèche", "Goélise"},
				maxLength: 5,
			},
			nominees{
				list:      []string{"Pikachu", "Dracofeu", "Salamèche", "Goélise", "Bidoof"},
				maxLength: 5,
			},
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
