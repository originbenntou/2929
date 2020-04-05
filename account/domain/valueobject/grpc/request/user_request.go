package request

type (
	InsertUserRequest struct {
		Email    string
		PassHash []byte
	}
)
