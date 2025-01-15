package test

import (
	"fmt"
	"testing"

	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"

	"example.com/sut-final-lab/entity"
)

func TestPrice(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("Price is required", func(t *testing.T){
		product := entity.Products{
			Name: "Nay",
			Price: 0,
			SKU: "ABC123456",
		}

		ok, err := govalidator.ValidateStruct(product)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Price is required"))
	})

	t.Run("Price is invalid", func(t *testing.T){
		product := entity.Products{
			Name: "Nay",
			Price: 123333,
			SKU: "ABC123456",
		}

		ok, err := govalidator.ValidateStruct(product)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal(fmt.Sprintf("Price: %.0f does not validate as range(1|1000)", product.Price)))
	})
}