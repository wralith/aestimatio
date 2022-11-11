package port

type Token interface {
	GenerateJWT(id, email string) (string, error)
	VerifyJWT(tkn string) bool
}
