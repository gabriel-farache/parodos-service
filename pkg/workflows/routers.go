package workflows

import (
	"github.com/gin-gonic/gin"
)

func InitWorkflows(router *gin.Engine) {
	workflowDefinitionHandler := newWorkflowDefinitionHandler()
	registerRouter(router, workflowDefinitionHandler)
}

func registerRouter(router *gin.Engine, handler *WorkflowDefinitionHandler) {
	workFlowRouter := router.Group("/v1/workflows")
	workFlowRouter.GET("", handler.GetWorkflows)
}
