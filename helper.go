package filters

func Normalize(arrPtr *[]float64) {
	arr := *arrPtr
	sum := 0.0
	for _, val := range arr {
		sum += val
	}
	for i, val := range arr {
		arr[i] = val / sum
	}
}
