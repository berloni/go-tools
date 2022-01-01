package tools

// StringIsInMap checks if a string is in a given map
func StringIsInMap(lookup string, list map[string]bool) bool {
	if _, ok := list[lookup]; ok {
		return true
	}

	return false
}

// StringIsInSlice checks if a string is in a given slice
func StringIsInSlice(lookup string, list []string) bool {
	for _, val := range list {
		if val == lookup {
			return true
		}
	}

	return false
}
