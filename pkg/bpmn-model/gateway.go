package bpmn_model

type GatewayControlType string

var (
	GatewayControlTypeExclusive          GatewayControlType = "exclusive"
	GatewayControlTypeEventBased         GatewayControlType = "event-based"
	GatewayControlTypeParallelEventBased GatewayControlType = "parallel-event-based"
	GatewayControlTypeInclusive          GatewayControlType = "inclusive"
	GatewayControlTypeComplex            GatewayControlType = "complex"
	GatewayControlTypeParallel           GatewayControlType = "parallel"
)

type Gateway interface {
	Flow
	Node
	Type() GatewayControlType
	Listeners() []Listener
	ExtProperty() []Property
	Note() Property
}
