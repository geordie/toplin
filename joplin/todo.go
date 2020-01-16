package joplin

import (
	"time"
)

/* Sample json

"id": "02a1f36c306465c6936802daccc729ae",
"title": "Take down the Christmas lights",
"is_todo": 1,
"todo_completed": 1578190192571,
"parent_id": "9d59b82caa574fcaba4a5f6ebb1bdf6e",
"updated_time": 1578190192581,
"user_updated_time": 1578190192581,
"user_created_time": 1578190045427,
"encryption_applied": 0,
"type_": 1
*/

// TodoItem struct
type TodoItem struct {
	ID                string `json:"id"`
	Title             string `json:"title"`
	IsTodo            int    `json:"is_todo"`
	TodoCompletedTime int64  `json:"todo_completed"`
	ParentID          string `json:"parent_id"`
	UpdatedTime       int    `json:"updated_time"`
	UserUpdatedTime   int    `json:"user_updated_time"`
	UserCreatedTime   int    `json:"user_created_time"`
	EncryptionApplied int    `json:"encryption_applied"`
	ItemType          int    `json:"type_"`
}

// GetCompletedDate func
func (t *TodoItem) GetCompletedDate() time.Time {
	return time.Unix(0, t.TodoCompletedTime*int64(time.Millisecond))
}

// TodoItems struct
type TodoItems struct {
	MyItems []TodoItem
}
