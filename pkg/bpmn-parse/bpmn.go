package bpmn_parse

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"

	"github.com/wedo-workflow/wedo/pkg/bpmn-parse/semantic"
	"github.com/wedo-workflow/wedo/pkg/bpmn-parse/xmltree"
)

func Parse(doc []byte) (interface{}, error) {
	tree, err := xmltree.Parse(doc)
	if err != nil {
		return nil, err
	}

	tree.SearchFunc(func(element *xmltree.Element) bool {
		parse(element)
		return true
	})

	return nil, nil
}

func parse(element *xmltree.Element) {
	switch element.Name.Local {
	case BPMN_ELEMENT_PROCESS:
		parseProcess(element)
	case BPMN_ELEMENT_LANE_SET:
		// parseLaneSet(dec, element)
	}
}

func parseProcess(element *xmltree.Element) {
	attrMap := map[string]interface{}{}
	for _, attr := range element.StartElement.Attr {
		fmt.Println(attr.Name.Local, attr.Value)
		if _, ok := attrMap[attr.Name.Local]; !ok {
			if attr.Value == "true" {
				attrMap[attr.Name.Local] = true
			} else if attr.Value == "false" {
				attrMap[attr.Name.Local] = false
			} else {
				attrMap[attr.Name.Local] = attr.Value
			}
		}
	}
	ab, err := json.Marshal(attrMap)
	if err != nil {
		log.Println(err)
		return
	}
	process := &semantic.TProcess{}

	if err := json.Unmarshal(ab, process); err != nil {
		log.Println("xml.Unmarshal err", err)
		return
	}
	pb, _ := json.Marshal(process)
	fmt.Println(string(pb))
}

func parseLaneSet(dec *xml.Decoder, element *xmltree.Element) {
	fmt.Println(element.Name, element.String())
	for _, child := range element.Children {
		fmt.Println(child.Name)
	}
}
