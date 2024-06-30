package driven

type PasswordService interface {
	HashPassword(string) (string, error)
}
