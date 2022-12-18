package paymentstatus

type Type string

const (
	Unpaid    Type = "unpaid"
	Paid      Type = "paid"
	Cancelled Type = "cancelled"
)
