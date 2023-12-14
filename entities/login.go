package entities

type Login struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type DataLogin struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	RoleCode string `json:"role_code"`
	Division string `json:"division"`
	Dept     string `json:"dept"`
	DeptCode string `json:"dept_code"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	ImageUrl string `json:"image_url"`
	Message  string `json:"message"`
}

type ChangePassword struct {
	UserId      string `json:"user_id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type DataChangePassword struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
