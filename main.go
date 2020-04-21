package main

import (
	"fmt"
	config2 "github.com/unladenSwallow/sm808/config"
	"github.com/unladenSwallow/sm808/sound"
	"os"
)

func main() {
	args := os.Args[1:]
	config, err := config2.NewArgs(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	if config.Bpm == 0 {
		config.Bpm = 128
	}
	if config.Duration == 0 {
		config.Duration = 30
	}
	pattern := make(map[string][8]int)
	pattern["kick"] = [8]int{0, 0, 0, 0, 1, 0, 0, 1}
	pattern["snare"] = [8]int{1, 0, 1, 0, 1, 0, 1, 0}
	pattern["hi-hat"] = [8]int{0, 1, 1, 0, 0, 1, 1, 0}

	song, err := sound.NewSong(config.Bpm, "Hazaa!", config.Duration, pattern)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Playing %v\n", song.Title)
	song.Play()
}
