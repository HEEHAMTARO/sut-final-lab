package test

import (
	"fmt"
	"testing"

	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"

	"example.com/sut-final-lab/entity"
)

func TestSKU(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("SKU is required", func(t *testing.T){
		product := entity.Products{
			Name: "Nay",
			Price: 102,
			SKU: "",
		}

		ok, err := govalidator.ValidateStruct(product)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("SKU is required"))
	})

	t.Run("SKU is invalid", func(t *testing.T){
		product := entity.Products{
			Name: "Nay",
			Price: 123333,
			SKU: "ABC1234564",
		}

		ok, err := govalidator.ValidateStruct(product)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal(fmt.Sprintf("SKU: %d does not validate as match(^[A-Z]\\d{5}$)", product.SKU)))
	})
}