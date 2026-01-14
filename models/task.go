package models

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}
