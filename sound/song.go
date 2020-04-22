package sound

import "fmt"

// Song represents a song to played by SM-808
type Song struct {
	bpm      int      // the beats per minute of the song
	Title    string   // the title of the song
	Sequence []string // the sequence that makes up the beat of the song (built from the supplied pattern)
	Duration float64  // length of the song in seconds
	Step     float64  // length of a step in seconds
}

// NewSong returns a song to be played by SM-808, or an error indicating why the song could not be created
// Accepts a bpm (integer), title (string), duration (float64), and a pattern (map[string][8]int)
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
		bpm:      bpm,
		Title:    title,
		Duration: duration,
	}
	s.generateSongSequence(pattern)
	s.calculateStep()
	return s, nil
}

// generateSongSequence converts a pattern into a a sequence that can be played by the Player
// accepts a pattern (map[string][8]int)
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

// calculateStep calculates the duration of a step in the song given the bpm
func (s *Song) calculateStep() {
	s.Step = ((60.0 / float64(s.bpm)) * 4.0) / 8.0
}

// Play creates a new player and plays the song
func (s *Song) Play() {
	p := NewPlayer()
	p.Play(*s)
}
