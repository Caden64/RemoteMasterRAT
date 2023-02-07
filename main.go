package main

import (
	"RemoteMasterRAT/common"
	"RemoteMasterRAT/darwin"
	"RemoteMasterRAT/windows"
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	switch os {

	case "windows":
		build, _ := windows.BuildNumberWindows()
		fmt.Println(build)
		shell := windows.GetShellWindows()
		fmt.Println(shell)
		common.PrintProcessName()
	case "darwin":
		build, _ := darwin.BuildNumberDarwin()
		fmt.Println(build)
		shell := darwin.GetShellDarwin()
		fmt.Println(shell)
		common.PrintProcessName()
	}
}
