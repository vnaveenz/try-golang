package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	succCars := (successRate/float64(100)) * float64(productionRate)
	return succCars
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	prodRatePerMinute := float64(productionRate / 60)
	successRatePerMinute := prodRatePerMinute * (successRate/100)
	return int(successRatePerMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	effBatch := carsCount/10 * 95000
	expBatch := carsCount%10 * 10000
	totalCost := uint(effBatch + expBatch)
	return  totalCost
}
