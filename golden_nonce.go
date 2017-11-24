//Golden Nonce strategies simulation
package main

import (
	"fmt"
	"math/rand"
)

//Settings
var ITERATIONS_NUMBER int = 1     	//Number of iteration by each nonce size
var NONCE_SIZES []int32 = []int32{	//Nonse maximum sizes
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
	10000000,
	100000000,
}

func main() {
	fmt.Println("Simulation started")
	fmt.Println()
	fmt.Println("Nonce size | GoOver avg. percent | Picking avg. percent ")

	//Make simulation for each nonce size
	for _, nonceSize := range NONCE_SIZES {
		//Totals for current ITERATIONS_NUMBER
		var totalPercentGoOver, totalPercentPicking float32
		//Target nonce
		var targetNonce int32
		//Random sequence
		randomSequence := randomSequenceGenerator(nonceSize)

		//Simulation
		for j := 0; j < ITERATIONS_NUMBER; j++ {
			targetNonce = rand.Int31n(nonceSize)
			//Try GoOver strategy
			var i int32 = 0
			for i < nonceSize {
				if targetNonce == i {
					//We found target nonce for this iteration
					totalPercentGoOver = totalPercentGoOver + (float32(i) / float32(nonceSize) * 100)
					break
				} else {
					i++
				}
			}
			//Try Picking strategy
			i = 0
			for k, _ := range randomSequence {
				if targetNonce == k {
					//We found target nonce for this iteration
					totalPercentPicking = totalPercentPicking + (float32(i) / float32(nonceSize) * 100)
					break
				} else {

					i++
				}
			}
		}

		fmt.Println(
			nonceSize, "|",
			totalPercentGoOver/float32(ITERATIONS_NUMBER), "|",
			totalPercentPicking/float32(ITERATIONS_NUMBER), "|",
		)
	}
}

// Generate random numbers sequence with constantly low access time
func randomSequenceGenerator(size int32) map[int32]bool {
	var num int32 = 0
	var total int32 = 0
	var result map[int32]bool = make(map[int32]bool)

	for {
		num = rand.Int31n(size)
		if _, ok := result[num]; !ok {
			result[num] = true
			total++
			if total == size {
				break
			}
		}
	}

	return result
}
