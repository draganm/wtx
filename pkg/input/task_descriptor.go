package input

type TaskDescriptor struct {
	JobClass  string            `json:"job_class"`
	JobID     uint64            `json:"job_id"`
	Args      []string          `json:"args"`
	Env       map[string]string `json:"env"`
	FSURLs    map[string]string `json:"fs_urls"`
	InputURL  string            `json:"input_urls"`
	OutputURL string            `json:"output_url"`
	CodeURL   string            `json:"code_url"`
}
