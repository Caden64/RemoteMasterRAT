//go:build windows

package windows

import (
	"github.com/gonutz/w32/v2"
	"os/exec"
	"strconv"
)

func BuildNumberWindows() (string, error) {
	v := w32.RtlGetVersion()
	return strconv.Itoa(int(v.BuildNumber)), nil
}

func GetShellWindows() string {
	exec.Command("(dir 2>&1 *`|echo CMD);&<# rem #>echo ($PSVersionTable).PSEdition")
}
