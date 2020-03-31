package request

type (
	InsertUserRequest struct {
		Email    string
		Password []byte
	}
)
