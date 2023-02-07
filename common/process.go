package common

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
)

func PrintProcessName() {
	data, _ := process.Processes()
	for proc := 0; proc < len(data); proc++ {
		fmt.Println(data[proc].Name())
	}
}
