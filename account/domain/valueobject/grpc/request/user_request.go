package request

type (
	InsertUserRequest struct {
		Email     string
		PassHash  []byte
		Name      string
		CompanyId uint64
	}
)
