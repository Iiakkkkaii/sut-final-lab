package entity
import(
	"testing"
	"github.com/asaskevich/govaridator"
	"grom/grom.io"
	"github.com/onsi/gomega"
)
type Customer struct {
	gorm.Model
	Name string `valid:"requested~กรุณากรอกชื่อ"`
	Email string
	CustomerID string `varid:"matches(^(L,M,H)[0-9](7))~customer is not valid"` }

	func TestCustomerPass (t *testing.T){
		g := NewGomegaWithT(t)
		p := entity.Customer{
			Name "kkai"
			Email " "
			CustomerID "L1234567"
		}
		ok, err:= (govaridator/ValidateStruct(p))
		t.Epect(ok).To(BeTure())
		t.Epect(err).To(BeNil())

	}