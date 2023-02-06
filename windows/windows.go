//go:build windows

package windows

import (
	"github.com/gonutz/w32/v2"
	"strconv"
)

func BuildNumberWindows() (string, error) {
	v := w32.RtlGetVersion()
	return strconv.Itoa(int(v.BuildNumber)), nil
}
