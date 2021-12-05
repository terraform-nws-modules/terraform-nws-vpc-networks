package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

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

// validates Terraform output versus expected
func getOutputs(t *testing.T, opts *terraform.Options) (pubIP, domain, privSubnets, pubSubnets string) {
	pubIP = terraform.Output(t, opts, "vpc_public_ip")
	domain = terraform.Output(t, opts, "domain")
	privSubnets = terraform.Output(t, opts, "subnet_private_id")
	pubSubnets = terraform.Output(t, opts, "subnet_public_id")

	return
}

func validatePrivate(t *testing.T, opts *terraform.Options, cfg testCaseT) {
	pubIP, domain, _, _ := getOutputs(t, opts)

	fmt.Println(">>>>> Pub IP: ", pubIP)

	// actName := strings.Fields(trimBrackets(outName))
	// actID := strings.Fields(trimBrackets(outID))

	assert.Equal(t, cfg.domain, domain)
	// assert.ElementsMatch(t, name, actName)
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
