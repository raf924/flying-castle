package utils

func Reverse(array []byte) []byte {
	newArray := make([]byte, len(array))
	for i, j := 0, len(array)-1; i <= j; i, j = i+1, j-1 {
		newArray[i], newArray[j] = array[j], array[i]
	}
	return newArray
}

func Flatten(array [][]byte) []byte {
	var result = make([]byte, 0)
	for _, d := range array {
		result = append(result, d...)
	}
	return result
}
