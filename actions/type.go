package actions

import "time"

type Item struct {
	Task string `json:"todo"`
	Done bool `json:"done"`
	Created_at time.Time `json:"created_at"`
	Completed_at time.Time `json:"completed_at"`
}

type Todos []Item