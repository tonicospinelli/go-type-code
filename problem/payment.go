package problem

const (
	statusInitial   = "Initial"
	statusWaiting   = "Waiting"
	statusPaid      = "Paid"
	statusCancelled = "Cancelled"
)

type Payment struct {
	status string
}

// you should initiate the Payment with beginning status
func New() Payment {
	return Payment{status: statusInitial}
}

// just Payment knows about available Status, because of it
// this func and status field are protected
func (p *Payment) changeStatusTo(status string) {
	p.status = status
}

// you need to do a defensive programming to avoid invalid attribution
func (p *Payment) ChangeStatusToWaiting() {
	p.status = statusWaiting
}
