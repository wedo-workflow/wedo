package model

import wedo "github.com/wedo-workflow/wedo/api"

type ProcessInstance struct {
	Id     string                    `json:"id"`
	Status wedo.ProcessInstanceState `json:"status"`
}
