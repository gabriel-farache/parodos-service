package workflows

type WorkflowQueryHandler struct {
}

func newWorkflowQueryHandler() *WorkflowQueryHandler {
	return &WorkflowQueryHandler{}
}

func (h *WorkflowQueryHandler) List() ([]Workflow, error) {
	return nil, nil
}
