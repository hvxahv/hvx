package mailer

type ValidateCodeCache interface {
	Set() error
	Get() error
}

type SMTP interface {
	Send() error
}
