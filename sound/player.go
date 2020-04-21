package sound

import (
	"fmt"
	"time"
)

type Player struct {
	StepTicker *time.Ticker
}

func NewPlayer() Player {
	return Player {
	}
}

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
	time.Sleep(time.Duration(song.Length) * time.Second)
	p.StepTicker.Stop()
	done <- true
	fmt.Println("FIN!")
}

func (p Player) PlayStep(step string, i int) {
		if i%8 == 0 {
			step += "\n"
		}
		fmt.Print(step)
}
