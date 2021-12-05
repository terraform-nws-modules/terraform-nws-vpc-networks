package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func validatePrivate(t *testing.T, opts *terraform.Options, cfg testCaseT) {
	pubIP := terraform.Output(t, opts, "vpc_public_ip")
	domain := terraform.Output(t, opts, "domain")
	tmpPrivSubnets := terraform.Output(t, opts, "subnet_private_id")

	privSubnets := strings.Fields(trimBrackets(tmpPrivSubnets))

	fmt.Println(">>>>> Pub IP: ", pubIP)
	fmt.Println(">>>>> Domain: ", domain)
	fmt.Println(">>>>> Private subnets: ", privSubnets)

	assert.NotEmpty(t, pubIP)
	assert.Equal(t, cfg.domain, domain)
	assert.Equal(t, len(cfg.subnetPrivCidr), len(privSubnets))
}

func validatePublic(t *testing.T, opts *terraform.Options, cfg testCaseT) {
	pubIP := terraform.Output(t, opts, "vpc_public_ip")
	domain := terraform.Output(t, opts, "domain")
	tmpPubSubnets := terraform.Output(t, opts, "subnet_public_id")
	tmpACLIDs := terraform.Output(t, opts, "subnet_public_acl_id")
	tmpACLRuleIDs := terraform.Output(t, opts, "subnet_public_acl_rule_id")

	pubSubnets := strings.Fields(trimBrackets(tmpPubSubnets))
	aclIDs := strings.Fields(trimBrackets(tmpACLIDs))
	aclRuleIDs := strings.Fields(trimBrackets(tmpACLRuleIDs))

	fmt.Println(">>>>> Pub IP: ", pubIP)
	fmt.Println(">>>>> Domain: ", domain)
	fmt.Println(">>>>> Public subnets: ", pubSubnets)
	fmt.Println(">>>>> ACL IDs: ", aclIDs)
	fmt.Println(">>>>> ACL Rule IDs: ", aclRuleIDs)

	assert.NotEmpty(t, pubIP)
	assert.Equal(t, cfg.domain, domain)
	assert.Equal(t, len(cfg.subnetPubCidr), len(pubSubnets))
	assert.Equal(t, len(cfg.subnetPubCidr), len(aclIDs))
	assert.Equal(t, len(cfg.subnetPubCidr), len(aclRuleIDs))
}

func trimBrackets(s string) string {
	str0 := strings.TrimLeft(s, "[")
	str1 := strings.TrimRight(str0, "]")
	return str1
}
