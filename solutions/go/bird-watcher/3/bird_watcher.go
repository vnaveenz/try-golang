package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var tc int = 0
	for i := 0; i < len(birdsPerDay) ; i ++ {
		tc = tc + birdsPerDay[i]
	}
	return tc
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	wbc := 0
	for i := (week - 1) * 7 ; i < week * 7 ; i ++ {
		wbc = wbc + birdsPerDay[i]
	}
	return wbc
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0 ; i < len(birdsPerDay); i +=2 {
		birdsPerDay[i]++
	}
	return birdsPerDay
}
