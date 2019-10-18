package arg

type ValidatorFunc func(arg string) error

type ArgConfig struct {
	Default   []string
	Usage     string
	Required  bool
	Override  bool
	IsCommand bool
	Validate  ValidatorFunc
}
