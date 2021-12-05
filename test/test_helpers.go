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
	domain  = "my.local"
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
		// Vars: map[string]interface{}{
		// 	"vpc_cidr":            vpcCidr,
		// 	"subnet_private_name": cfg.subnetPrivName,
		// 	"subnet_private_cidr": cfg.subnetPrivCidr,
		// 	"subnet_public_name":  cfg.subnetPubName,
		// 	"subnet_public_cidr":  cfg.subnetPubCidr,
		// },
		Vars: buildPrivateConfig(cfg),
		// Add retries to eliminate trasilent errors
		MaxRetries:         3,
		TimeBetweenRetries: 5 * time.Second,
	})
}

func buildPrivateConfig(cfg testCaseT) map[string]interface{} {
	conf := map[string]interface{}{
		"vpc_name":            vpcName,
		"vpc_cidr":            vpcCidr,
		"domain":              domain,
		"subnet_private_name": cfg.subnetPrivName,
		"subnet_private_cidr": cfg.subnetPrivCidr,
	}
	return conf
}

func buildPublicConfig(cfg testCaseT) map[string]interface{} {
	conf := map[string]interface{}{
		"vpc_cidr":           vpcCidr,
		"subnet_public_name": cfg.subnetPubName,
		"subnet_public_cidr": cfg.subnetPubCidr,
	}
	return conf
}

func buildFullConfig(cfg testCaseT) map[string]interface{} {
	conf := map[string]interface{}{
		"vpc_cidr":            vpcCidr,
		"subnet_private_name": cfg.subnetPrivName,
		"subnet_private_cidr": cfg.subnetPrivCidr,
		"subnet_public_name":  cfg.subnetPubName,
		"subnet_public_cidr":  cfg.subnetPubCidr,
	}
	return conf
}

// validates Terraform output versus expected
func getOutputs(t *testing.T, opts *terraform.Options) (pubIP, domain, privSubnets, pubSubnets string) {
	pubIP = terraform.Output(t, opts, "vpc_public_ip")
	domain = terraform.Output(t, opts, "domain")
	privSubnets = terraform.Output(t, opts, "subnet_private_id")
	pubSubnets = terraform.Output(t, opts, "subnet_public_id")

	return
}
func validate(t *testing.T, opts *terraform.Options) {
	// pubIP := terraform.Output(t, opts, "vpc_public_ip")
	// domain := terraform.Output(t, opts, "domain")
	// privSubnets := terraform.Output(t, opts, "subnet_private_id")
	// privSubnets := terraform.Output(t, opts, "subnet_private_id")
	// pubIP, domain, privSubnets, pubSubnets := getOutputs()
	pubIP, _, _, _ := getOutputs(t, opts)

	fmt.Println(">>>>> Pub IP: ", pubIP)

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

// func configEnvInputs(kind string){
// 	// func configEnvInputs(kind string) map[string]interface{}{
// 	switch kind {

// 	case "private":
// 		 return map[string]interface{}{
// 			"vpc_cidr":            vpcCidr,
// 			"subnet_private_name": cfg.subnetPrivName,
// 			"subnet_private_cidr": cfg.subnetPrivCidr,
// 			"subnet_public_name":  cfg.subnetPubName,
// 			"subnet_public_cidr":  cfg.subnetPubCidr,
// 		},

// 	}
// }
