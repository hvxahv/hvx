package cache

type cache interface {
	Dial(db int) error
}
