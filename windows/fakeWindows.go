//go:build !windows

package windows

func BuildNumberWindows() (string, error) {
	return "", nil
}

func GetShellWindows() string {
	return ""
}
