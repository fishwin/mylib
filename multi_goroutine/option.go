package multi_goroutine

type OptionFunc func(option *Option)

type Option struct {
	goroutinePrefixName string
	submitStrategy      SubmitStrategyType
	errChan             chan error
}

func WithGoroutinePrefixName(name string) OptionFunc {
	return func(option *Option) {
		option.goroutinePrefixName = name
	}
}

func WithSubmitStrategy(strategy SubmitStrategyType) OptionFunc {
	return func(option *Option) {
		option.submitStrategy = strategy
	}
}

func WithErrChan(errChan chan error) OptionFunc {
	return func(option *Option) {
		option.errChan = errChan
	}
}
