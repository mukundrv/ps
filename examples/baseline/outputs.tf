output "instance_name" {
  description = "Name of the parallestore instance"
  value       = module.parallelstore_test.instance_name
}

output "instance_id" {
  description = "Name of the parallestore instance"
  value       = module.parallelstore_test.instance_id
}

output "tags" {
  description = "Name of the parallestore instance"
  value       = module.parallelstore_test.effective_labels
}

output "region" {
  description = "Region of the parallestore instance"
  value       = module.parallelstore_test.region
}

output "capacity_gb" {
  description = "Capacity of the parallestore instance"
  value       = module.parallelstore_test.capacity_gb
}

output "deployment_type" {
  description = "Deployment type of the parallestore instance"
  value       = module.parallelstore_test.deployment_type
}

output "file_stripe_level" {
  description = "File stripe level of the parallestore instance"
  value       = module.parallelstore_test.file_stripe_level
}

output "directory_stripe_level" {
  description = "Directory stripe level of the parallestore instance"
  value       = module.parallelstore_test.directory_stripe_level
}
