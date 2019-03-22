package payment

type Status string

const (
	Initial   = Status("Initial")
	Waiting   = Status("Waiting")
	Paid      = Status("Paid")
	Cancelled = Status("Cancelled")
)

func (s Status) Equals(status Status) bool {
	return s == status
}

func (s Status) IsClosed() bool {
	return s.Equals(Paid) || s.Equals(Cancelled)
}
