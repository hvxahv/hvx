package mailer

type SMTPServer struct {
	from string
	to   string
	subj string
	body string
}

func (s *SMTPServer) Send() error {
	return nil
}
