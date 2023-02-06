package main

import (
	"bytes"
	"fmt"
	"github.com/gonutz/w32/v2"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	os := runtime.GOOS
	fmt.Println(os)
	switch os {

	case "windows":
		v := w32.RtlGetVersion()
		fmt.Println(v.BuildNumber)
	case "darwin":
		cmd := exec.Command("sw_vers")
		cmd.Stdin = strings.NewReader("Version input")
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

}
