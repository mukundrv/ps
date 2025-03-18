#output "instance_name" {
#  description = "Name of the parallestore instance"
#  value       = google_parallelstore_instance.main[*].instance_id
#}
#

output "instance_name" {
  description = "Name of the parallestore instance"
  value       = google_parallelstore_instance.main[0].name
}

output "instance_id" {
  description = "Name of the parallestore instance"
  value       = google_parallelstore_instance.main[0].id
}

output "effective_labels" {
  description = "Name of the parallestore instance"
  value       = google_parallelstore_instance.main[0].effective_labels
}

output "region" {
  description = "Region of the parallestore instance"
  value       = google_parallelstore_instance.main[0].location
}

output "capacity_gb" {
  description = "Capacity of the parallestore instance"
  value       = google_parallelstore_instance.main[0].capacity_gib
}

output "deployment_type" {
  description = "Deployment type of the parallestore instance"
  value       = google_parallelstore_instance.main[0].deployment_type
}

output "file_stripe_level" {
  description = "File stripe level of the parallestore instance"
  value       = google_parallelstore_instance.main[0].file_stripe_level
}

output "directory_stripe_level" {
  description = "Directory stripe level of the parallestore instance"
  value       = google_parallelstore_instance.main[0].directory_stripe_level
}
