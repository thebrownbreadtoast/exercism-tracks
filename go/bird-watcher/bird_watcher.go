package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalBirds := 0

	for idx := 0; idx < len(birdsPerDay); idx++ {
		totalBirds += birdsPerDay[idx]
	}

	return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekStart := (week - 1) * 7
	weekEnd := weekStart + 7

	slicedWeek := birdsPerDay[weekStart:weekEnd]

	totalBirdsInWeek := 0

	for idx := 0; idx < len(slicedWeek); idx++ {
		totalBirdsInWeek += slicedWeek[idx]
	}

	return totalBirdsInWeek
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for idx := 0; idx < len(birdsPerDay); idx++ {
		if idx % 2 == 0 {
			birdsPerDay[idx] += 1
		}
	}

	return birdsPerDay
}
