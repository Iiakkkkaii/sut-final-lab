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
	CustomerID string `valid:"matches(^(L,M,H)[0-9](7))~customer is not valid"` }

	func TestCustomerNamenotnull (t *testing.T){
		g := NewGomegaWithT(t)
		p := entity.Customer{
			Name ""
			Email " "
			CustomerID "L1234567"
		}
		ok, err:= govaridator.ValidateStruct(p)
		t.Epect(ok).ToNot(BeTure())
		t.Epect(err).ToNot(BeNil())
		t.Epect(err.Error()).(Equal("กรุณากรอกชื่อ"))

	}