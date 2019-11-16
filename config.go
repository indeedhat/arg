package arg

type ValidatorFunc func(arg string) error

type ArgConfig struct {
	Default  []string
	Usage    string
	Required bool
	Override bool
	Validate ValidatorFunc
}

type CommandConfig struct {
	Title string
	Usage string
}
