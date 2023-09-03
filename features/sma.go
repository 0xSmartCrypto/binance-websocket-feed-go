package features

import "fmt"

type SMA struct {
	Period int
	Sequential []float64
}

func (ma *SMA) Add(value float64) {
	ma.Sequential = append(ma.Sequential, value)

	if len(ma.Sequential) > ma.Period {
		ma.Sequential = ma.Sequential[1:]
	}
	fmt.Println(ma.Sequential)
}

func (ma *SMA) Value() float64 {
	var sum float64

	for _, value := range ma.Sequential {
		sum += value
	}

	return sum / float64(len(ma.Sequential))
}