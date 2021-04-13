package bpmn_model

// ListenerType defined listener type of listener.
type ListenerType string

// fixed listener type.
var (
	ListenerTypeExpression      ListenerType = "expression"
	ListenerTypeProxyExpression ListenerType = "proxy_expression"
)

type Listener interface {
	Type() ListenerType
	EventType() EventType
	Fields() []ListenerField
}

type ListenerFieldType string

var (
	ListenerFieldTypeExpression ListenerFieldType = "expression"
	ListenerFieldTypeString     ListenerFieldType = "string"
)

type ListenerField interface {
	Name() string
	Type() ListenerFieldType
}
