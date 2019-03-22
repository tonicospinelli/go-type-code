package payment_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tonicospinelli/go-type-code/solution_int/payment"
	"testing"
)

var labels = map[payment.Status]string{
	payment.Initial:   "Initial",
	payment.Waiting:   "Waiting",
	payment.Paid:      "Paid",
	payment.Cancelled: "Cancelled",
}

func TestPaymentTypeStatus(t *testing.T) {
	for status, label := range labels {
		testLabel := fmt.Sprintf("should status code %d must be represents string %s status", status, label)
		t.Run(testLabel, func(t *testing.T) {
			assert.Equal(t, label, status.String(), "Failed asserting that status code has the correct label")
		})

		testLabel = fmt.Sprintf("should creates Status(%d) from Scan(%s)", status, label)
		t.Run(testLabel, func(t *testing.T) {
			var actual payment.Status
			err := actual.Scan(status.String())
			assert.NoError(t, err)
			assert.Equal(t, status, actual)
			assert.True(t, status.Equals(actual))
		})

		testLabel = fmt.Sprintf("should returns Status(%d) from Value(%s)", status, label)
		t.Run(testLabel, func(t *testing.T) {
			actual, err := status.Value()
			assert.NoError(t, err)
			assert.Equal(t, label, actual)
		})

		testLabel = fmt.Sprintf("should creates Status(%d) from MarshalText(%s)", status, label)
		t.Run(testLabel, func(t *testing.T) {
			actual, err := status.MarshalText()
			assert.NoError(t, err)
			assert.Equal(t, []byte(label), actual)
		})

		testLabel = fmt.Sprintf("should returns Status(%s) from UnmarshalText(%d)", label, status)
		t.Run(testLabel, func(t *testing.T) {
			var actual payment.Status
			err := actual.UnmarshalText([]byte(status.String()))
			assert.NoError(t, err)
			assert.Equal(t, status, actual)
		})

		testLabel = fmt.Sprintf("should creates Status(%d) from MarshalJSON(%s)", status, label)
		t.Run(testLabel, func(t *testing.T) {
			actual, err := status.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, []byte(label), actual)
		})

		testLabel = fmt.Sprintf("should returns Status(%s) from UnmarshalJSON(%d)", label, status)
		t.Run(testLabel, func(t *testing.T) {
			var actual payment.Status
			err := actual.UnmarshalJSON([]byte(status.String()))
			assert.NoError(t, err)
			assert.Equal(t, status, actual)
		})

		testLabel = fmt.Sprintf("should creates Status(%d) from MarshalBinary(%s)", status, label)
		t.Run(testLabel, func(t *testing.T) {
			actual, err := status.MarshalBinary()
			assert.NoError(t, err)
			assert.Equal(t, []byte(label), actual)
		})

		testLabel = fmt.Sprintf("should returns Status(%s) from UnmarshalBinary(%d)", label, status)
		t.Run(testLabel, func(t *testing.T) {
			var actual payment.Status
			err := actual.UnmarshalBinary([]byte(status.String()))
			assert.NoError(t, err)
			assert.Equal(t, status, actual)
		})
	}
}
