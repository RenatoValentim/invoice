package usecase_test

import (
	"invoice/internal/config"
	"invoice/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateInvoice(t *testing.T) {
	assert := assert.New(t)
	config.LoadConfig(`../../`)

	t.Run(`Should to calculate invoice`, func(t *testing.T) {
		calculateInvoice := usecase.NewCalculateInvoice()
		total, _ := calculateInvoice.Execute(`1234`)

		assert.Equal(1050.0, total)
	})
}
