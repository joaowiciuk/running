package running

import (
	"fmt"
	"math"
	"time"

	"github.com/montanaflynn/stats"
)

// Sizer is implemented by any value that has a Size method.
// The Size method is used to measure the size of inputs to algorithms.
type Sizer interface {
	Size() int
}

// R is a type passed to Run functions to check running time and support formatted run logs.
type R struct {
	// runningTimes maps each input with size n to its running time.
	runningTimes map[int]time.Duration

	// N is the number of times the algorithm should run for each input.
	N int
}

// Run analyzes a single input and registers its running time.
func (r *R) Run(input Sizer, algorithm func(interface{}) interface{}) {
	runningTimes := make([]float64, 0)
	for i := 0; i < r.N; i++ {
		start := time.Now()
		algorithm(input)
		runningTimes = append(runningTimes, float64(time.Since(start).Nanoseconds()))
	}
	mean, err := stats.Mean(runningTimes)
	if err != nil {
		panic(fmt.Errorf("running.Run: calculating mean delay: %v", err))
	}
	rounded := math.Round(mean)
	r.runningTimes[input.Size()] = time.Duration(int64(rounded))
}
