package client

import (
	"math/rand"
	"time"
)

// windSimulation - simulates how the wind affects to the coordinates
func (cl *client) windSimulation() {
	t := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-cl.done:
			t.Stop()
			return
		case <-t.C:
			cl.lock.Lock()
			longRand := randomBool()
			latRand := randomBool()
			if longRand {
				cl.coordinate.Longitude = cl.coordinate.Longitude + randomFloats(0.001, 0.005)
			} else {
				cl.coordinate.Longitude = cl.coordinate.Longitude - randomFloats(0.001, 0.005)
			}
			if latRand {
				cl.coordinate.Latitude = cl.coordinate.Latitude + randomFloats(0.001, 0.005)
			} else {
				cl.coordinate.Latitude = cl.coordinate.Latitude - randomFloats(0.001, 0.005)
			}
			cl.lock.Unlock()
		}
	}
}

func randomFloats(min, max float64) float64 {
	return rand.Float64() * (max - min)
}

func randomBool() bool {
	return rand.Intn(2) == 1
}
