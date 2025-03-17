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
	"strconv"
)

// Define GCP Project and Region
const (
	projectID = "lab-gke-se"
	location  = "us-central1-a"
)

type ParallelstoreInstanceDetails struct {
	Name            		string
	Region          		string
	DeploymentType  		string
	CapacityGb      		string
	FileStripeLevel 		string
	DirectoryStripeLevel    string
	Labels		  			string
}

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
	t.Run("TerraformParallelStoreLabels", testTerraformParallelStoreLabels)
	t.Run("TerraformParallelStoreCapacity", testTerraformParallelStoreCapacity)
	t.Run("TerraformParallelStoreStripeLevel", testTerraformParallelStoreStripeLevel)

    // Register cleanup to ensure Terraform destroy runs last
    t.Cleanup(func() {
        terraform.Destroy(t, terraformOptions)
    })
}

// Extracts the region from the instance's Parent field
func extractLocation(parent string) string {
	parts := strings.Split(parent, "/")
	for i, part := range parts {
		if part == "locations" && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return "unknown-region"
}

// Function to fetch Parallelstore instance from GCP API
func getParallelstoreInstance(expectedName string) (*ParallelstoreInstanceDetails, error) {
	ctx := context.Background()

    // Initialize Parallelstore API client
    client, err := parallelstore.NewClient(ctx)
    if err != nil {
        return nil, err
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
            return nil, err
        }

		// ✅ Return only if the instance name matches the expected name
		if strings.Contains(instance.Name, expectedName) {
			instanceDetails := &ParallelstoreInstanceDetails{
				Name:            		instance.Name,
				Region:          		extractLocation(instance.Name),
				DeploymentType:  		instance.DeploymentType.String(), // Adjust this field according to API response
				CapacityGb:      		strconv.FormatInt(instance.CapacityGib, 10), // Adjust this field according to API response
				FileStripeLevel: 		instance.FileStripeLevel.String(), // Adjust as necessary
				DirectoryStripeLevel:  	instance.DirectoryStripeLevel.String(), // Adjust as necessary
				Labels:          		fmt.Sprintf("%v", instance.Labels),
			}
			return instanceDetails, nil
		}
	}

	return nil, fmt.Errorf("No Parallelstore instance matching '%s' found in project %s, region %s", expectedName, projectID, location)
}

// Test if the instance exists in GCP
func testParallelstoreInstanceExists(t *testing.T) {
	expectedInstanceName := terraform.Output(t, terraformOptions, "instance_name")
	instanceDetails, err := getParallelstoreInstance(expectedInstanceName)
	assert.NoError(t, err, "Failed to retrieve Parallelstore instance from GCP")
	assert.NotNil(t, instanceDetails, "No Parallelstore instance found matching Terraform output")
}

// Test if the actual instance name & deployment type from GCP match with the terraform output
func testTerraformParallelStoreDefault(t *testing.T) {
	expectedInstanceName := terraform.Output(t, terraformOptions, "instance_name")
	expectedRegion := terraform.Output(t, terraformOptions, "region")
	expectedDeploymentType := terraform.Output(t, terraformOptions, "deployment_type")

	// Fetch actual instance name from GCP Parallelstore API
	instanceDetails, err := getParallelstoreInstance(expectedInstanceName)
	assert.NoError(t, err, "Failed to retrieve Parallelstore instance from GCP")

	// Assert Terraform output matches the actual instance name
	assert.Equal(t, expectedInstanceName, instanceDetails.Name, "Parallelstore instance name does not match expected value")
	// Assert Terraform output matches the actual instance name
	assert.Equal(t, expectedRegion, instanceDetails.Region, "Parallelstore instance region does not match expected value")
	// Assert Terraform output matches the actual deployment type
	assert.Equal(t, expectedDeploymentType, instanceDetails.DeploymentType, "Parallelstore instance deployment type does not match expected value")
}

// Test if the actual instance labels from GCP match with the terraform output
func testTerraformParallelStoreLabels(t *testing.T) {
	expectedInstanceName := terraform.Output(t, terraformOptions, "instance_name")
	expectedLabels := terraform.Output(t, terraformOptions, "tags")

	// Fetch actual instance name from GCP Parallelstore API
	instanceDetails, err := getParallelstoreInstance(expectedInstanceName)
	assert.NoError(t, err, "Failed to retrieve Parallelstore instance from GCP")

	// Check if the labels are empty
	assert.NotEmpty(t, instanceDetails.Labels, "Parallelstore instance labels are empty")
	// Assert Terraform output matches the actual instance name
	assert.Equal(t, expectedLabels, instanceDetails.Labels, "Parallelstore instance labels does not match expected value")
}

// Test if the actual instance capacity from GCP match with the terraform output
func testTerraformParallelStoreCapacity(t *testing.T) {
	expectedInstanceName := terraform.Output(t, terraformOptions, "instance_name")
	expectedCapacity := terraform.Output(t, terraformOptions, "capacity_gb")

	// Fetch actual instance name from GCP Parallelstore API
	instanceDetails, err := getParallelstoreInstance(expectedInstanceName)
	assert.NoError(t, err, "Failed to retrieve Parallelstore instance from GCP")

	// Assert Terraform output matches the actual instance name
	assert.Equal(t, expectedCapacity, instanceDetails.CapacityGb, "Parallelstore instance capacity settings does not match expected value")
}

// Test if the actual instance file/directory stripe level from GCP match with the terraform output
func testTerraformParallelStoreStripeLevel(t *testing.T) {
	expectedInstanceName := terraform.Output(t, terraformOptions, "instance_name")
	expectedFileStripeLevel := terraform.Output(t, terraformOptions, "file_stripe_level")
	expectedDirectoryStripeLevel := terraform.Output(t, terraformOptions, "directory_stripe_level")

	// Fetch actual instance name from GCP Parallelstore API
	instanceDetails, err := getParallelstoreInstance(expectedInstanceName)
	assert.NoError(t, err, "Failed to retrieve Parallelstore instance from GCP")

	// Assert Terraform output matches the actual instance name
	assert.Equal(t, expectedFileStripeLevel, instanceDetails.FileStripeLevel, "Parallelstore instance file stripe level settings does not match expected value")
	assert.Equal(t, expectedDirectoryStripeLevel, instanceDetails.DirectoryStripeLevel, "Parallelstore instance directory stripe level settings does not match expected value")
}