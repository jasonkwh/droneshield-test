package ballast

import (
	"time"

	"k8s.io/apimachinery/pkg/api/resource"
)

// Create - allocates a byte array to the configured size.
func Create(quantity string) (int64, error) {
	interval := time.Minute

	q, err := resource.ParseQuantity(quantity)
	if err != nil {
		return 0, err
	}

	b := make([]byte, q.Value())
	i := 0
	go func() {
		for {
			// enter an endless loop which reads the first byte of the array every minute
			// this tricks the GC into thinking that the memory is live so it doesn't clean it up.
			n := b[i]
			i += int(n) + 1 // n is always 0 BUT the compiler doesnt know that
			if i >= len(b) {
				i = 0
			}
			time.Sleep(interval)
		}
	}()
	return q.Value(), nil
}
