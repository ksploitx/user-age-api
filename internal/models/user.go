package models

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	DOB  string `json:"dob"  validate:"required"` // "YYYY-MM-DD"
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	DOB  string `json:"dob"  validate:"required"`
}

type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
}

type UserDetailResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age"`
}
