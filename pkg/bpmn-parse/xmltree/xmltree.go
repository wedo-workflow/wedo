package xmltree

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"

	"golang.org/x/net/html/charset"
)

const (
	xmlNamespaceURI = "http://www.w3.org/2000/xmlns/"
	xmlLangURI      = "http://www.w3.org/XML/1998/namespace"
	recursionLimit  = 3000
)

var errDeepXML = errors.New("xmltree: xml document too deeply nested")

// Unmarshal parses the XML encoding of the Element and stores the result
// in the value pointed to by v. Unmarshal follows the same rules as
// xml.Unmarshal, but only parses the portion of the XML document
// contained by the Element.
func Unmarshal(el *Element, v interface{}) error {
	return xml.Unmarshal(Marshal(el), v)
}

// Save some typing when scanning xml
type scanner struct {
	*xml.Decoder
	tok xml.Token
	err error
}

func (s *scanner) scan() bool {
	if s.err != nil {
		return false
	}
	s.tok, s.err = s.Token()
	return s.err == nil
}

// Parse builds a tree of Elements by reading an XML document.  The
// byte slice passed to Parse is expected to be a valid XML document
// with a single root element.
func Parse(doc []byte) (*Element, error) {
	d := xml.NewDecoder(bytes.NewReader(doc))

	// The xmltree package, when constructing the tree, takes slices
	// of the source document for chardata (data between tags). To do
	// this, it takes the position of the Decoder in the utf-8 input
	// stream. If the source document is not utf8, the position may be
	// incorrect and cause invalid data or a run-time panic. So we copy
	// the utf8 conversion to an internal buffer.
	utf8buf := bytes.NewBuffer(doc[:0])
	d.CharsetReader = func(label string, r io.Reader) (io.Reader, error) {
		utf8input, err := charset.NewReaderLabel(label, r)
		if err != nil {
			return nil, err
		}
		// At this point, the encoding/xml package has already
		// parsed the <?xml?> header. To be able to index
		// into the document, we need to account for this.
		padding := make([]byte, int(d.InputOffset()))
		utf8buf.Write(padding)

		_, err = io.Copy(utf8buf, utf8input)
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(utf8buf.Bytes()[len(padding)+1:]), nil
	}
	scanner := scanner{Decoder: d}
	root := new(Element)

	for scanner.scan() {
		if start, ok := scanner.tok.(xml.StartElement); ok {
			root.StartElement = start
			break
		}
	}
	if scanner.err != nil {
		return nil, scanner.err
	}
	if err := root.parse(&scanner, utf8buf.Bytes(), 0); err != nil {
		return nil, err
	}
	return root, nil
}
