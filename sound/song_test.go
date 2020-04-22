package sound

import (
	"fmt"
	"testing"
)

func Test_NewSong(t *testing.T) {
	type testData struct {
		bpm          int
		title        string
		duration     float64
		pattern      map[string][8]int
		errResponse  error
		songResponse *Song
	}
	tests := map[string]testData{
		"HappyPath": {
			bpm:      128,
			title:    "test",
			duration: 100,
			pattern: map[string][8]int{
				"test1": {0, 1, 0, 1, 0, 1, 1, 0},
				"test2": {0, 1, 0, 1, 0, 1, 1, 0},
				"test3": {0, 1, 0, 1, 0, 1, 1, 0},
			},
			songResponse: &Song{
				Title:    "test",
				BPM:      128,
				Duration: 100,
				Sequence: []string{"|_", "|test1+test2+test3", "|_", "|test1+test2+test3", "|_", "|test1+test2+test3", "|test1+test2+test3", "|_|"},
				Step:     0.234375,
			},
		},
		"Error-NegativeBPM": {
			bpm:      -128,
			title:    "test",
			duration: 100,
			pattern: map[string][8]int{
				"test1": {0, 1, 0, 1, 0, 1, 1, 0},
				"test2": {0, 1, 0, 1, 0, 1, 1, 0},
				"test3": {0, 1, 0, 1, 0, 1, 1, 0},
			},
			errResponse: fmt.Errorf("bpm must be between 1 and 999"),
		},
		"Error-NegativeDuration": {
			bpm:      128,
			title:    "test",
			duration: -100,
			pattern: map[string][8]int{
				"test1": {0, 1, 0, 1, 0, 1, 1, 0},
				"test2": {0, 1, 0, 1, 0, 1, 1, 0},
				"test3": {0, 1, 0, 1, 0, 1, 1, 0},
			},
			errResponse: fmt.Errorf("song duration must be between 1 and 999"),
		},
		"Error-NoPattern": {
			bpm:         128,
			title:       "test",
			duration:    100,
			errResponse: fmt.Errorf("song requires a pattern"),
		},
	}
	for _, v := range tests {
		s, err := NewSong(v.bpm, v.title, v.duration, v.pattern)
		if v.errResponse != nil {
			if v.errResponse.Error() != err.Error() {
				t.Error(fmt.Sprintf("Expected err to be %v but got %v", v.errResponse, err))
			} else if s != nil {
				t.Error(fmt.Sprintf("Expected song to be nil but got: %v", s))
			}
		} else {
			if v.songResponse != s {
				if s != nil && len(s.Sequence) != len(v.songResponse.Sequence) {
					t.Error(fmt.Sprintf("Expected song to be %v but got %v", v.songResponse, s))
				}
			}
		}
	}
}
