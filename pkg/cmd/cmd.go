package cmd

func InputPath(base, variant string) string {
	if variant == "" {
		return base + ".txt"
	}
	return base + "_" + variant + ".txt"
}
