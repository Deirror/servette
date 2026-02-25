package mail

type Sender interface {
	Send(to, subject, body string) error
}
