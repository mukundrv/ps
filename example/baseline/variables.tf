variable "region" {
  description = "region name"
  type        = string
}

variable "project_id" {
  description = "gcp project id"
  type        = string
}

############################# Existing vpc name which has PSA enabled #########################

variable "network_name" {
  description = "vpc network name which has PSA enabled"
  type        = string
}

############################# Parallelstore #############################

variable "create_instance" {
  description = "create/destroy parallestore instance"
  type        = string
}

variable "capacity_gib" {
  description = "parallestore instance size in storage n"
  type        = string
  default     = "21000"

  validation {
    condition     = var.capacity_gib >= 21000 && var.capacity_gib <= 301000 && var.capacity_gib % 7000 == 0
    error_message = "Invalid input, the instance capacity can only be configured starting from 21TiB ( 21000GiB ) up to 301TiB, with increments of 7Ti."
  }
}

variable "file_stripe_level" {
  description = "Stripe level for files"
  type        = string
  default     = "FILE_STRIPE_LEVEL_BALANCED"

  validation {
    condition     = contains(["FILE_STRIPE_LEVEL_BALANCED", "FILE_STRIPE_LEVEL_UNSPECIFIED", "FILE_STRIPE_LEVEL_MIN", "FILE_STRIPE_LEVEL_MAX"], var.file_stripe_level)
    error_message = "Invalid input, options: \"FILE_STRIPE_LEVEL_BALANCED\", \"FILE_STRIPE_LEVEL_UNSPECIFIED\", \"FILE_STRIPE_LEVEL_MIN\", \"FILE_STRIPE_LEVEL_MAX\"."
  }
}

variable "directory_stripe_level" {
  description = "Stripe level for directories"
  type        = string
  default     = "DIRECTORY_STRIPE_LEVEL_BALANCED"

  validation {
    condition     = contains(["DIRECTORY_STRIPE_LEVEL_UNSPECIFIED", "DIRECTORY_STRIPE_LEVEL_MIN", "DIRECTORY_STRIPE_LEVEL_BALANCED", "DIRECTORY_STRIPE_LEVEL_MAX"], var.directory_stripe_level)
    error_message = "Invalid input, options: \"DIRECTORY_STRIPE_LEVEL_UNSPECIFIED\", \"DIRECTORY_STRIPE_LEVEL_MIN\", \"DIRECTORY_STRIPE_LEVEL_BALANCED\", \"DIRECTORY_STRIPE_LEVEL_MAX\"."
  }
}

############################# Labels ########################## We might need to check ways of getting these values dynamically ####################

variable "tags" {
  description = "Map of labels associate to the resources"
  type        = map(string)
  default     = {}
}

variable "application_id" {
  description = "The application id"
  type        = string
  default     = ""
}

variable "environment" {
  description = "environment where the infra deployed"
  type        = string
  default     = ""
}

variable "chargeback_profile_id" {
  description = "chanrge back id for finops to charge accordingly"
  type        = string
  default     = ""
}

variable "public_cloud_support" {
  description = "Cloud support email address"
  type        = string
  default     = ""
}

variable "architecture_review_id" {
  description = "A unique id that represents architecture "
  type        = string
  default     = ""
}

variable "sector" {
  description = "Represents the business sector"
  type        = string
  default     = ""
}

variable "platform_environment" {
  description = "Environment"
  type        = string
  default     = ""
}

variable "env_code" {
  description = "The environment code"
  type        = string
  default     = ""
}
