package utils

func BytesToKB(b uint64) float64 {
	return float64(b) / 1024
}

func BytesToMB(b uint64) float64 {
	return float64(b) / (1024 * 1024)
}
