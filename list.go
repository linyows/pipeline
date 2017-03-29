package main

var cmdList = &Command{
	Run:       runList,
	UsageLine: "list ",
	Short:     "",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdList.Flag.BoolVar(&flagA, "a", false, "")
}

// runList executes list command and return exit code.
func runList(args []string) int {

	return 0
}
