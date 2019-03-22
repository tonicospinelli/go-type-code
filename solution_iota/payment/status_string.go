package payment

import (
	"fmt"
)

var labels = [...]string{
	"Initial",
	"Waiting",
	"Paid",
	"Cancelled",
}

func (s Status) String() string {
	return labels[int(s)]
}

func Parse(status string) (Status, error) {
	for key, label := range labels {
		if label == status {
			return Status(key), nil
		}
	}
	return Initial, fmt.Errorf("status %s is not available", status)
}
