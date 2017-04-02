package pipeline

// Pipeline is structure
type Pipeline struct {
}

// New for pipeline
func New() *Pipeline {
	return &Pipeline{}
}

// Run invokes the CLI with the given arguments.
func (p *Pipeline) Run(args []string) int {
	return 0
}
