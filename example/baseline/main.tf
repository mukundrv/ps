terraform {

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">=6.15.0"
    }
    random = {
      source  = "hashicorp/random"
      version = ">=3.3.2"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
}

module "paralleltstore_test_baseline" {
  source                 = "../.."
  project_id             = var.project_id
  create_instance        = var.create_instance
  region                 = var.region
  capacity_gib           = var.capacity_gib
  network_name           = var.network_name
  file_stripe_level      = var.file_stripe_level
  directory_stripe_level = var.directory_stripe_level
  tags                   = var.tags
}