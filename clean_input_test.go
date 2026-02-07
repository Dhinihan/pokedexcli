package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Hello World   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "Hello",
			expected: []string{"hello"},
		},
		{
			input:    "This is a VERY Long sentence",
			expected: []string{"this", "is", "a", "very", "long", "sentence"},
		},
	}
	for _, testCase := range cases {
		result := cleanInput(testCase.input)
		if len(result) != len(testCase.expected) {
			t.Errorf("Não tem o mesmo número de palavras do que o esperado")
		}
		for i, word := range result {
			if word != testCase.expected[i] {
				t.Errorf("%s é diferente de %s", word, testCase.expected[i])
			}
		}
	}
}
