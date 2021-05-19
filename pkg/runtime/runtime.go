package runtime

import (
	"log"

	"github.com/wedo-workflow/wedo/configs"
	"github.com/wedo-workflow/wedo/pkg/bpmn"
)

type WedoRuntime struct {
	bpmn *bpmn.Wedo
}

func NewWedoRuntime(config *configs.Config) *WedoRuntime {
	b, err := bpmn.NewWedo(config)
	if err != nil {
		log.Fatal(err)
	}
	wr := &WedoRuntime{
		bpmn: b,
	}

	return wr
}

func (wr *WedoRuntime) Run(opts ...Option) error {

	return nil
}

func (wr *WedoRuntime) Deploy(xml []byte) error {
	err := wr.bpmn.ParseDoc(xml)
	return err
}
