package learning

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	mondayTimes := []float64{8.33, 10, 15}
	tuesdayTimes := []float64{8.49, 10.30, 12.30}
	wednesdayTimes := []float64{8.15, 10, 14.37}
	thursdayTimes := []float64{8.30, 9.30, 13, 15.20}
	fridayTimes := []float64{8.37, 10.15, 13.23, 15.57}

	wts := [][]float64{mondayTimes, tuesdayTimes, wednesdayTimes, thursdayTimes, fridayTimes}
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("NormFloat64 %f \n", rand.NormFloat64())
	fmt.Printf("NormFloat64 %f \n", rand.NormFloat64())

	// rand.Seed(123455)
	rand.Seed(time.Now().UnixNano())

	r := randInts(0, 3.0)
	fmt.Printf("\nRandom number %d \n", r)
	fmt.Printf("Random number %d \n", r)

	r1 := randInts(0, 3.0)
	fmt.Printf("\nRandom number %d \n", r1)
	fmt.Printf("Random number %d \n\n", r1)

	// distribution := gaussian.NewGausian(mean, variance)

	var s string
	s = fmt.Sprintf("%.2f", randFloats(7.0, 10.0, r))
	fmt.Printf(s)

	fmt.Println("  ")

	rand.Seed(123456)
	s = fmt.Sprintf("%.2f", randFloats(7.0, 10.0, r1))
	fmt.Printf(s)

	fmt.Println("  ")

	_ = wts
}

//randFloats max is not included
func randFloats(min, max float64, n int) []float64 {
	var allValues []float64

	for i := 0; i < n; i++ {
		tmp := math.Round((min+rand.Float64()*(max-min))*100) / 100
		fmt.Printf("tmp %v \n", tmp)
		if i == 0 {
			allValues = append(allValues, tmp)
		} else if i > 0 {
			if allValues[i-1] < tmp {
				allValues = append(allValues, tmp)
			}
		}
	}

	fmt.Printf("new values %v \n", allValues)
	return allValues
}

func randInts(min, max float64) int {
	res := min + rand.Float64()*(max-min)
	return int(math.Round((res)))
}
