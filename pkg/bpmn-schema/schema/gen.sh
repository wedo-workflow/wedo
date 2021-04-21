#!/usr/bin/env bash

# using https://github.com/droyo/go-xml/tree/master/xsdgen

xsdgen -o ../bpmn/BPMN20_xsdgen.go -pkg bpmn -f -vv BPMN20.xsd
xsdgen -o ../bpmndi/BPMNDI_xsdgen.go -pkg bpmndi -f -vv BPMNDI.xsd
xsdgen -o ../dc/DC_xsdgen.go -pkg dc -f -vv DC.xsd
xsdgen -o ../di/DI_xsdgen.go -pkg di -f -vv DI.xsd
xsdgen -o ../semantic/Semantic_xsdgen.go -pkg semantic -f -vv Semantic.xsd