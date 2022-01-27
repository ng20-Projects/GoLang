package main

import "fmt"

func SuccessRate(workSpeed int) float64 {
	if workSpeed == 0 {
		return 0.0
	} else if workSpeed >= 1 && workSpeed <= 4 {
		return 1.0
	} else if workSpeed >= 5 && workSpeed <= 8 {
		return 0.9
	} else if workSpeed >= 9 && workSpeed <= 10 {
		return 0.77
	} else {
		return -1.0
	}
}

func CalculateProductionRatePerHour(workSpeed int) float64 {
	return float64(workSpeed * 221)
}

func CalculateProductionRatePerMinute(workSpeed int) int {
	return (workSpeed * 221) / 60
}

func CalculateLimitedProductionRatePerHour(workSpeed int, limitProduction float64) float64 {
	production := float64(workSpeed * 221)
	if production <= limitProduction {
		return production
	} else {
		return limitProduction
	}
}

func main() {
	var workSpeed int
	fmt.Print("Work Speed: ")
	fmt.Scanln(&workSpeed)

	var limitProduction float64
	fmt.Print("\nLimit Production Per Hour: ")
	fmt.Scanln(&limitProduction)

	SR := SuccessRate(workSpeed)
	CPRPH := CalculateProductionRatePerHour(workSpeed) * SR
	CPRPM := int(float64(CalculateProductionRatePerMinute(workSpeed)) * SR)
	CLPRPH := CalculateLimitedProductionRatePerHour(workSpeed, limitProduction)

	fmt.Printf("\n Success Rate is %v percent \n Prod. Rate per hour is %v \n Prod. Rate per minute is %v \n Limited Production per hour is %v", SR*100, CPRPH, CPRPM, CLPRPH)
	fmt.Println("")
}
