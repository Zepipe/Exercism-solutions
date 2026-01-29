package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var totalBirds int

    for _, val := range birdsPerDay {
        totalBirds += val
    }

    return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    var totalBirds, countWeekDays int
	var firstDayInWeek int = 7 * week - 7
	// To escape out of range mistake
    if firstDayInWeek >= len(birdsPerDay) || len(birdsPerDay) == 0 || week <= 0 {
        return 0
    }

    for _, val := range birdsPerDay[firstDayInWeek:] {
        totalBirds += val
        countWeekDays += 1
        
        if countWeekDays == 7 {
            break
        }
    }
    return totalBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    // Modify the original slice in place
    for i := 0; i < len(birdsPerDay); i += 2 {
        birdsPerDay[i] += 1
    }
    
    return birdsPerDay
}
