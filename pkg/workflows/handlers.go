package workflows

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/swaggo/swag/example/celler/httputil"
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

// @Summary 	Get a list of groups
// @Description return the list of groups registered.
// @Accept 		json
// @Produce  	json
// @Success 	200 {array} Group "ok"
// @Failure		400  {object}  httputil.HTTPError
// @Failure		404  {object}  httputil.HTTPError
// @Failure		500  {object}  httputil.HTTPError
// @Router 		/groups [get]
func (handler WorkflowDefinitionHandler) GetGroups(ctx *gin.Context) {
	ctx.JSON(200, []Group{
		{Name: "parodos", Repository: "https://github.com/parodos-dev/parodos"},
		{Name: "parodos-service", Repository: "https://github.com/parodos-dev/parodos-service"},
	})
}

// @Summary 	Get the details of a registered group
// @Description return the details of a given registered group.
// @Accept 		json
// @Produce  	json
// @Success 	200 {object} GroupDetails "ok"
// @Failure		400  {object}  httputil.HTTPError
// @Failure		404  {object}  httputil.HTTPError
// @Failure		500  {object}  httputil.HTTPError
// @Router 		/groups/{group_id} [get]
func (handler WorkflowDefinitionHandler) GetGroup(ctx *gin.Context) {
	groupId, _ := ctx.Get(GroupIdParam)
	// TODO actually checks if the group exists
	groupExists := true
	if !groupExists {
		httputil.NewError(ctx, http.StatusNotFound, fmt.Errorf("Group %q not found", GroupIdParam))
		return
	}
	glog.Infof("Get details of groups %q", groupId)

	ctx.JSON(200, GroupDetails{
		Group: Group{
			Name:       "parodos-service",
			Repository: "https://github.com/parodos-dev/parodos-service",
		},
		Workflows: []Workflow{
			{Name: "test1"},
			{Name: "test2"},
		},
	})
}

// @Summary Get a list of workflows definitions in the group
// @Description return the list of workflows definitions registered in the given group.
// @Accept  	json
// @Produce  	json
// @Success 	200 {array} Workflow "ok"
// @Failure		400  {object}  httputil.HTTPError
// @Failure		404  {object}  httputil.HTTPError
// @Failure		500  {object}  httputil.HTTPError
// @Router 		/groups/{group_id}/workflows [get]
func (handler WorkflowDefinitionHandler) GetWorkflows(ctx *gin.Context) {
	groupId, _ := ctx.Get(GroupIdParam)
	// TODO actually checks if the group exists
	groupExists := true
	if !groupExists {
		httputil.NewError(ctx, http.StatusNotFound, fmt.Errorf("Group %q not found", GroupIdParam))
		return
	}
	glog.Infof("Get workflows of groups %q", groupId)
	ctx.JSON(200, []Workflow{
		{Name: "test1"},
		{Name: "test2"},
	})
}

// @Summary Get a workflow definition
// @Description return the workflow definition registered.
// @Accept  	json
// @Produce  	json
// @Success 	200 {object} Workflow "ok"
// @Failure		400  {object}  httputil.HTTPError
// @Failure		404  {object}  httputil.HTTPError
// @Failure		500  {object}  httputil.HTTPError
// @Router 		/groups/{group_id}/workflows/{workflow_id} [get]
func (handler WorkflowDefinitionHandler) GetWorkflow(ctx *gin.Context) {
	groupId, _ := ctx.Get(GroupIdParam)
	//if !groupExists {
	//	httputil.NewError(ctx, http.StatusNotFound, fmt.Errorf("Group %q not found", GroupIdParam))
	//	return
	//}
	workflowId, _ := ctx.Get(WorkflowIdParam)
	//if !workflowExists {
	//	httputil.NewError(ctx, http.StatusNotFound, fmt.Errorf("Workflow %q of group %q not found", WorkflowIdParam, GroupIdParam))
	//	return
	//}
	glog.Infof("Get workflow %q of groups %q", workflowId, groupId)

	ctx.JSON(200, Workflow{Name: workflowId.(string)})
}
