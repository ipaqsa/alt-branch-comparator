package VersionsComparator

import "testing"

func TestSecondVersionLessFirst(t *testing.T) {
	testTable := []struct {
		v1       string
		v2       string
		expected bool
	}{
		{
			v1:       "1.4.1",
			v2:       "3.5",
			expected: false,
		},
		{
			v1:       "6.2.0",
			v2:       "6.10.2",
			expected: false,
		},
		{
			v1:       "12",
			v2:       "11",
			expected: true,
		},
		{
			v1:       "12.2",
			v2:       "11",
			expected: true,
		},
		{
			v1:       "0.1.2",
			v2:       "0.11.2",
			expected: false,
		},
	}
	for _, testCase := range testTable {
		result := SecondVersionLessFirst(testCase.v1, testCase.v2)
		t.Logf("V1 %s V2 %s, result %t\n", testCase.v1, testCase.v2, result)
		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %t, got %t", testCase.expected, result)
		}
	}
}
