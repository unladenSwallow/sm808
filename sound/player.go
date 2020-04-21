package sound

import (
	"fmt"
	"time"
)

// Player represents a music player
type Player struct {
	StepTicker *time.Ticker // StepTicker ticks at an interval that corresponds to one "step" in a song
}

// NewPlayer returns a new player ready to play a song
func NewPlayer() Player {
	return Player{}
}

// Play plays a song
// Accepts: song, a struct representing the song to be played
func (p Player) Play(song Song) {
	p.StepTicker = time.NewTicker(time.Duration(song.Step*1000) * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			for i := 1; i < 9; i++ {
				select {
				case <-done:
					return
				case <-p.StepTicker.C:
					p.PlayStep(song.Sequence[i-1], i)
				}
			}
		}
	}()
	time.Sleep(time.Duration(song.Duration) * time.Second)
	p.StepTicker.Stop()
	done <- true
	fmt.Println("\nFIN!")
}

// PlayStep "plays" one step of the song.
func (p Player) PlayStep(step string, i int) {
	if i%8 == 0 {
		step += "\n"
	}
	fmt.Print(step)
}
