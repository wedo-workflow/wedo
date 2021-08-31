package element

// Constants used in the BPMN 2.0 Language (DI + Semantic)
var (
	// XSI_NS The XSI namespace
	XSI_NS = "http://www.w3.org/2001/XMLSchema-instance"

	// BPMN20_NS The BPMN 2.0 namespace
	BPMN20_NS = "http://www.omg.org/spec/BPMN/20100524/MODEL"

	// BPMNDI_NS The BPMNDI namespace
	BPMNDI_NS = "http://www.omg.org/spec/BPMN/20100524/DI"

	// DC_NS The DC namespace
	DC_NS = "http://www.omg.org/spec/DD/20100524/DC"

	// DI_NS The DI namespace
	DI_NS = "http://www.omg.org/spec/DD/20100524/DI"

	// BPMN_20_SCHEMA_LOCATION The location of the BPMN 2.0 XML schema.
	BPMN_20_SCHEMA_LOCATION = "github.com/wedo-workflow/wedo/pkg/bpmn/schema/BPMN20.xsd"

	// XML_SCHEMA_NS Xml Schema is the default type language
	XML_SCHEMA_NS = "http://www.w3.org/2001/XMLSchema"

	XPATH_NS = "http://www.w3.org/1999/XPath"

	WEDO_NS = "https://github.com/wedo-workflow/wedo/bpmn/1.0"

	BPMN_ELEMENT_BASE_ELEMENT                              = "baseElement"
	BPMN_ELEMENT_DEFINITIONS                               = "definitions"
	BPMN_ELEMENT_DOCUMENTATION                             = "documentation"
	BPMN_ELEMENT_EXTENSION                                 = "extension"
	BPMN_ELEMENT_EXTENSION_ELEMENTS                        = "extensionElements"
	BPMN_ELEMENT_IMPORT                                    = "import"
	BPMN_ELEMENT_RELATIONSHIP                              = "relationship"
	BPMN_ELEMENT_SOURCE                                    = "source"
	BPMN_ELEMENT_TARGET                                    = "target"
	BPMN_ELEMENT_ROOT_ELEMENT                              = "rootElement"
	BPMN_ELEMENT_AUDITING                                  = "auditing"
	BPMN_ELEMENT_MONITORING                                = "monitoring"
	BPMN_ELEMENT_CATEGORY_VALUE                            = "categoryValue"
	BPMN_ELEMENT_FLOW_ELEMENT                              = "flowElement"
	BPMN_ELEMENT_FLOW_NODE                                 = "flowNode"
	BPMN_ELEMENT_CATEGORY_VALUE_REF                        = "categoryValueRef"
	BPMN_ELEMENT_EXPRESSION                                = "expression"
	BPMN_ELEMENT_CONDITION_EXPRESSION                      = "conditionExpression"
	BPMN_ELEMENT_SEQUENCE_FLOW                             = "sequenceFlow"
	BPMN_ELEMENT_INCOMING                                  = "incoming"
	BPMN_ELEMENT_OUTGOING                                  = "outgoing"
	BPMN_ELEMENT_DATA_STATE                                = "dataState"
	BPMN_ELEMENT_ITEM_DEFINITION                           = "itemDefinition"
	BPMN_ELEMENT_ERROR                                     = "error"
	BPMN_ELEMENT_IN_MESSAGE_REF                            = "inMessageRef"
	BPMN_ELEMENT_OUT_MESSAGE_REF                           = "outMessageRef"
	BPMN_ELEMENT_ERROR_REF                                 = "errorRef"
	BPMN_ELEMENT_OPERATION                                 = "operation"
	BPMN_ELEMENT_IMPLEMENTATION_REF                        = "implementationRef"
	BPMN_ELEMENT_OPERATION_REF                             = "operationRef"
	BPMN_ELEMENT_DATA_OUTPUT                               = "dataOutput"
	BPMN_ELEMENT_FROM                                      = "from"
	BPMN_ELEMENT_TO                                        = "to"
	BPMN_ELEMENT_ASSIGNMENT                                = "assignment"
	BPMN_ELEMENT_ITEM_AWARE_ELEMENT                        = "itemAwareElement"
	BPMN_ELEMENT_DATA_OBJECT                               = "dataObject"
	BPMN_ELEMENT_DATA_OBJECT_REFERENCE                     = "dataObjectReference"
	BPMN_ELEMENT_DATA_STORE                                = "dataStore"
	BPMN_ELEMENT_DATA_STORE_REFERENCE                      = "dataStoreReference"
	BPMN_ELEMENT_DATA_INPUT                                = "dataInput"
	BPMN_ELEMENT_FORMAL_EXPRESSION                         = "formalExpression"
	BPMN_ELEMENT_DATA_ASSOCIATION                          = "dataAssociation"
	BPMN_ELEMENT_SOURCE_REF                                = "sourceRef"
	BPMN_ELEMENT_TARGET_REF                                = "targetRef"
	BPMN_ELEMENT_TRANSFORMATION                            = "transformation"
	BPMN_ELEMENT_DATA_INPUT_ASSOCIATION                    = "dataInputAssociation"
	BPMN_ELEMENT_DATA_OUTPUT_ASSOCIATION                   = "dataOutputAssociation"
	BPMN_ELEMENT_INPUT_SET                                 = "inputSet"
	BPMN_ELEMENT_OUTPUT_SET                                = "outputSet"
	BPMN_ELEMENT_DATA_INPUT_REFS                           = "dataInputRefs"
	BPMN_ELEMENT_OPTIONAL_INPUT_REFS                       = "optionalInputRefs"
	BPMN_ELEMENT_WHILE_EXECUTING_INPUT_REFS                = "whileExecutingInputRefs"
	BPMN_ELEMENT_OUTPUT_SET_REFS                           = "outputSetRefs"
	BPMN_ELEMENT_DATA_OUTPUT_REFS                          = "dataOutputRefs"
	BPMN_ELEMENT_OPTIONAL_OUTPUT_REFS                      = "optionalOutputRefs"
	BPMN_ELEMENT_WHILE_EXECUTING_OUTPUT_REFS               = "whileExecutingOutputRefs"
	BPMN_ELEMENT_INPUT_SET_REFS                            = "inputSetRefs"
	BPMN_ELEMENT_CATCH_EVENT                               = "catchEvent"
	BPMN_ELEMENT_THROW_EVENT                               = "throwEvent"
	BPMN_ELEMENT_END_EVENT                                 = "endEvent"
	BPMN_ELEMENT_IO_SPECIFICATION                          = "ioSpecification"
	BPMN_ELEMENT_LOOP_CHARACTERISTICS                      = "loopCharacteristics"
	BPMN_ELEMENT_RESOURCE_PARAMETER                        = "resourceParameter"
	BPMN_ELEMENT_RESOURCE                                  = "resource"
	BPMN_ELEMENT_RESOURCE_PARAMETER_BINDING                = "resourceParameterBinding"
	BPMN_ELEMENT_RESOURCE_ASSIGNMENT_EXPRESSION            = "resourceAssignmentExpression"
	BPMN_ELEMENT_RESOURCE_ROLE                             = "resourceRole"
	BPMN_ELEMENT_RESOURCE_REF                              = "resourceRef"
	BPMN_ELEMENT_PERFORMER                                 = "performer"
	BPMN_ELEMENT_HUMAN_PERFORMER                           = "humanPerformer"
	BPMN_ELEMENT_POTENTIAL_OWNER                           = "potentialOwner"
	BPMN_ELEMENT_ACTIVITY                                  = "activity"
	BPMN_ELEMENT_IO_BINDING                                = "ioBinding"
	BPMN_ELEMENT_INTERFACE                                 = "interface"
	BPMN_ELEMENT_EVENT                                     = "event"
	BPMN_ELEMENT_MESSAGE                                   = "message"
	BPMN_ELEMENT_START_EVENT                               = "startEvent"
	BPMN_ELEMENT_PROPERTY                                  = "property"
	BPMN_ELEMENT_EVENT_DEFINITION                          = "eventDefinition"
	BPMN_ELEMENT_EVENT_DEFINITION_REF                      = "eventDefinitionRef"
	BPMN_ELEMENT_MESSAGE_EVENT_DEFINITION                  = "messageEventDefinition"
	BPMN_ELEMENT_CANCEL_EVENT_DEFINITION                   = "cancelEventDefinition"
	BPMN_ELEMENT_COMPENSATE_EVENT_DEFINITION               = "compensateEventDefinition"
	BPMN_ELEMENT_CONDITIONAL_EVENT_DEFINITION              = "conditionalEventDefinition"
	BPMN_ELEMENT_CONDITION                                 = "condition"
	BPMN_ELEMENT_ERROR_EVENT_DEFINITION                    = "errorEventDefinition"
	BPMN_ELEMENT_LINK_EVENT_DEFINITION                     = "linkEventDefinition"
	BPMN_ELEMENT_SIGNAL_EVENT_DEFINITION                   = "signalEventDefinition"
	BPMN_ELEMENT_TERMINATE_EVENT_DEFINITION                = "terminateEventDefinition"
	BPMN_ELEMENT_TIMER_EVENT_DEFINITION                    = "timerEventDefinition"
	BPMN_ELEMENT_SUPPORTED_INTERFACE_REF                   = "supportedInterfaceRef"
	BPMN_ELEMENT_CALLABLE_ELEMENT                          = "callableElement"
	BPMN_ELEMENT_PARTITION_ELEMENT                         = "partitionElement"
	BPMN_ELEMENT_FLOW_NODE_REF                             = "flowNodeRef"
	BPMN_ELEMENT_CHILD_LANE_SET                            = "childLaneSet"
	BPMN_ELEMENT_LANE_SET                                  = "laneSet"
	BPMN_ELEMENT_LANE                                      = "lane"
	BPMN_ELEMENT_ARTIFACT                                  = "artifact"
	BPMN_ELEMENT_CORRELATION_PROPERTY_RETRIEVAL_EXPRESSION = "correlationPropertyRetrievalExpression"
	BPMN_ELEMENT_MESSAGE_PATH                              = "messagePath"
	BPMN_ELEMENT_DATA_PATH                                 = "dataPath"
	BPMN_ELEMENT_CALL_ACTIVITY                             = "callActivity"
	BPMN_ELEMENT_CORRELATION_PROPERTY_BINDING              = "correlationPropertyBinding"
	BPMN_ELEMENT_CORRELATION_PROPERTY                      = "correlationProperty"
	BPMN_ELEMENT_CORRELATION_PROPERTY_REF                  = "correlationPropertyRef"
	BPMN_ELEMENT_CORRELATION_KEY                           = "correlationKey"
	BPMN_ELEMENT_CORRELATION_SUBSCRIPTION                  = "correlationSubscription"
	BPMN_ELEMENT_SUPPORTS                                  = "supports"
	BPMN_ELEMENT_PROCESS                                   = "process"
	BPMN_ELEMENT_TASK                                      = "task"
	BPMN_ELEMENT_SEND_TASK                                 = "sendTask"
	BPMN_ELEMENT_SERVICE_TASK                              = "serviceTask"
	BPMN_ELEMENT_SCRIPT_TASK                               = "scriptTask"
	BPMN_ELEMENT_USER_TASK                                 = "userTask"
	BPMN_ELEMENT_RECEIVE_TASK                              = "receiveTask"
	BPMN_ELEMENT_BUSINESS_RULE_TASK                        = "businessRuleTask"
	BPMN_ELEMENT_MANUAL_TASK                               = "manualTask"
	BPMN_ELEMENT_SCRIPT                                    = "script"
	BPMN_ELEMENT_RENDERING                                 = "rendering"
	BPMN_ELEMENT_BOUNDARY_EVENT                            = "boundaryEvent"
	BPMN_ELEMENT_SUB_PROCESS                               = "subProcess"
	BPMN_ELEMENT_TRANSACTION                               = "transaction"
	BPMN_ELEMENT_GATEWAY                                   = "gateway"
	BPMN_ELEMENT_PARALLEL_GATEWAY                          = "parallelGateway"
	BPMN_ELEMENT_EXCLUSIVE_GATEWAY                         = "exclusiveGateway"
	BPMN_ELEMENT_INTERMEDIATE_CATCH_EVENT                  = "intermediateCatchEvent"
	BPMN_ELEMENT_INTERMEDIATE_THROW_EVENT                  = "intermediateThrowEvent"
	BPMN_ELEMENT_END_POINT                                 = "endPoint"
	BPMN_ELEMENT_PARTICIPANT_MULTIPLICITY                  = "participantMultiplicity"
	BPMN_ELEMENT_PARTICIPANT                               = "participant"
	BPMN_ELEMENT_PARTICIPANT_REF                           = "participantRef"
	BPMN_ELEMENT_INTERFACE_REF                             = "interfaceRef"
	BPMN_ELEMENT_END_POINT_REF                             = "endPointRef"
	BPMN_ELEMENT_MESSAGE_FLOW                              = "messageFlow"
	BPMN_ELEMENT_MESSAGE_FLOW_REF                          = "messageFlowRef"
	BPMN_ELEMENT_CONVERSATION_NODE                         = "conversationNode"
	BPMN_ELEMENT_CONVERSATION                              = "conversation"
	BPMN_ELEMENT_SUB_CONVERSATION                          = "subConversation"
	BPMN_ELEMENT_GLOBAL_CONVERSATION                       = "globalConversation"
	BPMN_ELEMENT_CALL_CONVERSATION                         = "callConversation"
	BPMN_ELEMENT_PARTICIPANT_ASSOCIATION                   = "participantAssociation"
	BPMN_ELEMENT_INNER_PARTICIPANT_REF                     = "innerParticipantRef"
	BPMN_ELEMENT_OUTER_PARTICIPANT_REF                     = "outerParticipantRef"
	BPMN_ELEMENT_CONVERSATION_ASSOCIATION                  = "conversationAssociation"
	BPMN_ELEMENT_MESSAGE_FLOW_ASSOCIATION                  = "messageFlowAssociation"
	BPMN_ELEMENT_CONVERSATION_LINK                         = "conversationLink"
	BPMN_ELEMENT_COLLABORATION                             = "collaboration"
	BPMN_ELEMENT_ASSOCIATION                               = "association"
	BPMN_ELEMENT_SIGNAL                                    = "signal"
	BPMN_ELEMENT_TIME_DATE                                 = "timeDate"
	BPMN_ELEMENT_TIME_DURATION                             = "timeDuration"
	BPMN_ELEMENT_TIME_CYCLE                                = "timeCycle"
	BPMN_ELEMENT_ESCALATION                                = "escalation"
	BPMN_ELEMENT_ESCALATION_EVENT_DEFINITION               = "escalationEventDefinition"
	BPMN_ELEMENT_ACTIVATION_CONDITION                      = "activationCondition"
	BPMN_ELEMENT_COMPLEX_GATEWAY                           = "complexGateway"
	BPMN_ELEMENT_EVENT_BASED_GATEWAY                       = "eventBasedGateway"
	BPMN_ELEMENT_INCLUSIVE_GATEWAY                         = "inclusiveGateway"
	BPMN_ELEMENT_TEXT_ANNOTATION                           = "textAnnotation"
	BPMN_ELEMENT_TEXT                                      = "text"
	BPMN_ELEMENT_COMPLEX_BEHAVIOR_DEFINITION               = "complexBehaviorDefinition"
	BPMN_ELEMENT_MULTI_INSTANCE_LOOP_CHARACTERISTICS       = "multiInstanceLoopCharacteristics"
	BPMN_ELEMENT_LOOP_CARDINALITY                          = "loopCardinality"
	BPMN_ELEMENT_COMPLETION_CONDITION                      = "completionCondition"
	BPMN_ELEMENT_OUTPUT_DATA_ITEM                          = "outputDataItem"
	BPMN_ELEMENT_INPUT_DATA_ITEM                           = "inputDataItem"
	BPMN_ELEMENT_LOOP_DATA_OUTPUT_REF                      = "loopDataOutputRef"
	BPMN_ELEMENT_LOOP_DATA_INPUT_REF                       = "loopDataInputRef"
	BPMN_ELEMENT_IS_SEQUENTIAL                             = "isSequential"
	BPMN_ELEMENT_BEHAVIOR                                  = "behavior"
	BPMN_ELEMENT_ONE_BEHAVIOR_EVENT_REF                    = "oneBehaviorEventRef"
	BPMN_ELEMENT_NONE_BEHAVIOR_EVENT_REF                   = "noneBehaviorEventRef"
	BPMN_ELEMENT_GROUP                                     = "group"
	BPMN_ELEMENT_CATEGORY                                  = "category"

	// DC
	DC_ELEMENT_FONT   = "Font"
	DC_ELEMENT_POINT  = "Point"
	DC_ELEMENT_BOUNDS = "Bounds"

	// DI
	DI_ELEMENT_DIAGRAM_ELEMENT = "DiagramElement"
	DI_ELEMENT_DIAGRAM         = "Diagram"
	DI_ELEMENT_EDGE            = "Edge"
	DI_ELEMENT_EXTENSION       = "extension"
	DI_ELEMENT_LABELED_EDGE    = "LabeledEdge"
	DI_ELEMENT_LABEL           = "Label"
	DI_ELEMENT_LABELED_SHAPE   = "LabeledShape"
	DI_ELEMENT_NODE            = "Node"
	DI_ELEMENT_PLANE           = "Plane"
	DI_ELEMENT_SHAPE           = "Shape"
	DI_ELEMENT_STYLE           = "Style"
	DI_ELEMENT_WAYPOINT        = "waypoint"

	// BPMNDI
	BPMNDI_ELEMENT_BPMN_DIAGRAM     = "BPMNDiagram"
	BPMNDI_ELEMENT_BPMN_PLANE       = "BPMNPlane"
	BPMNDI_ELEMENT_BPMN_LABEL_STYLE = "BPMNLabelStyle"
	BPMNDI_ELEMENT_BPMN_SHAPE       = "BPMNShape"
	BPMNDI_ELEMENT_BPMN_LABEL       = "BPMNLabel"
	BPMNDI_ELEMENT_BPMN_EDGE        = "BPMNEdge"

	// model extensions
	WEDO_ELEMENT_CONNECTOR                   = "connector"
	WEDO_ELEMENT_CONNECTOR_ID                = "connectorId"
	WEDO_ELEMENT_CONSTRAINT                  = "constraint"
	WEDO_ELEMENT_ENTRY                       = "entry"
	WEDO_ELEMENT_ERROR_EVENT_DEFINITION      = "errorEventDefinition"
	WEDO_ELEMENT_EXECUTION_LISTENER          = "executionListener"
	WEDO_ELEMENT_EXPRESSION                  = "expression"
	WEDO_ELEMENT_FAILED_JOB_RETRY_TIME_CYCLE = "failedJobRetryTimeCycle"
	WEDO_ELEMENT_FIELD                       = "field"
	WEDO_ELEMENT_FORM_DATA                   = "formData"
	WEDO_ELEMENT_FORM_FIELD                  = "formField"
	WEDO_ELEMENT_FORM_PROPERTY               = "formProperty"
	WEDO_ELEMENT_IN                          = "in"
	WEDO_ELEMENT_INPUT_OUTPUT                = "inputOutput"
	WEDO_ELEMENT_INPUT_PARAMETER             = "inputParameter"
	WEDO_ELEMENT_LIST                        = "list"
	WEDO_ELEMENT_MAP                         = "map"
	WEDO_ELEMENT_OUTPUT_PARAMETER            = "outputParameter"
	WEDO_ELEMENT_OUT                         = "out"
	WEDO_ELEMENT_POTENTIAL_STARTER           = "potentialStarter"
	WEDO_ELEMENT_PROPERTIES                  = "properties"
	WEDO_ELEMENT_PROPERTY                    = "property"
	WEDO_ELEMENT_SCRIPT                      = "script"
	WEDO_ELEMENT_STRING                      = "string"
	WEDO_ELEMENT_TASK_LISTENER               = "taskListener"
	WEDO_ELEMENT_VALIDATION                  = "validation"
	WEDO_ELEMENT_VALUE                       = "value"

	// attributes //

	// XSI attributes
	XSI_ATTRIBUTE_TYPE = "type"

	// BPMN attributes
	BPMN_ATTRIBUTE_EXPORTER                    = "exporter"
	BPMN_ATTRIBUTE_EXPORTER_VERSION            = "exporterVersion"
	BPMN_ATTRIBUTE_EXPRESSION_LANGUAGE         = "expressionLanguage"
	BPMN_ATTRIBUTE_ID                          = "id"
	BPMN_ATTRIBUTE_NAME                        = "name"
	BPMN_ATTRIBUTE_TARGET_NAMESPACE            = "targetNamespace"
	BPMN_ATTRIBUTE_TYPE_LANGUAGE               = "typeLanguage"
	BPMN_ATTRIBUTE_NAMESPACE                   = "namespace"
	BPMN_ATTRIBUTE_LOCATION                    = "location"
	BPMN_ATTRIBUTE_IMPORT_TYPE                 = "importType"
	BPMN_ATTRIBUTE_TEXT_FORMAT                 = "textFormat"
	BPMN_ATTRIBUTE_PROCESS_TYPE                = "processType"
	BPMN_ATTRIBUTE_IS_CLOSED                   = "isClosed"
	BPMN_ATTRIBUTE_IS_EXECUTABLE               = "isExecutable"
	BPMN_ATTRIBUTE_MESSAGE_REF                 = "messageRef"
	BPMN_ATTRIBUTE_DEFINITION                  = "definition"
	BPMN_ATTRIBUTE_MUST_UNDERSTAND             = "mustUnderstand"
	BPMN_ATTRIBUTE_TYPE                        = "type"
	BPMN_ATTRIBUTE_DIRECTION                   = "direction"
	BPMN_ATTRIBUTE_SOURCE_REF                  = "sourceRef"
	BPMN_ATTRIBUTE_TARGET_REF                  = "targetRef"
	BPMN_ATTRIBUTE_IS_IMMEDIATE                = "isImmediate"
	BPMN_ATTRIBUTE_VALUE                       = "value"
	BPMN_ATTRIBUTE_STRUCTURE_REF               = "structureRef"
	BPMN_ATTRIBUTE_IS_COLLECTION               = "isCollection"
	BPMN_ATTRIBUTE_ITEM_KIND                   = "itemKind"
	BPMN_ATTRIBUTE_ITEM_REF                    = "itemRef"
	BPMN_ATTRIBUTE_ITEM_SUBJECT_REF            = "itemSubjectRef"
	BPMN_ATTRIBUTE_ERROR_CODE                  = "errorCode"
	BPMN_ATTRIBUTE_LANGUAGE                    = "language"
	BPMN_ATTRIBUTE_EVALUATES_TO_TYPE_REF       = "evaluatesToTypeRef"
	BPMN_ATTRIBUTE_PARALLEL_MULTIPLE           = "parallelMultiple"
	BPMN_ATTRIBUTE_IS_INTERRUPTING             = "isInterrupting"
	BPMN_ATTRIBUTE_IS_REQUIRED                 = "isRequired"
	BPMN_ATTRIBUTE_PARAMETER_REF               = "parameterRef"
	BPMN_ATTRIBUTE_IS_FOR_COMPENSATION         = "isForCompensation"
	BPMN_ATTRIBUTE_START_QUANTITY              = "startQuantity"
	BPMN_ATTRIBUTE_COMPLETION_QUANTITY         = "completionQuantity"
	BPMN_ATTRIBUTE_DEFAULT                     = "default"
	BPMN_ATTRIBUTE_OPERATION_REF               = "operationRef"
	BPMN_ATTRIBUTE_INPUT_DATA_REF              = "inputDataRef"
	BPMN_ATTRIBUTE_OUTPUT_DATA_REF             = "outputDataRef"
	BPMN_ATTRIBUTE_IMPLEMENTATION_REF          = "implementationRef"
	BPMN_ATTRIBUTE_PARTITION_ELEMENT_REF       = "partitionElementRef"
	BPMN_ATTRIBUTE_CORRELATION_PROPERTY_REF    = "correlationPropertyRef"
	BPMN_ATTRIBUTE_CORRELATION_KEY_REF         = "correlationKeyRef"
	BPMN_ATTRIBUTE_IMPLEMENTATION              = "implementation"
	BPMN_ATTRIBUTE_SCRIPT_FORMAT               = "scriptFormat"
	BPMN_ATTRIBUTE_INSTANTIATE                 = "instantiate"
	BPMN_ATTRIBUTE_CANCEL_ACTIVITY             = "cancelActivity"
	BPMN_ATTRIBUTE_ATTACHED_TO_REF             = "attachedToRef"
	BPMN_ATTRIBUTE_TRIGGERED_BY_EVENT          = "triggeredByEvent"
	BPMN_ATTRIBUTE_GATEWAY_DIRECTION           = "gatewayDirection"
	BPMN_ATTRIBUTE_CALLED_ELEMENT              = "calledElement"
	BPMN_ATTRIBUTE_MINIMUM                     = "minimum"
	BPMN_ATTRIBUTE_MAXIMUM                     = "maximum"
	BPMN_ATTRIBUTE_PROCESS_REF                 = "processRef"
	BPMN_ATTRIBUTE_CALLED_COLLABORATION_REF    = "calledCollaborationRef"
	BPMN_ATTRIBUTE_INNER_CONVERSATION_NODE_REF = "innerConversationNodeRef"
	BPMN_ATTRIBUTE_OUTER_CONVERSATION_NODE_REF = "outerConversationNodeRef"
	BPMN_ATTRIBUTE_INNER_MESSAGE_FLOW_REF      = "innerMessageFlowRef"
	BPMN_ATTRIBUTE_OUTER_MESSAGE_FLOW_REF      = "outerMessageFlowRef"
	BPMN_ATTRIBUTE_ASSOCIATION_DIRECTION       = "associationDirection"
	BPMN_ATTRIBUTE_WAIT_FOR_COMPLETION         = "waitForCompletion"
	BPMN_ATTRIBUTE_ACTIVITY_REF                = "activityRef"
	BPMN_ATTRIBUTE_ERROR_REF                   = "errorRef"
	BPMN_ATTRIBUTE_SIGNAL_REF                  = "signalRef"
	BPMN_ATTRIBUTE_ESCALATION_CODE             = "escalationCode"
	BPMN_ATTRIBUTE_ESCALATION_REF              = "escalationRef"
	BPMN_ATTRIBUTE_EVENT_GATEWAY_TYPE          = "eventGatewayType"
	BPMN_ATTRIBUTE_DATA_OBJECT_REF             = "dataObjectRef"
	BPMN_ATTRIBUTE_DATA_STORE_REF              = "dataStoreRef"
	BPMN_ATTRIBUTE_METHOD                      = "method"
	BPMN_ATTRIBUTE_CAPACITY                    = "capacity"
	BPMN_ATTRIBUTE_IS_UNLIMITED                = "isUnlimited"
	BPMN_ATTRIBUTE_CATEGORY_VALUE_REF          = "categoryValueRef"

	// DC
	DC_ATTRIBUTE_NAME              = "name"
	DC_ATTRIBUTE_SIZE              = "size"
	DC_ATTRIBUTE_IS_BOLD           = "isBold"
	DC_ATTRIBUTE_IS_ITALIC         = "isItalic"
	DC_ATTRIBUTE_IS_UNDERLINE      = "isUnderline"
	DC_ATTRIBUTE_IS_STRIKE_THROUGH = "isStrikeThrough"
	DC_ATTRIBUTE_X                 = "x"
	DC_ATTRIBUTE_Y                 = "y"
	DC_ATTRIBUTE_WIDTH             = "width"
	DC_ATTRIBUTE_HEIGHT            = "height"

	// DI
	DI_ATTRIBUTE_ID            = "id"
	DI_ATTRIBUTE_NAME          = "name"
	DI_ATTRIBUTE_DOCUMENTATION = "documentation"
	DI_ATTRIBUTE_RESOLUTION    = "resolution"

	// BPMNDI
	BPMNDI_ATTRIBUTE_BPMN_ELEMENT                = "bpmnElement"
	BPMNDI_ATTRIBUTE_SOURCE_ELEMENT              = "sourceElement"
	BPMNDI_ATTRIBUTE_TARGET_ELEMENT              = "targetElement"
	BPMNDI_ATTRIBUTE_MESSAGE_VISIBLE_KIND        = "messageVisibleKind"
	BPMNDI_ATTRIBUTE_IS_HORIZONTAL               = "isHorizontal"
	BPMNDI_ATTRIBUTE_IS_EXPANDED                 = "isExpanded"
	BPMNDI_ATTRIBUTE_IS_MARKER_VISIBLE           = "isMarkerVisible"
	BPMNDI_ATTRIBUTE_IS_MESSAGE_VISIBLE          = "isMessageVisible"
	BPMNDI_ATTRIBUTE_PARTICIPANT_BAND_KIND       = "participantBandKind"
	BPMNDI_ATTRIBUTE_CHOREOGRAPHY_ACTIVITY_SHAPE = "choreographyActivityShape"
	BPMNDI_ATTRIBUTE_LABEL_STYLE                 = "labelStyle"

	// model extensions
	WEDO_ATTRIBUTE_ASSIGNEE                             = "assignee"
	WEDO_ATTRIBUTE_ASYNC                                = "async"
	WEDO_ATTRIBUTE_ASYNC_BEFORE                         = "asyncBefore"
	WEDO_ATTRIBUTE_ASYNC_AFTER                          = "asyncAfter"
	WEDO_ATTRIBUTE_BUSINESS_KEY                         = "businessKey"
	WEDO_ATTRIBUTE_CALLED_ELEMENT_BINDING               = "calledElementBinding"
	WEDO_ATTRIBUTE_CALLED_ELEMENT_VERSION               = "calledElementVersion"
	WEDO_ATTRIBUTE_CALLED_ELEMENT_VERSION_TAG           = "calledElementVersionTag"
	WEDO_ATTRIBUTE_CALLED_ELEMENT_TENANT_ID             = "calledElementTenantId"
	WEDO_ATTRIBUTE_CANDIDATE_GROUPS                     = "candidateGroups"
	WEDO_ATTRIBUTE_CANDIDATE_STARTER_GROUPS             = "candidateStarterGroups"
	WEDO_ATTRIBUTE_CANDIDATE_STARTER_USERS              = "candidateStarterUsers"
	WEDO_ATTRIBUTE_CANDIDATE_USERS                      = "candidateUsers"
	WEDO_ATTRIBUTE_CLASS                                = "class"
	WEDO_ATTRIBUTE_COLLECTION                           = "collection"
	WEDO_ATTRIBUTE_CONFIG                               = "config"
	WEDO_ATTRIBUTE_DATE_PATTERN                         = "datePattern"
	WEDO_ATTRIBUTE_DECISION_REF                         = "decisionRef"
	WEDO_ATTRIBUTE_DECISION_REF_BINDING                 = "decisionRefBinding"
	WEDO_ATTRIBUTE_DECISION_REF_VERSION                 = "decisionRefVersion"
	WEDO_ATTRIBUTE_DECISION_REF_VERSION_TAG             = "decisionRefVersionTag"
	WEDO_ATTRIBUTE_DECISION_REF_TENANT_ID               = "decisionRefTenantId"
	WEDO_ATTRIBUTE_DEFAULT                              = "default"
	WEDO_ATTRIBUTE_DEFAULT_VALUE                        = "defaultValue"
	WEDO_ATTRIBUTE_DELEGATE_EXPRESSION                  = "delegateExpression"
	WEDO_ATTRIBUTE_DUE_DATE                             = "dueDate"
	WEDO_ATTRIBUTE_FOLLOW_UP_DATE                       = "followUpDate"
	WEDO_ATTRIBUTE_ELEMENT_VARIABLE                     = "elementVariable"
	WEDO_ATTRIBUTE_EVENT                                = "event"
	WEDO_ATTRIBUTE_ERROR_CODE_VARIABLE                  = "errorCodeVariable"
	WEDO_ATTRIBUTE_ERROR_MESSAGE_VARIABLE               = "errorMessageVariable"
	WEDO_ATTRIBUTE_ERROR_MESSAGE                        = "errorMessage"
	WEDO_ATTRIBUTE_EXCLUSIVE                            = "exclusive"
	WEDO_ATTRIBUTE_EXPRESSION                           = "expression"
	WEDO_ATTRIBUTE_FORM_HANDLER_CLASS                   = "formHandlerClass"
	WEDO_ATTRIBUTE_FORM_KEY                             = "formKey"
	WEDO_ATTRIBUTE_ID                                   = "id"
	WEDO_ATTRIBUTE_INITIATOR                            = "initiator"
	WEDO_ATTRIBUTE_JOB_PRIORITY                         = "jobPriority"
	WEDO_ATTRIBUTE_TASK_PRIORITY                        = "taskPriority"
	WEDO_ATTRIBUTE_KEY                                  = "key"
	WEDO_ATTRIBUTE_LABEL                                = "label"
	WEDO_ATTRIBUTE_LOCAL                                = "local"
	WEDO_ATTRIBUTE_MAP_DECISION_RESULT                  = "mapDecisionResult"
	WEDO_ATTRIBUTE_NAME                                 = "name"
	WEDO_ATTRIBUTE_PRIORITY                             = "priority"
	WEDO_ATTRIBUTE_READABLE                             = "readable"
	WEDO_ATTRIBUTE_REQUIRED                             = "required"
	WEDO_ATTRIBUTE_RESOURCE                             = "resource"
	WEDO_ATTRIBUTE_RESULT_VARIABLE                      = "resultVariable"
	WEDO_ATTRIBUTE_SCRIPT_FORMAT                        = "scriptFormat"
	WEDO_ATTRIBUTE_SOURCE                               = "source"
	WEDO_ATTRIBUTE_SOURCE_EXPRESSION                    = "sourceExpression"
	WEDO_ATTRIBUTE_STRING_VALUE                         = "stringValue"
	WEDO_ATTRIBUTE_TARGET                               = "target"
	WEDO_ATTRIBUTE_TOPIC                                = "topic"
	WEDO_ATTRIBUTE_TYPE                                 = "type"
	WEDO_ATTRIBUTE_VALUE                                = "value"
	WEDO_ATTRIBUTE_VARIABLE                             = "variable"
	WEDO_ATTRIBUTE_VARIABLE_MAPPING_CLASS               = "variableMappingClass"
	WEDO_ATTRIBUTE_VARIABLE_MAPPING_DELEGATE_EXPRESSION = "variableMappingDelegateExpression"
	WEDO_ATTRIBUTE_VARIABLES                            = "variables"
	WEDO_ATTRIBUTE_WRITEABLE                            = "writeable"
	WEDO_ATTRIBUTE_CASE_REF                             = "caseRef"
	WEDO_ATTRIBUTE_CASE_BINDING                         = "caseBinding"
	WEDO_ATTRIBUTE_CASE_VERSION                         = "caseVersion"
	WEDO_ATTRIBUTE_CASE_TENANT_ID                       = "caseTenantId"
	WEDO_ATTRIBUTE_VARIABLE_NAME                        = "variableName"
	WEDO_ATTRIBUTE_VARIABLE_EVENTS                      = "variableEvents"
	WEDO_ATTRIBUTE_HISTORY_TIME_TO_LIVE                 = "historyTimeToLive"
	WEDO_ATTRIBUTE_IS_STARTABLE_IN_TASKLIST             = "isStartableInTasklist"
	WEDO_ATTRIBUTE_VERSION_TAG                          = "versionTag"
)
