terraform {
  required_version = ">=1.6"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 6.15.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.3.2"
    }
  }
}


provider "google" {
  project = var.project_id
  region  = var.region
}


module "parallelstore_test" {
  source                 = "../.."
  enable_instance        = var.enable_instance
  ps_description         = var.ps_description
  capacity_gib           = var.capacity_gib
  region                 = var.region
  file_stripe_level      = var.file_stripe_level
  directory_stripe_level = var.directory_stripe_level
  tags                   = var.tags
  network_name           = var.network_name
  project_id             = var.project_id
}
