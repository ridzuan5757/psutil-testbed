package hddmetrics

import (
	"github.com/shirou/gopsutil/v3/disk"
)

type HddMetrics struct {
	Total       uint64
	Free        uint64
	Used        uint64
	UsedPercent float64
}

func GetHddMetrics() (HddMetrics, error) {
	partition, err := disk.Partitions(false)
	if err != nil {
		return HddMetrics{}, err
	}

	var pName []string
	var hddMetrics HddMetrics

	for _, p := range partition {
		pName = append(pName, p.Mountpoint)
	}

	for _, name := range pName {
		usage, err := disk.Usage(name)
		if err != nil {
			return HddMetrics{}, err
		}

		hddMetrics.Free += usage.Free
		hddMetrics.Used += usage.Used
		hddMetrics.Total += usage.Total

	}

	hddMetrics.UsedPercent = float64(hddMetrics.Used) / float64(hddMetrics.Total)

	return hddMetrics, nil
}
