package utils

import (
	"math/rand"
	"sync"
	"time"
)

type RandomWeight struct {
	Gen rand.Rand
}

var (
	instance *RandomWeight
	once     = sync.Once{}
)

func GetRngInstance() *RandomWeight {
	once.Do(func() {
		instance = &RandomWeight{Gen: *rand.New(rand.NewSource(time.Now().UnixNano()))}
	})
	return instance
}
