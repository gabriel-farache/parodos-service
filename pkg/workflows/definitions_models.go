package workflows

import "encoding/json"

type Group struct {
	Name       string `json:"name" minLength:"4" maxLength:"16" example:"kogito-examples"`
	Repository string `json:"repository" minLength:"10" example:"https://github.com/kiegroup/kogito-examples/tree/stable"`
}

type GroupDetails struct {
	Group
	Workflows []Workflow `json:"workflows"`
}

type Workflow struct {
	Meta           map[string]string `json:"meta"`
	Name           string            `json:"name" minLength:"3" example:"fahrenheit_to_celsius"`
	InputArguments json.RawMessage   `json:"input_arguments" swaggertype:"string" example:"{ 'fahrenheit': 100 }"`
}

type WorkflowExecution struct {
	Workflow
	Result    json.RawMessage `json:"result"  example:"{ 'fahrenheit': 100, 'subtractValue': 32.0, 'multiplyValue': 0.5556, 'difference': 68.0, 'product': 37.7808 }"`
	Timestamp string          `json:"timestamp" format:"date-time"`
}
