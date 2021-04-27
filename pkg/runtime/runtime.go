package runtime

type WedoRuntime struct {
}

func NewWedoRuntime() *WedoRuntime {
	return &WedoRuntime{}
}

func (wr *WedoRuntime) Run(opts ...Option) error {
	return nil
}
