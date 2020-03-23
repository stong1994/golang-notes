package builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	manufacturingDirector := ManufacturingDirector{}
	laptop := &Laptop{}
	manufacturingDirector.SetBuilder(laptop)
	manufacturingDirector.Construct()
	manufacturingDirector.PrintProduct()

	phone := &SmartPhone{}
	manufacturingDirector.SetBuilder(phone)
	manufacturingDirector.Construct()
	manufacturingDirector.PrintProduct()
}
