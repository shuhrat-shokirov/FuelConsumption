package main

import "testing"

func Test_fuels(t *testing.T) {
	testCases := []struct {
		name string
		consumption int
		volume int
		want int
	}{
		{"First car", 8,5,62},
		{"Second car", 8,15,187},
		{"Third car", 10,0,0},
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got :=fuels(testCase.consumption,testCase.volume)
		if got != testCase.want {
			t.Error("Fuels with test:", testCase.name, "got:", got, "want:", testCase.want)
		}
	}
}