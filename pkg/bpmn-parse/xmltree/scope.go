package xmltree

import (
	"encoding/xml"
	"sort"
	"strings"
)

// A Scope represents the xml namespace scope at a given position in
// the document.
type Scope struct {
	ns []xml.Name
}

// The JoinScope method joins two Scopes together. When resolving
// prefixes using the returned scope, the prefix list in the argument
// Scope is searched before that of the receiver Scope.
func (scope *Scope) JoinScope(inner *Scope) *Scope {
	return &Scope{append(scope.ns, inner.ns...)}
}

// Resolve translates an XML QName (namespace-prefixed string) to an
// xml.Name with a canonicalized namespace in its Space field.  This can
// be used when working with XSD documents, which put QNames in attribute
// values. If qname does not have a prefix, the default namespace is used.If
// a namespace prefix cannot be resolved, the returned value's Space field
// will be the unresolved prefix. Use the ResolveNS function to detect when
// a namespace prefix cannot be resolved.
func (scope *Scope) Resolve(qname string) xml.Name {
	name, _ := scope.ResolveNS(qname)
	return name
}

// The ResolveNS method is like Resolve, but returns false for its second
// return value if a namespace prefix cannot be resolved.
func (scope *Scope) ResolveNS(qname string) (xml.Name, bool) {
	var prefix, local string
	parts := strings.SplitN(qname, ":", 2)
	if len(parts) == 2 {
		prefix, local = parts[0], parts[1]
	} else {
		prefix, local = "", parts[0]
	}
	switch prefix {
	case "xml":
		return xml.Name{Space: xmlLangURI, Local: local}, true
	case "xmlns":
		return xml.Name{Space: xmlNamespaceURI, Local: local}, true
	}
	for i := len(scope.ns) - 1; i >= 0; i-- {
		if scope.ns[i].Local == prefix {
			return xml.Name{Space: scope.ns[i].Space, Local: local}, true
		}
	}
	return xml.Name{Space: prefix, Local: local}, false
}

// ResolveDefault is like Resolve, but allows for the default namespace to
// be overridden. The namespace of strings without a namespace prefix
// (known as an NCName in XML terminology) will be defaultns.
func (scope *Scope) ResolveDefault(qname, defaultns string) xml.Name {
	if defaultns == "" || strings.Contains(qname, ":") {
		return scope.Resolve(qname)
	}
	return xml.Name{Space: defaultns, Local: qname}
}

// Prefix is the inverse of Resolve. It uses the closest prefix
// defined for a namespace to create a string of the form
// prefix:local. If the namespace cannot be found, or is the
// default namespace, an unqualified name is returned.
func (scope *Scope) Prefix(name xml.Name) (qname string) {
	switch name.Space {
	case "":
		return name.Local
	case xmlLangURI:
		return "xml:" + name.Local
	case xmlNamespaceURI:
		return "xmlns:" + name.Local
	}
	for i := len(scope.ns) - 1; i >= 0; i-- {
		if scope.ns[i].Space == name.Space {
			if scope.ns[i].Local == "" {
				// Favor default NS if there is an extra
				// qualified NS declaration
				qname = name.Local
			} else if len(qname) == 0 {
				qname = scope.ns[i].Local + ":" + name.Local
			}
		}
	}
	return qname
}

func (scope *Scope) pushNS(tag xml.StartElement) []xml.Attr {
	var ns []xml.Name
	var newAttrs []xml.Attr
	for _, attr := range tag.Attr {
		if attr.Name.Space == "xmlns" {
			ns = append(ns, xml.Name{attr.Value, attr.Name.Local})
		} else if attr.Name.Local == "xmlns" {
			ns = append(ns, xml.Name{attr.Value, ""})
		} else {
			newAttrs = append(newAttrs, attr)
		}
	}
	// Within a single tag, all ns declarations are sorted. This reduces
	// differences between xmlns declarations between tags when
	// modifying the xml tree.
	sort.Sort(byXMLName(ns))
	if len(ns) > 0 {
		scope.ns = append(scope.ns, ns...)
		// Ensure that future additions to the scope create
		// a new backing array. This prevents the scope from
		// being clobbered during parsing.
		scope.ns = scope.ns[:len(scope.ns):len(scope.ns)]
	}
	return newAttrs
}
