package project

// project entity map to table
type ProjectEntity struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	ProjectType string `json:"project_type"`
	CreatedUserId string `json:"created_user_id"`
	CreatedAt string `json:"created_at"`
}

// project create req dto
type ProjectCreateReqDto struct {
	Name string `json:"name" binding:"required,min=3,max=50"`
	Description string `json:"description" binding:"required,min=3,max=355"`	
	ProjectType string `json:"projectType" binding:"required,min=3,max=100"`	
	CreatedUserId string `json:"createdUserId" binding:"required,min=3,max=255"`	
}

// project create input
type ProjectCreateInput struct {
	*ProjectCreateReqDto
	Id string
}