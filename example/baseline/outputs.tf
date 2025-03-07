output "instance_name" {
  description = "Name of the parallestore instance"
  value       = module.parallelstore_test.instance_name
}

output "instance_id" {
  description = "Name of the parallestore instance"
  value       = module.parallelstore_test.instance_id
}

output "instance_labels" {
  description = "Name of the parallestore instance"
  value       = module.parallelstore_test.effective_labels
}
