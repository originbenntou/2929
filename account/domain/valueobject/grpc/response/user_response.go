package response

type (
	InsertUserResponse struct {
		Id           uint64
		Email        string
		PasswordHash []byte
		Name         string
		CompanyId    uint64
	}
)
