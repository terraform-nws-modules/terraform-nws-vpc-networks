package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

const (
	vpcName = "vpc-test"
	vpcCidr = "10.0.1.0/24"
	domain  = "test.local"
)

type testCaseT struct {
	testName       string
	subnetPrivName []string
	subnetPrivCidr []string
	subnetPubName  []string
	subnetPubCidr  []string
}

func config(t *testing.T, cfg testCaseT, servicePath string) *terraform.Options {

	return terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: servicePath,
		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"vpc_name":            vpcName,
			"vpc_cidr":            vpcCidr,
			"domain":              domain,
			"subnet_private_name": cfg.subnetPrivName,
			"subnet_private_cidr": cfg.subnetPrivCidr,
			"subnet_public_name":  cfg.subnetPubName,
			"subnet_public_cidr":  cfg.subnetPubCidr,
		},
		// Add retries to eliminate trasilent errors
		MaxRetries:         3,
		TimeBetweenRetries: 5 * time.Second,
	})
}

// validates Terraform output versus expected
func validate(t *testing.T, opts *terraform.Options) {
	// outName := terraform.Output(t, opts, "name")
	// outID := terraform.Output(t, opts, "id")

	// actName := strings.Fields(trimBrackets(outName))
	// actID := strings.Fields(trimBrackets(outID))

	// // assert.Equal(t, len(name), len(actID))
	// assert.ElementsMatch(t, name, actName)
}

func genSubnetPrivateName() string {
	return fmt.Sprintf("net-private-%s", random.UniqueId())
}

func genSubnetPublicName() string {
	return fmt.Sprintf("net-public-%s", random.UniqueId())
}
