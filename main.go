package main

import (
	"RemoteMasterRAT/darwin"
	"RemoteMasterRAT/windows"
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	fmt.Println(os)
	switch os {

	case "windows":
		build, _ := windows.BuildNumberWindows()
		fmt.Println(build)
	case "darwin":
		build, _ := darwin.BuildNumberDarwin()
		fmt.Println(build)

	}
}
