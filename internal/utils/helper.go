package utils

func InSlicesInt64(item int64, slice []int64) bool {
	for _, i := range slice{
		if i == item {
			return true
		}
	}
	return false
}