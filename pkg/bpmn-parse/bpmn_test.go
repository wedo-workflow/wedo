package bpmn_parse

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/wedo-workflow/wedo/pkg/bpmn-parse/semantic"
)

func TestXML(t *testing.T) {
	f, err := os.ReadFile("testdata/shenpi.bpmn")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(Parse(f))
}

func TestXML_Process(t *testing.T) {
	f, err := os.ReadFile("testdata/process.bpmn")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	dec := xml.NewDecoder(bytes.NewReader(f))
	process := &semantic.TProcess{}
	if err := process.UnmarshalXML(dec, xml.StartElement{
		Name: xml.Name{},
		Attr: nil,
	}); err != nil {
		log.Println("process Unmarshal err", err)

	}
	pb, _ := json.Marshal(process)
	fmt.Println(string(pb))

	xml.Unmarshal(f, process)

	pb, _ = json.Marshal(process)
	fmt.Println(string(pb))
}
