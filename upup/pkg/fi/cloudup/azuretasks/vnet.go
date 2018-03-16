package azuretasks

import (
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/upup/pkg/fi/cloudup/terraform"
)

//go:generate fitask -type=VNet
type VNet struct {
	Name         *string
	ID           *string
	AddressSpace *[]string
}

var _ fi.CompareWithID = &VPC{}

func (e *VNet) CompareWithID() *string {
	return e.ID
}

type terraformVNet struct {
	AddressSpace *string `json:"address_space,omitempty"`
}

func (_ *VNet) RenderTerraform(t *terraform.TerraformTarget, a, e, changes *VNet) error {
	tf := &terraformVNet{
		AddressSpace: e.AddressSpace,
	}

	return t.RenderResource("azurerm_virtual_network", *e.Name, tf)
}

func (e *VNet) TerraformName() *terraform.Literal {
	return terraform.LiteralProperty("azurerm_virtual_network", *e.Name, "name")
}
