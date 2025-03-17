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