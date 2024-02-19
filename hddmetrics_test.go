package hddmetrics

import (
	"testing"
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
