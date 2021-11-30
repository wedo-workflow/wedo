package runtime

import (
	"context"
	"testing"

	wedo "github.com/wedo-workflow/wedo/api"
)

func TestRuntime_TaskCreate(t *testing.T) {
	ret, err := initRuntime().TaskCreate(context.Background(), &wedo.TaskCreateRequest{
		Name:              "task1",
		Description:       "",
		Assignee:          "",
		Owner:             "",
		DelegationState:   "",
		Due:               nil,
		Priority:          "",
		ParentTaskId:      "",
		ProcessInstanceId: "1234",
		TaskDefinitionKey: "",
		NamespaceId:       "",
		FormKey:           "",
		FollowUpDate:      nil,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
