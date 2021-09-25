package utils

func SearchString(arr []string, key string) bool {
	for _, item := range arr {
		if item == key {
			return true
		}
	}
	return false
}
