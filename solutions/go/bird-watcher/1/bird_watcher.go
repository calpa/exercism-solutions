package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	ans := 0
	for i := 0; i < len(birdsPerDay); i++ {
        ans += birdsPerDay[i]
    }

    return ans
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekStart := (week - 1) * 7
    weekEnd := weekStart + 7
    total := 0
    for i := weekStart; i < weekEnd; i++ {
        total += birdsPerDay[i]
    }

    return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
        birdsPerDay[i] += 1
    }

    return birdsPerDay
}
