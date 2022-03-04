package raceconditioninportgeneration

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	var wg sync.WaitGroup
	var ports sync.Map

	getRamdonPort := func(index int) {
		defer wg.Done()

		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		port, err := RandomUnusedPort()
		if err != nil {
			t.Errorf("Index %d failed to get ramdom port: %s", index, err)
		}
		_, loaded := ports.LoadOrStore(port, true)
		if loaded {
			t.Errorf("Index %d got repeated port: %d", index, port)
		}
		t.Logf("Index %d got port: %d", index, port)
	}

	const N = 1000
	wg.Add(N)
	for i := 0; i < N; i++ {
		go getRamdonPort(i)
	}
	wg.Wait()
}
