//go:build windows

package windows

import (
	"fmt"
	"github.com/gonutz/w32/v2"
	"os/exec"
	"strconv"
	"strings"
)

func BuildNumberWindows() (string, error) {
	v := w32.RtlGetVersion()
	return strconv.Itoa(int(v.BuildNumber)), nil
}

func GetShellWindows() string {
	cmd := exec.Command("powershell", "(dir 2>&1 *`|echo CMD);&<# rem #>echo ($PSVersionTable).PSEdition")

	cmd.Stdin = strings.NewReader("Version data")

	data, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(data)
}
