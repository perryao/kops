package azuretasks

import (
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/upup/pkg/fi/cloudup/terraform"
)

//go:generate fitask -type=Subnet
type Subnet struct {
	Name          *string
	ID            *string
	VNet          *VNet
	AddressPrefix *string
}

var _ fi.CompareWithID = &Subnet{}

func (e *Subnet) CompareWithID() *string {
	return e.ID
}

type terraformSubnet struct {
	Name               *string            `json:"name"`
	ResourceGroupName  *terraform.Literal `json:"resource_group_name"`
	VirtualNetworkName *terraform.Literal `json:"virtual_network_name"`
	AddressPrefix      *string            `json:"address_prefix"`
}

func (_ *Subnet) RenderTerraform(t *terraform.TerraformTarget, a, e, changes *Subnet) error {
	tf := &terraformSubnet{
		Name:               e.Name,
		VirtualNetworkName: e.VNet.TerraformName(),
		AddressPrefix:      e.AddressPrefix,
	}

	return t.RenderResource("azurerm_subnet", *e.Name, tf)
}
