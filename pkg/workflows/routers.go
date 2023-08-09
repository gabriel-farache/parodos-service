package workflows

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitWorkflows(router *gin.Engine) {
	workflowDefinitionHandler := newWorkflowDefinitionHandler()
	registerRouter(router, workflowDefinitionHandler)
}

func registerRouter(router *gin.Engine, handler *WorkflowDefinitionHandler) {
	api := router.Group("/api")
	v1 := api.Group("/v1")
	group := v1.Group("/groups")
	workflow := group.Group(fmt.Sprintf("/:%s/workflows", GroupIdParam))

	group.GET("", handler.GetGroups)

	workflow.GET("", handler.GetWorkflows)
	workflow.GET(fmt.Sprintf("/:%s", WorkflowIdParam), handler.GetWorkflow)
}
