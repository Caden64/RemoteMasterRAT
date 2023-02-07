//go:build darwin

package darwin

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

func BuildNumberDarwin() (ProductInfo, error) {
	cmd := exec.Command("sw_vers")
	cmd.Stdin = strings.NewReader("Version data")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return ProductInfo{}, err
	}

	data := string(out.Bytes())
	lines := strings.Split(data, "\n")

	var info ProductInfo
	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], "ProductName:") {
			pns := strings.Split(lines[i], ":")
			pnss := strings.TrimSpace(pns[1])
			info.Name = pnss
		} else if strings.Contains(lines[i], "ProductVersion:") {
			pvs := strings.Split(lines[i], ":")
			pvss := strings.TrimSpace(pvs[1])
			info.Version = pvss
		} else if strings.Contains(lines[i], "BuildVersion:") {
			bvs := strings.Split(lines[i], ":")
			bvss := strings.TrimSpace(bvs[1])
			info.Build = bvss
		}
	}
	return info, nil
}

func GetShellDarwin() string {

	return os.ExpandEnv("$SHELL")

}
