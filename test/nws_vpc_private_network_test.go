package test

import (
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestVpcPrivateExample(t *testing.T) {
	t.Parallel()

	stage := test_structure.RunTestStage

	testCases := []testCaseT{
		{
			"VPC with private subnetworks",
			vpcName,
			vpcCidr,
			domain,
			[]string{genSubnetPrivateName()},
			[]string{"10.0.1.0/30"},
			[]string{},
			[]string{},
			alcName,
			[]string{},
			[]string{},
		},
	}
	for _, testCase := range testCases {

		// capture range variable so that it doesn't update when the subtest goroutine swaps.
		testCase := testCase

		// generate a random service path for each parallel test
		servicePath := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/vpc-private")

		// fork a parallel test with all stages
		t.Run(testCase.testName, func(t *testing.T) {
			t.Parallel()
			stage(t, "deploy", func() {
				opts := configPrivate(t, testCase, servicePath)
				test_structure.SaveTerraformOptions(t, servicePath, opts)
				terraform.InitAndApply(t, opts)
			})

			defer stage(t, "destroy", func() {
				opts := test_structure.LoadTerraformOptions(t, servicePath)
				terraform.Destroy(t, opts)
			})

			stage(t, "validate", func() {
				opts := test_structure.LoadTerraformOptions(t, servicePath)
				validatePrivate(t, opts, testCase)
			})
		})
	}
}

func configPrivate(t *testing.T, cfg testCaseT, servicePath string) *terraform.Options {

	return terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: servicePath,

		// Variables to pass to our Terraform code using -var options
		Vars: buildPrivateConfig(cfg),
		// Add retries to eliminate trasilent errors
		MaxRetries:         3,
		TimeBetweenRetries: 5 * time.Second,
	})
}
