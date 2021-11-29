package element

import (
	"github.com/wedo-workflow/wedo/element/bpmn"
	"github.com/wedo-workflow/xmltree"
)

type Element interface {
	// EID is id of element.
	EID() string

	// TypeName is element's type name.
	TypeName() string
	SetTypeName(string) error

	// Parse parse xml dom element to element object.
	Parse(element *xmltree.Element) error
}

// todo : dynamic register element instance
func DefaultRegister() map[string]Element {
	return map[string]Element{
		BPMN_ELEMENT_DEFINITIONS:              bpmn.NewDefinitions(),
		BPMN_ELEMENT_PROCESS:                  bpmn.NewProcess(),
		BPMN_ELEMENT_MESSAGE:                  bpmn.NewMessage(),
		BPMN_ELEMENT_GROUP:                    bpmn.NewGroup(),
		BPMN_ELEMENT_SIGNAL:                   bpmn.NewSignal(),
		BPMN_ELEMENT_TASK:                     bpmn.NewTask(),
		BPMN_ELEMENT_INCOMING:                 bpmn.NewIncoming(),
		BPMN_ELEMENT_OUTGOING:                 bpmn.NewOutgoing(),
		BPMN_ELEMENT_COLLABORATION:            bpmn.NewCollaboration(),
		BPMN_ELEMENT_END_EVENT:                bpmn.NewEndEvent(),
		BPMN_ELEMENT_START_EVENT:              bpmn.NewStartEvent(),
		BPMN_ELEMENT_EXTENSION_ELEMENTS:       bpmn.NewExtensionElements(),
		BPMN_ELEMENT_INTERMEDIATE_THROW_EVENT: bpmn.NewIntermediateThrowEvent(),
		BPMN_ELEMENT_MANUAL_TASK:              bpmn.NewManualTask(),
		BPMN_ELEMENT_MESSAGE_EVENT_DEFINITION: bpmn.NewMessageEventDefinition(),
		BPMN_ELEMENT_PARTICIPANT:              bpmn.NewParticipant(),
		BPMN_ELEMENT_RECEIVE_TASK:             bpmn.NewReceiveTask(),
		BPMN_ELEMENT_SEND_TASK:                bpmn.NewSendTask(),
		BPMN_ELEMENT_SEQUENCE_FLOW:            bpmn.NewSequenceFlow(),
		BPMN_ELEMENT_SIGNAL_EVENT_DEFINITION:  bpmn.NewSignalEventDefinition(),
		WEDO_ELEMENT_FORM_DATA:                bpmn.NewFormData(),

		BPMNDI_ELEMENT_BPMN_DIAGRAM: bpmn.NewBPMNDiagram(),
		BPMNDI_ELEMENT_BPMN_PLANE:   bpmn.NewBPMNPlane(),
	}
}
