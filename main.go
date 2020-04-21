package main

import (
	"fmt"
	"github.com/unladenSwallow/sm808/sound"
)

func main() {
	pattern := make(map[string][8]int)
	pattern["kick"] = [8]int{0, 1, 0, 0, 1, 0, 0, 1}
	pattern["snare"] = [8]int{1, 1, 1, 0, 1, 1, 1, 0}
	pattern["hi-hat"] = [8]int{0, 1, 1, 0, 0, 1, 1, 0}
	song := sound.NewSong(128.0, "Hazaa!", 30.0, pattern)
	fmt.Printf("Playing %v\n", song.Title)
	song.Play()
}
