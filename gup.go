package gup

// GetPkg takes the location of an executable and produces the
// import name.
func GetPkg(bin string) (string, error) {
	path, err := getExecPath(bin)
	if err != nil {
		return "", err
	}

	return getMainPath(path)
}
