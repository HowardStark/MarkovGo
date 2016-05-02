package main

// Cache interface describes the basic
// functions a cache module must be able
// to perform.
type Cache interface {
	AddPair(key string, value string) error
	GetRandom(key string) (string, error)
}
