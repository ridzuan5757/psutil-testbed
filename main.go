package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	hddMetrics, err := GetMountMetrics()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, mm := range hddMetrics {

		ppName, _ := json.MarshalIndent(mm, "", " ")
		fmt.Println(string(ppName))

	}
}
