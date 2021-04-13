package bpmn_model

// Data is represented with the four elements:
// 1. Data Objects
// 2. Data Inputs
// 3. Data Outputs
// 4. Data Stores
type Data interface {
	ToString() string
}
