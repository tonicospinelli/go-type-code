package solution_string

import "github.com/tonicospinelli/go-type-code/solution_string/payment"

type Payment struct {
	status payment.Status
}

func New() Payment {
	return Payment{status: payment.Initial}
}

// now you may do less defensive code, because you are not able
// to make invalid attribution
func (p *Payment) ChangeStatusTo(status payment.Status) {
	p.status = status
}
