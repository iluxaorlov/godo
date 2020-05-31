package todo

type Task struct {
	Id int `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
	Status bool `json:"status,omitempty"`
}