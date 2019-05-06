package cpt2

import (
	"fmt"
	"math"
	"time"
)

func TimeCount() {
	for x := 0; x <= 9; x++ {
		n := int(math.Pow(10, float64(x)))
		startTime := time.Now()

		sum := 0
		for i := 0; i < n; i++ {
			sum += i
		}

		timePass := time.Since(startTime)
		fmt.Printf("10^%d : %f s\n", x, timePass.Seconds())
	}
}
