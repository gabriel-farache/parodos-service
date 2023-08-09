package workflows

import "github.com/parodos-dev/parodos-service/pkg/errors"

//go:generate mockgen -package=workflows -destination=mock_workflows_query.go . WorkflowsQuery
type WorkflowsQuery interface {
	ListGroups() ([]Group, error)
	GetGroup(group string) (*GroupDetails, error)
	ListWorkflows(group string) ([]Workflow, error)
	GetWorkflow(group, workflow string) (*Workflow, error)
}

type WorkflowsQueryHandler struct {
}

func newWorkflowsQueryHandler() *WorkflowsQueryHandler {
	return &WorkflowsQueryHandler{}
}

func (h *WorkflowsQueryHandler) ListGroups() ([]Group, error) {
	return []Group{
		{Name: "parodos", Repository: "https://github.com/parodos-dev/parodos"},
		{Name: "parodos-service", Repository: "https://github.com/parodos-dev/parodos-service"},
	}, nil
}

func (h *WorkflowsQueryHandler) GetGroup(group string) (*GroupDetails, error) {
	if group == "" {
		return nil, errors.BadRequestError{Message: "No group provided"}
	}
	return &GroupDetails{
		Group: Group{
			Name:       "parodos-service",
			Repository: "https://github.com/parodos-dev/parodos-service",
		},
		Workflows: []Workflow{
			{Name: "test1"},
			{Name: "test2"},
		},
	}, nil
}

func (h *WorkflowsQueryHandler) ListWorkflows(group string) ([]Workflow, error) {
	if group == "" {
		return nil, errors.BadRequestError{Message: "No group provided"}
	}
	return []Workflow{
		{Name: "test1"},
		{Name: "test2"},
	}, nil
}
func (h *WorkflowsQueryHandler) GetWorkflow(group, workflow string) (*Workflow, error) {
	if group == "" {
		return nil, errors.BadRequestError{Message: "No group provided"}
	}
	if workflow == "" {
		return nil, errors.BadRequestError{Message: "No workflow provided"}
	}
	return &Workflow{Name: "test1"}, nil
}
