package push

type Push interface {
	Subscription()
	Send() error
}
