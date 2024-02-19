package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	hddMetrics, err := GetHddMetrics()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ppName, _ := json.MarshalIndent(hddMetrics, "", " ")
	fmt.Println(string(ppName))
}
