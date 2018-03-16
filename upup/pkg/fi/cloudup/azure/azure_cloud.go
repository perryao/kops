package azure

import (
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/upup/pkg/fi"
)

// AzureCloud type for Azure
type AzureCloud interface {
	fi.Cloud
}

type azureCloudImplementation struct {
}

var _ fi.Cloud = &azureCloudImplementation{}

func (c *azureCloudImplementation) ProviderID() kops.CloudProviderID {
	return kops.CloudProviderAzure
}

func NewAzureCloud() (AzureCloud, error) {

}

// fi.Cloud
// type Cloud interface {
// 	ProviderID() kops.CloudProviderID

// 	DNS() (dnsprovider.Interface, error)

// 	// FindVPCInfo looks up the specified VPC by id, returning info if found, otherwise (nil, nil)
// 	FindVPCInfo(id string) (*VPCInfo, error)

// 	// DeleteInstance deletes a cloud instance
// 	DeleteInstance(instance *cloudinstances.CloudInstanceGroupMember) error

// 	// DeleteGroup deletes the cloud resources that make up a CloudInstanceGroup, including the instances
// 	DeleteGroup(group *cloudinstances.CloudInstanceGroup) error

// 	// GetCloudGroups returns a map of cloud instances that back a kops cluster
// 	GetCloudGroups(cluster *kops.Cluster, instancegroups []*kops.InstanceGroup, warnUnmatched bool, nodes []v1.Node) (map[string]*cloudinstances.CloudInstanceGroup, error)
// }
