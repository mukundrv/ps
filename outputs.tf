#output "instance_name" {
#  description = "Name of the parallestore instance"
#  value       = google_parallelstore_instance.main[*].instance_id
#}
#

output "instance_name" {
  description = "Name of the parallestore instance"
  value       = google_parallelstore_instance.main[*].name
}

output "instance_id" {
  description = "Name of the parallestore instance"
  value       = google_parallelstore_instance.main[*].id
}

output "effective_labels" {
  description = "Name of the parallestore instance"
  value       = google_parallelstore_instance.main[*].effective_labels
}

