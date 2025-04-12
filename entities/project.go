package entities

import "time"

type ProjectEntity struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	ProjectType string `json:"project_type"`
	CreatedUserId string `json:"created_user_id"`
	CreatedAt time.Time `json:"created_at"`
}