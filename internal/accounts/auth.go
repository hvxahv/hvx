package accounts

type AccountAuth interface {
	// Login to the account and generate token, Return token and custom error message.
	Login() (string, error)
}

func NewAccountAuth(mail string, password string) AccountAuth {
	return &Accounts{Mail: mail, Password: password}
}

func (a *Accounts) Login() (string, error) {
	name, err := NewAccountLogin(a.Mail, a.Password)
	if err != nil {
		return "", err
	}
	return name, nil
}