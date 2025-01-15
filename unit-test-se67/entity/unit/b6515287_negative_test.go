package unit

import (
	"githib/unit-test-se67/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"testing"
)

func TestPrice(t *testing.T) {
	g := NewGomegaWiithT(t)
	t.Run(`Price must be between 1 and 1000`, func(t *testing.T) {
		products = entity.Products{
			Name:  "SE",
			Price: "105",
		}
		ok, err := govalidator.ValidateStruct(products)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Price must be between 1 and 1000"))
	})
	t.Run(`Price is valid`, func(t *testing.T) {
		products = entity.Products{
			Name:  "SE",
			Price: "99",
		}
		ok, err := govalidator.ValidateStruct(products)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

	})

}
