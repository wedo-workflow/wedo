package runtime

type (
	// runtimeOpts encapsulates the components to include in the runtime.
	runtimeOpts struct {
	}

	// Option is a function that customizes the runtime.
	Option func(o *runtimeOpts)
)
