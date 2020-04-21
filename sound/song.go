package sound

import "fmt"

type Song struct {
	BPM      int      // the beats per minute of the song
	Title    string   // the title of the song
	Sequence []string // the sequence that makes up the beat of the song
	Duration float64  // length of the song in seconds
	Step     float64  // length of a step in seconds
}

func NewSong(bpm int, title string, duration float64, pattern map[string][8]int) (*Song, error) {
	if pattern == nil {
		return nil, fmt.Errorf("song requires a pattern")
	}
	if bpm < 1 || bpm > 999 {
		return nil, fmt.Errorf("bpm must be between 1 and 999")
	}
	if duration < 1 || duration > 999 {
		return nil, fmt.Errorf("song duration must be between 1 and 999")
	}
	s := &Song{
		BPM:      bpm,
		Title:    title,
		Duration: duration,
	}
	s.generateSongSequence(pattern)
	s.calculateStep()
	return s, nil
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
	s.Step = ((60.0 / float64(s.BPM)) * 4.0) / 8.0
}

func (s *Song) Play() {
	p := NewPlayer()
	p.Play(*s)
}
