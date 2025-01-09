package ans
func avg (data [] float64) float64 {
	total := float64(0)
	for _, value := range data {
		total += value
	}
	return total / float64(len(data))
}