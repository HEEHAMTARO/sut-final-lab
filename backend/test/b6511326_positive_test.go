package test

import (
	"testing"

	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"

	"example.com/sut-final-lab/entity"
)

func TestPrice(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("valid", func(t *testing.T){
		product := entity.Products{
			Name: "Nay",
			Price: 20.3,
			SKU: "ABC123456",
		}

		ok, err := govalidator.ValidateStruct(product)

		g.Expect(ok).To(BeTrue())

		g.Expect(err).To(BeNil())
	})
}