package features

type VolumeMA struct {
	Period int
	Sequential []float64
}

func (ma *VolumeMA) Add(value float64) {
	ma.Sequential = append(ma.Sequential, value)

	if len(ma.Sequential) > ma.Period {
		ma.Sequential = ma.Sequential[1:]
	}
}

func (ma *VolumeMA) Value() float64 {
	var sum float64

	for _, value := range ma.Sequential {
		sum += value
	}

	return sum / float64(len(ma.Sequential))
}