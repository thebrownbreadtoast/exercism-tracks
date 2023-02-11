package cars


const SingleCarCost = 10000

const TenCarBatchCost = 95000


// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	productionRateFloat := float64(productionRate)

	return productionRateFloat * (successRate / 100.0)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	successfulCarsCreatedPerHour := int(CalculateWorkingCarsPerHour(productionRate, successRate))

	return successfulCarsCreatedPerHour / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	singleCars := carsCount % 10

	carBatches := carsCount / 10

	singleCarsCost := singleCars * SingleCarCost
	batchedCarsCost := carBatches * TenCarBatchCost

	return uint(singleCarsCost + batchedCarsCost)
}
