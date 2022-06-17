package mut

import (
	"math/rand"
	"sync"
)

var mut sync.RWMutex

func TenPercentsWriting() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	m := make(map[int]float32, 10)
	for i := 0; i < 10; i++ {
		m[i] = float32(i) / 3.0
	}

	go func() {
		mut.Lock()
		randKey := rand.Intn(10)
		m[randKey] = m[randKey] / 3.0

		wg.Done()
		mut.Unlock()
	}()

	for i := 0; i < 9; i++ {
		go func(i int) {
			mut.RLock()

			wg.Done()
			mut.RUnlock()
		}(i)
	}
	wg.Wait()
}

func FiftyPercentsWriting() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	m := make(map[int]float32, 10)
	for i := 0; i < 10; i++ {
		m[i] = float32(i) / 3.0
	}

	for i := 0; i < 5; i++ {
		go func(i int) {
			mut.Lock()
			m[i] = m[i] / 3.0
			wg.Done()
			mut.Unlock()
		}(i)
	}

	for i := 5; i < 10; i++ {
		go func(i int) {
			mut.RLock()
			wg.Done()
			mut.RUnlock()
		}(i)
	}
	wg.Wait()
}

func NintyPercentsWriting() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	m := make(map[int]float32, 10)
	for i := 0; i < 10; i++ {
		m[i] = float32(i) / 3.0
	}

	for i := 0; i < 9; i++ {
		go func(i int) {
			mut.Lock()
			m[i] = m[i] / 3.0

			wg.Done()
			mut.Unlock()
		}(i)
	}

	go func(i int) {
		mut.RLock()

		wg.Done()
		mut.RUnlock()
	}(9)

	wg.Wait()
}
