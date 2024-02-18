package main

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
)

func main() {
	diskUsage, _ := disk.Usage("/")
	ppdp, _ := json.MarshalIndent(diskUsage, "", " ")
	fmt.Println(string(ppdp))
}
