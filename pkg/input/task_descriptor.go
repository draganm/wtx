package input

type Args map[string]interface{}

type TaskDescriptor struct {
	JobClass        string   `json:"job_class"`
	JobID           uint64   `json:"job_id"`
	Args            Args     `json:"args"`
	InputURLs       []string `json:"input_urls"`
	OutputURL       string   `json:"output_url"`
	CodeURL         string   `json:"code_url"`
	StatusUpdateURL string   `json:"status_update_url"`
}
