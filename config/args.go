package config

import (
	"strconv"
)

type Args struct {
	Bpm      int
	Duration float64
}

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
