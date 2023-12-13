package mongo

import (
	"errors"
	"math/rand"
	"short-it/internal/logger"
	"strconv"
	"sync"
)

// Mongo is a struct that implements DB interface
type Mongo struct {
	m map[int]string
}

var (
	instance *Mongo
	one      sync.Once
)

func init() {
	one.Do(func() {
		instance = &Mongo{m: make(map[int]string)}
	})
}

func GetInstance() *Mongo {
	return instance
}

// Save saves the url to the database
func (m *Mongo) Save(url string) (int, error) {
	// generate a random int
	id := rand.Intn(10000)
	logger.Info("generated id: " + strconv.Itoa(id))
	m.m[id] = url
	return id, nil
}

// Find finds the url from the database
func (m *Mongo) Find(id int) (string, error) {
	url, ok := m.m[id]
	if !ok {
		logger.Error("url not found")
		return "", errors.New("url not found")
	}
	return url, nil
}
