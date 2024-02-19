package main

import (
	"testing"

	"github.com/shirou/gopsutil/v3/disk"
)

func TestGetHddMetrics(t *testing.T) {
	// Run the function
	metrics, err := GetHddMetrics()

	// Check if there was an error
	if err != nil {
		t.Errorf("Error fetching HDD metrics: %v", err)
	}

	// Check if the metrics are reasonable (assuming we have at least one disk)
	if metrics.Total == 0 {
		t.Errorf("Unexpected total disk space: %d", metrics.Total)
	}
}

func BenchmarkGetHddMetrics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GetHddMetrics()
		if err != nil {
			b.Errorf("Error fetching HDD metrics: %v", err)
		}
	}
}

func TestGetMountMetrics(t *testing.T) {

	partitions := []disk.PartitionStat{
		{
			Device:     "foo",
			Mountpoint: "/mnt/drive1",
			Fstype:     "foo",
			Opts:       []string{"foo"},
		},
		{
			Device:     "foo",
			Mountpoint: "/mnt/drive2",
			Fstype:     "foo",
			Opts:       []string{"foo"},
		},
	}

	usage := map[string]*disk.UsageStat{
		"mnt/drive1": {
			Total:       100,
			Used:        50,
			Free:        50,
			UsedPercent: 0.5,
		},
		"mnt/drive2": {
			Total:       100,
			Used:        50,
			Free:        50,
			UsedPercent: 0.5,
		},
	}

	disk.Partitions = func(all bool) ([]disk.PartitionStat, error) {
		return partitions, nil
	}

	disk.Usage = func(path string) *disk.UsageStat {
		return usage[path], nil
	}

	mountMetrics, err := GetMountMetrics()

	if err != nil {
		t.Errorf("Expected 2 mount metrics, got %d", len(mountMetrics))
	}
}
