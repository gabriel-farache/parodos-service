package workflows

type WorkflowCommandHandler struct {
}

func newWorkflowCommandHandler() *WorkflowCommandHandler {
	return &WorkflowCommandHandler{}
}

func (h *WorkflowCommandHandler) CreateWorkflow() error {
	return nil
}
