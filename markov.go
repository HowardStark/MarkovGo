package main

import "regexp"

// Markov is the main object for the
// markov chain system.
type Markov struct {
	cache Cache
}

// ParseText parses a string
func (markov *Markov) ParseText(text string) error {
	var validRegex = regexp.MustCompile(`([\w'-]+|[.,!?;&])`)
	_ = validRegex
	return nil
}
