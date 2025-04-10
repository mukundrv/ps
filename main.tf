resource "random_string" "main" {
  length           = 8
  special          = false
  upper            = false
  override_special = ""
}

resource "google_parallelstore_instance" "main" {
  count                  = var.enable_instance ? 1 : 0
  instance_id            = "citi-ps-${var.project_id}-${var.region}-${random_string.main.id}"
  location               = var.region
  description            = var.ps_description
  capacity_gib           = var.capacity_gib
  network                = var.network_name
  file_stripe_level      = var.file_stripe_level
  directory_stripe_level = var.directory_stripe_level
  deployment_type        = "PERSISTENT"
  labels                 = local.tags
  project                = var.project_id
}
