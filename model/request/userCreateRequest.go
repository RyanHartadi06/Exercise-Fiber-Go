package request

type UserCreateRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address"  validate:"required"`
	Phone   string `json:"phone"  validate:"required"`
	Email   string `json:"email"  validate:"required"`
}
