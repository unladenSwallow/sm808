package config

import (
	"strconv"
)

// Args struct matches accepted command line arguments for the sm-808
type Args struct {
	Bpm      int // the beats per minute
	Duration float64 // the song duration to play
}

// NewArgs returns the Args parsed from the command line arguments
// Accepts a slice of strings representing the command line arguments supplied
func NewArgs(args []string) (*Args, error) {
	var bpm int
	var duration float64
	var err error
	if len(args) > 0 {
		for k, a := range args {
			switch a {
			case "-bpm":
				bpm, err = strconv.Atoi(args[k+1])
				if err != nil {
					return nil, err
				}
				break
			case "-duration":
				duration, err = strconv.ParseFloat(args[k+1], 4)
				if err != nil {
					return nil, err
				}
				break
			}
		}
	}
	return &Args{
		Bpm:      bpm,
		Duration: duration,
	}, nil
}
