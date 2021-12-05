package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestVpcPrivateExample(t *testing.T) {
	t.Parallel()

	stage := test_structure.RunTestStage

	testCases := []testCaseT{
		{
			"VPC with private subnetworks",
			[]string{genSubnetPrivateName()},
			[]string{"10.0.1.0/30"},
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
				opts := config(t, testCase, servicePath)
				test_structure.SaveTerraformOptions(t, servicePath, opts)
				terraform.InitAndApply(t, opts)
			})

			defer stage(t, "destroy", func() {
				opts := test_structure.LoadTerraformOptions(t, servicePath)
				terraform.Destroy(t, opts)
			})

			stage(t, "validate", func() {
				opts := test_structure.LoadTerraformOptions(t, servicePath)
				validate(t, opts)
			})
		})
	}
}
