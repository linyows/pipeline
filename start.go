package main

var cmdStart = &Command{
	Run:       runStart,
	UsageLine: "start ",
	Short:     "",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdStart.Flag.BoolVar(&flagA, "a", false, "")
}

// runStart executes start command and return exit code.
func runStart(args []string) int {

	return 0
}
