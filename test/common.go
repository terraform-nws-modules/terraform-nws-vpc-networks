package test

import (
	"fmt"

	"github.com/gruntwork-io/terratest/modules/random"
)

const (
	vpcName = "vpc-test"
	vpcCidr = "10.0.1.0/24"
	domain  = "my.local"
	alcName = "my-acl-public"
)

type testCaseT struct {
	testName       string
	vpcName        string
	vpcCidr        string
	domain         string
	subnetPrivName []string
	subnetPrivCidr []string
	subnetPubName  []string
	subnetPubCidr  []string
	aclName        string
	aclCIDRList    []string
	aclPortList    []string
}

func genSubnetPrivateName() string {
	return fmt.Sprintf("net-private-%s", random.UniqueId())
}

func genSubnetPublicName() string {
	return fmt.Sprintf("net-public-%s", random.UniqueId())
}
