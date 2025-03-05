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
