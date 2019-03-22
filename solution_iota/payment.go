package solution_iota

import "github.com/tonicospinelli/go-type-code/solution_iota/payment"

type Payment struct {
	status payment.Status
}

func New() Payment {
	return Payment{status: payment.Initial}
}

// now you may do less defensive code, because you are not able
// to make invalid attribution
//
// NOTE: this file has no difference from solution_string file :)
func (p *Payment) ChangeStatusTo(status payment.Status) {
	p.status = status
}
