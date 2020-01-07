package main

import "testing"

func Test_fuels(t *testing.T) {
	testCases := []struct {
		name string
		consumption int
		volume int
		want int
	}{
		{"Fuel reserve more than consumption", 8,15,187},
		{"Fuel reserve equal to consumption", 8,8,100},
		{"Fuel reserve less than consumption", 10,5,50},
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got :=fuels(testCase.consumption,testCase.volume)
		if got != testCase.want {
			t.Error("Fuels with test:", testCase.name, "got:", got, "want:", testCase.want)
		}
	}
}