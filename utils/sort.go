package utils

// Check if source array contains destination element.
func Contains(source []int, destination int) bool {
	for _, value := range source {
		if value == destination {
			return true
		}
	}
	return false
}
