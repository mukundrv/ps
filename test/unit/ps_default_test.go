package test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTerraformParallelStoreDefault(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../../examples/baseline",
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	instanceNameDefault := terraform.Output(t, terraformOptions, "instance_name")
	instanceIDDefault := terraform.Output(t, terraformOptions, "instance_id")
	instanceLabelDefault := terraform.Output(t, terraformOptions, "tags")
	instanceRegion := terraform.Output(t, terraformOptions, "region")
	instanceCapacity := terraform.Output(t, terraformOptions, "capacity_gb")

	assert.NotEmpty(t, instanceNameDefault, "ParallelStore instance name should not be empty")
	assert.NotEmpty(t, instanceIDDefault, "ParallelStore ID should not be empty")
	assert.NotEmpty(t, instanceLabelDefault, "ParallelStore label should not be empty")
	assert.Equal(t, "[us-central1-a]", instanceRegion, "Region should be us-central1")
	assert.Equal(t, "[27000]", instanceCapacity, "Capacity should be 27000GB")
}