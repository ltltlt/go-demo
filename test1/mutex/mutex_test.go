package mutex_test

import (
	"sync"
	"testing"
)

func TestMutex(t *testing.T) {
	var m sync.Mutex
	m.Lock()
	m.Lock()
	m.Unlock()
}
