package runtime

import "github.com/wedo-workflow/wedo/runtime/config"

func initRuntime() *Runtime {
	rt, err := NewRuntime(config.NewDefaultConfig())
	if err != nil {
		panic(err)
	}
	return rt
}
