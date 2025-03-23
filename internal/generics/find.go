package generics

func Find[T comparable](slice []T, needle T) (int, bool) {
	for i, v := range slice {
		if v == needle {
			return i, true
		}
	}
	return -1, false
}
