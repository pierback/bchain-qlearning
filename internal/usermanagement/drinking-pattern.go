package usermanagement

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

func calcDrinkPerTimeslot(desiredStdDev, desiredMean float64) float64 {
	randfloat := rand.NormFloat64()
	// randfloat := test.Rand()
	normal := math.Abs(randfloat)*desiredStdDev + desiredMean

	// fmt.Printf("NormFloat64 %f random %f \n", normal, randfloat)

	return normal
}

//CreateDrinkinTimes creates drinking times for one day all timeslots
func CreateDrinkinTimes() []float64 {
	desiredStdDev := 1.0
	desiredMean := 3.0
	randfloat := rand.NormFloat64()
	ccPD := randfloat*desiredStdDev + desiredMean
	// fmt.Println("ccPD", math.Floor(ccPD))
	return randFloats(int(math.Round(ccPD)))
}

//randFloats max is not included
func randFloats(n int) []float64 {
	var allValues, drinkinTimes []float64
	//				mean, dev
	ts0 := []float64{1.0, 1.0}
	ts1 := []float64{1.5, 0.2}
	ts2 := []float64{1.5, 0.01}
	ts3 := []float64{0.2, 0.5}
	ts4 := []float64{0.0, 0.0}

	tsts := [][]float64{ts0, ts1, ts2, ts3, ts4}

	var sum int

	for i := 0; i < len(tsts); i++ {
		randfloat := rand.NormFloat64()
		desiredMean := tsts[i][0]
		desiredStdDev := tsts[i][1]
		// randfloat := test.Rand()
		normal := math.Abs(randfloat)*desiredStdDev + desiredMean
		allValues = append(allValues, math.Floor(normal))
		sum += int(math.Floor(normal))
	}

	// fmt.Printf("new values %v sum %d \n", allValues, sum)
	// var finalSum int
	for k, val := range allValues {
		desiredMean := tsts[k][0]
		if val > desiredMean+1 && sum > n {
			allValues[k] = 0.0
		}
		// finalSum += int(math.Floor(allValues[k]))
		for cc := 0; cc < int(math.Floor(allValues[k])); cc++ {
			drinkinTimes = append(drinkinTimes, randTime(k))
		}
		_ = desiredMean
		_ = val
	}

	// fmt.Printf("new values %v \n", allValues)
	// fmt.Printf("new values %v drinkinTimes \n", drinkinTimes)
	return drinkinTimes
}

func randInts(min, max float64) int {
	res := min + rand.Float64()*(max-min)
	return int(math.Round((res)))
}

func randTime(ts int) float64 {
	upperBoundary := 6 + 3*ts
	lowerBoundary := 9 + 3*ts
	min := time.Date(2019, 1, 0, upperBoundary, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2019, 1, 0, lowerBoundary, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	hour := time.Unix(sec, 0).Hour()
	minute := time.Unix(sec, 0).Minute()
	clock := strconv.Itoa(hour) + "." + strconv.Itoa(minute)
	tmStr, _ := strconv.ParseFloat(clock, 64)
	// fmt.Printf("clock: %s tmStr %0.2f \n", clock, tmStr)
	return tmStr
}
