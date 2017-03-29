package main

var cmdInit = &Command{
	Run:       runInit,
	UsageLine: "init ",
	Short:     "",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdInit.Flag.BoolVar(&flagA, "a", false, "")
}

// runInit executes init command and return exit code.
func runInit(args []string) int {

	return 0
}
