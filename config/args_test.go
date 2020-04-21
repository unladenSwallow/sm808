package config

import (
	"fmt"
	"testing"
)

func Test_NewArgs(t *testing.T) {
	type testData struct {
		argsReturn  *Args
		argsGiven   []string
		errResponse error
	}
	tests := map[string]testData{
		"HappyPath": {
			argsGiven: []string{"-bpm", "128", "-duration", "30"},
			argsReturn: &Args{
				Bpm:      128,
				Duration: 30,
			},
		}, "Error-NonInvalidBPM": {
			argsGiven:   []string{"-bpm", "abc", "-duration", "30"},
			errResponse: fmt.Errorf("strconv.Atoi: parsing \"abc\": invalid syntax"),
		}, "Error-InvalidDuration": {
			argsGiven:   []string{"-duration", "abc", "-bpm", "30"},
			errResponse: fmt.Errorf("strconv.ParseFloat: parsing \"abc\": invalid syntax"),
		},
	}
	for _, v := range tests {
		s, err := NewArgs(v.argsGiven)
		if v.errResponse != nil {
			if err != nil && v.errResponse.Error() != err.Error() {
				t.Error(fmt.Sprintf("Expected err to be %v but got %v", v.errResponse, err))
			} else if s != nil {
				t.Error(fmt.Sprintf("Expected args to be nil but got: %v", s))
			}
		} else {
			if v.argsReturn.Bpm != s.Bpm || v.argsReturn.Duration != s.Duration {
				t.Error(fmt.Sprintf("Expected song to be %v but got %v", v.argsReturn, s))
			}
		}
	}
}
