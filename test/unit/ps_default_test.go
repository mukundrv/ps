package test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"testing"
)

func TestTerraformParallelStoreDefault(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../../examples/baseline",
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}