package workflows

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parodos-dev/parodos-service/pkg/errors"
	"github.com/swaggo/swag/example/celler/httputil"
)

type WorkflowDefinitionHandler struct {
	command *WorkflowCommandHandler
	query   WorkflowsQuery
}

func newDefaultWorkflowDefinitionHandler() *WorkflowDefinitionHandler {
	return &WorkflowDefinitionHandler{
		command: newWorkflowCommandHandler(),
		query:   newWorkflowsQueryHandler(),
	}
}

func NewWorkflowDefinitionHandler(command *WorkflowCommandHandler, query WorkflowsQuery) *WorkflowDefinitionHandler {
	return &WorkflowDefinitionHandler{
		command: command,
		query:   query,
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
	groups, err := handler.query.ListGroups()
	if err != nil {
		processError(err, ctx)
		return
	}
	ctx.JSON(200, groups)
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
	groupId, ok := ctx.Get(GroupIdParam)
	if !ok {
		httputil.NewError(ctx, http.StatusBadRequest, fmt.Errorf("no group id provided"))
		return
	}
	group, err := handler.query.GetGroup(groupId.(string))
	if err != nil {
		processError(err, ctx)
		return
	}

	ctx.JSON(200, group)
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
	groupId, ok := ctx.Get(GroupIdParam)
	if !ok {
		httputil.NewError(ctx, http.StatusBadRequest, fmt.Errorf("no group id provided"))
		return
	}
	workflows, err := handler.query.ListWorkflows(groupId.(string))
	if err != nil {
		processError(err, ctx)
		return
	}
	ctx.JSON(200, workflows)
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
	groupId, ok := ctx.Get(GroupIdParam)
	if !ok {
		httputil.NewError(ctx, http.StatusBadRequest, fmt.Errorf("no group id provided"))
		return
	}
	workflowId, ok := ctx.Get(WorkflowIdParam)
	if !ok {
		httputil.NewError(ctx, http.StatusBadRequest, fmt.Errorf("no workflow id provided"))
		return
	}
	workflow, err := handler.query.GetWorkflow(groupId.(string), workflowId.(string))
	if err != nil {
		processError(err, ctx)
		return
	}
	ctx.JSON(200, workflow)
}

func processError(err error, ctx *gin.Context) {
	switch t := err.(type) {
	default:
		httputil.NewError(ctx, http.StatusInternalServerError, t)
	case *errors.BadRequestError, errors.BadRequestError:
		httputil.NewError(ctx, http.StatusBadRequest, t)
	case *errors.NotFoundError, errors.NotFoundError:
		httputil.NewError(ctx, http.StatusNotFound, t)
	}

}
