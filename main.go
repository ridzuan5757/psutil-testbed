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

	ppName, _ := json.MarshalIndent(hddMetrics, "", " ")
	fmt.Println(string(ppName))
}
