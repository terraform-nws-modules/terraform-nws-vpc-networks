package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

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

func validatePublic(t *testing.T, opts *terraform.Options, cfg testCaseT) {
	pubIP, domain, _, _ := getOutputs(t, opts)

	fmt.Println(">>>>> Pub IP: ", pubIP)

	// actName := strings.Fields(trimBrackets(outName))
	// actID := strings.Fields(trimBrackets(outID))

	assert.Equal(t, cfg.domain, domain)
	// assert.ElementsMatch(t, name, actName)
}
