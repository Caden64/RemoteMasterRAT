//go:build !darwin

package darwin

func BuildNumberDarwin() (string, error) {
	return "", nil
}
