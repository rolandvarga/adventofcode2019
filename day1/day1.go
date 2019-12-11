package main

import "fmt"

var (
	modules = []int{148319, 54894, 105685, 136247, 133339, 91401, 126939, 102276, 66395, 134572, 137988, 65709, 119909, 98439, 88926, 74563, 73275, 111063, 92623, 66649, 147991, 71108, 129082, 58677, 57884, 93284, 61455, 110362, 81726, 146604, 70037, 82328, 78802, 69496, 61390, 134525, 94583, 73669, 136417, 80424, 98700, 88578, 147217, 108332, 73965, 116009, 51599, 55617, 129014, 51962, 95443, 114358, 141826, 134605, 145837, 112074, 93422, 112897, 137077, 141584, 114605, 122589, 121933, 67088, 120788, 53216, 82633, 55215, 135617, 91439, 110237, 130445, 120865, 109484, 113596, 133240, 113525, 110473, 146059, 129811, 79370, 50142, 111149, 137107, 91647, 82978, 119456, 51924, 132979, 63215, 55209, 114474, 54585, 101761, 79878, 63987, 149324, 100155, 54601, 115686}
)

func FuelCounterUpper(mass int) int {
	return mass/3 - 2
}

func PartOne() int {
	sum := 0
	for _, m := range modules {
		sum += FuelCounterUpper(m)
	}
	return sum
}

func PartTwo() int {
	sum := 0

	for _, m := range modules {
		needMoreFuel := true
		totalFuel := 0
		mMass := m

		for needMoreFuel {
			fuel := FuelCounterUpper(mMass)

			if fuel > 0 {
				totalFuel += fuel
				mMass = fuel
				continue
			}
			needMoreFuel = false
		}
		sum += totalFuel
	}
	return sum
}

func DayOne() {
	sum1 := PartOne()
	fmt.Printf("Total Fuel requirement #1: '%d'\n", sum1)

	sum2 := PartTwo()
	fmt.Printf("Total Fuel requirement #2: '%d'\n", sum2)
}
