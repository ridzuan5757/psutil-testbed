package main

import (
	//	"encoding/json"
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
)

type HddMetrics struct {
	Total       uint64
	Free        uint64
	Used        uint64
	UsedPercent float64
}

func main() {
	partition, _ := disk.Partitions(false)
	var pName []string
	var hddMetrics HddMetrics

	for _, p := range partition {
		pName = append(pName, p.Mountpoint)
	}

	ppName, _ := json.MarshalIndent(pName, "", " ")
	fmt.Println(string(ppName))

	for _, name := range pName {
		usage, err := disk.Usage(name)

		if err != nil {
			fmt.Println("Error", err)
		}

		hddMetrics.Free += usage.Free
		hddMetrics.Used += usage.Used
		hddMetrics.Total += usage.Total

	}

	hddMetrics.UsedPercent = float64(hddMetrics.Used) / float64(hddMetrics.Total)

	fmt.Println(hddMetrics)

}
