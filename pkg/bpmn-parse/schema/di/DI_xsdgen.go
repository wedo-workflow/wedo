// Code generated by xsdgen. DO NOT EDIT.

package schema

import "encoding/xml"

type Bounds struct {
	X      float64 `xml:"x,attr"`
	Y      float64 `xml:"y,attr"`
	Width  float64 `xml:"width,attr"`
	Height float64 `xml:"height,attr"`
}

type Diagram struct {
	Name          string  `xml:"name,attr,omitempty"`
	Documentation string  `xml:"documentation,attr,omitempty"`
	Resolution    float64 `xml:"resolution,attr,omitempty"`
}

type DiagramElement struct {
	Extension Extension `xml:"http://www.omg.org/spec/DD/20100524/DI extension,omitempty"`
}

type Edge struct {
	Extension Extension `xml:"http://www.omg.org/spec/DD/20100524/DI extension,omitempty"`
	Waypoint  []Point   `xml:"http://www.omg.org/spec/DD/20100524/DC waypoint"`
}

type Extension []string

func (a Extension) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ArrayType string   `xml:"http://schemas.xmlsoap.org/wsdl/ arrayType,attr"`
		Items     []string `xml:" item"`
	}
	output.Items = []string(a)
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{"", "xmlns:ns1"}, Value: "http://www.w3.org/2001/XMLSchema"})
	output.ArrayType = "ns1:anyType[]"
	return e.EncodeElement(&output, start)
}
func (a *Extension) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var tok xml.Token
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if tok, ok := tok.(xml.StartElement); ok {
			var item string
			if err = d.DecodeElement(&item, &tok); err == nil {
				*a = append(*a, item)
			}
		}
		if _, ok := tok.(xml.EndElement); ok {
			break
		}
	}
	return err
}

type Font struct {
	Name            string  `xml:"name,attr,omitempty"`
	Size            float64 `xml:"size,attr,omitempty"`
	IsBold          bool    `xml:"isBold,attr,omitempty"`
	IsItalic        bool    `xml:"isItalic,attr,omitempty"`
	IsUnderline     bool    `xml:"isUnderline,attr,omitempty"`
	IsStrikeThrough bool    `xml:"isStrikeThrough,attr,omitempty"`
}

type Label struct {
	Extension Extension `xml:"http://www.omg.org/spec/DD/20100524/DI extension,omitempty"`
	Bounds    string    `xml:"http://www.omg.org/spec/DD/20100524/DI Bounds,omitempty"`
}

type LabeledEdge struct {
	Extension Extension `xml:"http://www.omg.org/spec/DD/20100524/DI extension,omitempty"`
	Waypoint  []Point   `xml:"http://www.omg.org/spec/DD/20100524/DC waypoint"`
}

type LabeledShape struct {
	Extension Extension `xml:"http://www.omg.org/spec/DD/20100524/DI extension,omitempty"`
	Bounds    string    `xml:"http://www.omg.org/spec/DD/20100524/DI Bounds"`
}

type Node struct {
	Extension Extension `xml:"http://www.omg.org/spec/DD/20100524/DI extension,omitempty"`
}

type Plane struct {
	Extension      Extension `xml:"http://www.omg.org/spec/DD/20100524/DI extension,omitempty"`
	DiagramElement []string  `xml:"http://www.omg.org/spec/DD/20100524/DI DiagramElement,omitempty"`
}

type Point struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

type Shape struct {
	Extension Extension `xml:"http://www.omg.org/spec/DD/20100524/DI extension,omitempty"`
	Bounds    string    `xml:"http://www.omg.org/spec/DD/20100524/DI Bounds"`
}

type Style struct {
}
