#!/usr/bin/env bash

 xsdgen -o BPMN20_xsdgen.go -pkg schema -f -vv BPMN20.xsd
 xsdgen -o bpmndi/BPMNDI_xsdgen.go -pkg schema -f -vv BPMNDI.xsd
 xsdgen -o dc/DC_xsdgen.go -pkg schema -f -vv DC.xsd
 xsdgen -o di/DI_xsdgen.go -pkg schema -f -vv DI.xsd
 xsdgen -o Semantic_xsdgen.go -pkg schema -f -vv Semantic.xsd