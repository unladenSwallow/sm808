package sound

import "fmt"

type Song struct {
	BPM      int      // the beats per minute of the song
	Title    string   // the title of the song
	Sequence []string // the sequence that makes up the beat of the song
	Length   float64  // length of the song in seconds
	Step 	 float64 // length of a step in seconds
}

func NewSong(bpm int, title string, length float64, pattern map[string][8]int) *Song {
	s := &Song{
		BPM: bpm,
		Title: title,
		Length: length,
	}
	s.generateSongSequence(pattern)
	s.calculateStep()
	return s
}

func (s *Song) generateSongSequence(pattern map[string][8]int) {
	s.Sequence = make([]string, 8)
	for i := 0; i < 8; i++ {
		beat := "|"
		for k, v := range pattern {
			if v[i] > 0 {
				if beat != "|" {
					beat += "+"
				}
				beat += fmt.Sprintf("%v", k)
			}
		}
		if beat == "|" {
			beat += "_"
		}
		s.Sequence[i] = beat
		if i == 7 {
			s.Sequence[i] += "|"
		}
	}
}

func (s *Song) calculateStep() {
	s.Step = ((60.0/float64(s.BPM)) * 4.0) / 8.0
	fmt.Println("LESLIE -- s.Step %v", s.Step)
}

func (s *Song) Play() {
	p := NewPlayer()
	p.Play(*s)
}
