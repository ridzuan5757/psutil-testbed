package main

import (
	"fmt"
	"sync"

	"github.com/shirou/gopsutil/v3/disk"
)

type HddMetrics struct {
	Total       uint64
	Free        uint64
	Used        uint64
	UsedPercent float64
}

func GetHddMetrics() (HddMetrics, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return HddMetrics{}, err
	}

	var (
		hddMetrics HddMetrics
		wg         sync.WaitGroup
		mu         sync.Mutex
		errors     []error
	)

	for _, p := range partitions {
		wg.Add(1)

		go func(name string) {
			defer wg.Done()

			usage, err := disk.Usage(p.Mountpoint)

			if err != nil {
				mu.Lock()
				errors = append(errors, fmt.Errorf("error getting usage for partition %s: %v", name, err))
				mu.Unlock()
			}

			mu.Lock()
			hddMetrics.Free += usage.Free
			hddMetrics.Used += usage.Used
			hddMetrics.Total += usage.Total
			mu.Unlock()
		}(p.Mountpoint)

		wg.Wait()

		if len(errors) > 0 {
			var errMsg string
			for _, e := range errors {
				errMsg += e.Error() + "\n"
			}
			return HddMetrics{}, fmt.Errorf(
				"encountered errors while fetching disk metrics:\n%s",
				errMsg,
			)
		}

	}

	hddMetrics.UsedPercent = float64(hddMetrics.Used) / float64(hddMetrics.Total)
	return hddMetrics, nil
}
