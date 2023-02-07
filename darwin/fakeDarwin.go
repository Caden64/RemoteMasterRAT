//go:build !darwin

package darwin

func BuildNumberDarwin() (ProductInfo, error) {
	return ProductInfo{}, nil
}

func GetShellDarwin() string {
	return ""
}
