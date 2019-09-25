package arg

type ValidatorFunc func(arg string) error

type ArgConfig struct {
	Default  interface{}
	Required bool
	Override bool
	Usage    string
	Validate ValidatorFunc
}
