package workflows

import (
	"github.com/gin-gonic/gin"
)

type WorkflowDefinitionHandler struct {
	command *WorkflowCommandHandler
	query   *WorkflowQueryHandler
}

func newWorkflowDefinitionHandler() *WorkflowDefinitionHandler {
	return &WorkflowDefinitionHandler{
		command: newWorkflowCommandHandler(),
		query:   newWorkflowQueryHandler(),
	}
}

// @Summary Get a list of workflows definitions
// @Description return the list of workflows definitions registered.
// @Accept  json
// @Produce  json
// @Success 200 {array} Workflow "ok"
// @Router /workflows [get]
func (handler WorkflowDefinitionHandler) GetWorkflows(ctx *gin.Context) {
	ctx.JSON(200, []Workflow{
		{Name: "test1"},
		{Name: "test2"},
	})
}
