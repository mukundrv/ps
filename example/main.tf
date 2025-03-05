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
  region  = var.location
}


module "test_parallelstore_module" {
  source                 = "../."
  create_instance        = var.create_instance
  capacity_gib           = var.capacity_gib
  location               = var.location
  file_stripe_level      = var.file_stripe_level
  directory_stripe_level = var.directory_stripe_level
  tags                   = var.tags
  network_name           = var.network_name
  project_id             = var.project_id
}
