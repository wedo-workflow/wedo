package xmltree

import (
	"encoding/xml"
	"fmt"
)

// An Element represents a single element in an XML document. Elements
// may have zero or more children. The byte array used by the Content
// field is shared among all elements in the document, and should not
// be modified. An Element also captures xml namespace prefixes, so
// that arbitrary QNames in attribute values can be resolved.
type Element struct {
	xml.StartElement
	// The XML namespace scope at this element's location in the
	// document.
	Scope
	// The raw content contained within this element's start and
	// end tags. Uses the underlying byte array passed to Parse.
	Content []byte
	// Sub-elements contained within this element.
	Children []Element
}

// Attr gets the value of the first attribute whose name matches the
// space and local arguments. If space is the empty string, only
// attributes' local names are considered when looking for a match.
// If an attribute could not be found, the empty string is returned.
func (ele *Element) Attr(space, local string) string {
	for _, v := range ele.StartElement.Attr {
		if v.Name.Local != local {
			continue
		}
		if space == "" || space == v.Name.Space {
			return v.Value
		}
	}
	return ""
}

func (ele *Element) parse(scanner *scanner, data []byte, depth int) error {
	if depth > recursionLimit {
		return errDeepXML
	}
	ele.StartElement.Attr = ele.pushNS(ele.StartElement)

	begin := scanner.InputOffset()
	end := begin
walk:
	for scanner.scan() {
		switch tok := scanner.tok.(type) {
		case xml.StartElement:
			child := Element{StartElement: tok.Copy(), Scope: ele.Scope}
			if err := child.parse(scanner, data, depth+1); err != nil {
				return err
			}
			ele.Children = append(ele.Children, child)
		case xml.EndElement:
			if tok.Name != ele.Name {
				return fmt.Errorf("Expecting </%s>, got </%s>", ele.Prefix(ele.Name), ele.Prefix(tok.Name))
			}
			ele.Content = data[int(begin):int(end)]
			break walk
		}
		end = scanner.InputOffset()
	}
	return scanner.err
}

// Flatten produces a slice of Element pointers referring to
// the children of el, and their children, in depth-first order.
func (ele *Element) Flatten() []*Element {
	return ele.SearchFunc(func(*Element) bool { return true })
}

// SetAttr adds an XML attribute to an Element's existing Attributes.
// If the attribute already exists, it is replaced.
func (ele *Element) SetAttr(space, local, value string) {
	for i, a := range ele.StartElement.Attr {
		if a.Name.Local != local {
			continue
		}
		if space == "" || a.Name.Space == space {
			ele.StartElement.Attr[i].Value = value
			return
		}
	}
	ele.StartElement.Attr = append(ele.StartElement.Attr, xml.Attr{
		Name:  xml.Name{Space: space, Local: local},
		Value: value,
	})
}

// Search searches the Element tree for Elements with an xml tag
// matching the name and xml namespace. If space is the empty string,
// any namespace is matched.
func (ele *Element) Search(space, local string) []*Element {
	return ele.SearchFunc(func(el *Element) bool {
		if local != el.Name.Local {
			return false
		}
		return space == "" || space == el.Name.Space
	})
}

// SearchFunc traverses the Element tree in depth-first order and returns
// a slice of Elements for which the function fn returns true.
func (ele *Element) SearchFunc(fn func(*Element) bool) []*Element {
	var results []*Element
	var search func(el *Element)

	search = func(el *Element) {
		if fn(el) {
			results = append(results, el)
		}
		el.walk(search)
	}
	ele.walk(search)
	return results
}

// walkFunc is the type of the function called for each of an Element's
// children.
type walkFunc func(*Element)

// The walk method calls the walkFunc for each of the Element's children.
// If the WalkFunc returns a non-nil error, Walk will return it
// immediately.
func (ele *Element) walk(fn walkFunc) error {
	for i := 0; i < len(ele.Children); i++ {
		fn(&ele.Children[i])
	}
	return nil
}

type byXMLName []xml.Name

func (x byXMLName) Len() int { return len(x) }
func (x byXMLName) Less(i, j int) bool {
	return x[i].Space+x[i].Local < x[j].Space+x[j].Local
}
func (x byXMLName) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// String returns the XML encoding of an Element
// and its children as a string.
func (ele *Element) String() string {
	return string(Marshal(ele))
}
