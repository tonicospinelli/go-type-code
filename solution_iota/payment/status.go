package payment

type Status int

const (
	Initial = iota
	Waiting
	Paid
	Cancelled
)

func (s Status) Equals(status Status) bool {
	return s == status
}

func (s Status) IsClosed() bool {
	return (Paid|Cancelled)&s == 0
}
