package workflows

type Workflow struct {
	Name string `json:"name" minLength:"4" maxLength:"16"`
}
