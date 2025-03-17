package test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	parallelstore "cloud.google.com/go/parallelstore/apiv1beta"
	parallelstorepb "cloud.google.com/go/parallelstore/apiv1beta/parallelstorepb"
	"google.golang.org/api/iterator"
	"testing"
	"context"
	"fmt"
	"strings"
)

// Define GCP Project and Region
const (
	projectID = "lab-gke-se"
	location  = "us-central1-a"
)

// Global Terraform options
var terraformOptions *terraform.Options

func TestParallelstoreSuite(t *testing.T) {
	// 🚀 Step 1: Terraform Setup
	terraformOptions = &terraform.Options{
		TerraformDir: "../../examples/baseline",
	}
	terraform.InitAndApply(t, terraformOptions)

	// 🚀 Step 2: Run Subtests Sequentially
	t.Run("ParallelstoreInstanceExists", testParallelstoreInstanceExists)
	t.Run("TerraformParallelStoreDefault", testTerraformParallelStoreDefault)

    // Register cleanup to ensure Terraform destroy runs last
    t.Cleanup(func() {
        terraform.Destroy(t, terraformOptions)
    })
}

// Function to fetch Parallelstore instance from GCP API
func getParallelstoreInstance(expectedName string) (string, error) {
	ctx := context.Background()

	// Initialize Parallelstore API client
	client, err := parallelstore.NewClient(ctx) // No credentials file needed!
	if err != nil {
		return "", err
	}
	defer client.Close()

	// List Parallelstore instances in the specified project and region
	req := &parallelstorepb.ListInstancesRequest{
		Parent: "projects/" + projectID + "/locations/" + location,
	}

	it := client.ListInstances(ctx, req)
	for {
		instance, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return "", err
		}

		// ✅ Return only if the instance name matches the expected name
		if strings.Contains(instance.Name, expectedName) {
			return instance.Name, nil
		}
	}

	return "", fmt.Errorf("No Parallelstore instance matching '%s' found in project %s, region %s", expectedName, projectID, location)
}

// Test if the instance exists in GCP
func testParallelstoreInstanceExists(t *testing.T) {
	expectedInstanceName := terraform.Output(t, terraformOptions, "instance_name")
	instanceName, err := getParallelstoreInstance(expectedInstanceName)
	assert.NoError(t, err, "Failed to retrieve Parallelstore instance from GCP")
	assert.NotEmpty(t, instanceName, "No Parallelstore instance found with 'ps' in its name")
}

// Test if the actual instance name from GCP match with the terraform output
func testTerraformParallelStoreDefault(t *testing.T) {
	expectedInstanceName := terraform.Output(t, terraformOptions, "instance_name")

	// Fetch actual instance name from GCP Parallelstore API
	actualInstanceName, err := getParallelstoreInstance(expectedInstanceName)
	assert.NoError(t, err, "Failed to retrieve Parallelstore instance from GCP")

	// Assert Terraform output matches the actual instance name
	assert.Equal(t, expectedInstanceName, actualInstanceName, "Parallelstore instance name does not match expected value")
}