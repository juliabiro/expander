package cmd

func ParseInput(args []string) ([]string, error) {
	var ret []string
	for _, c := range args {
		ret = append(ret, c)
	}
	return ret, nil
}
