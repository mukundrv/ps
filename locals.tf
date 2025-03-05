locals {
  tags = merge(
    {
      application_id         = var.application_id
      environment            = var.environment
      chargeback_profile_id  = var.chargeback_profile_id
      public_cloud_support   = var.public_cloud_support
      architecture_review_id = var.architecture_review_id
      sector                 = var.sector
      platform_environment   = var.platform_environment
      env_code               = var.env_code
      module_name            = "gcp-parallelstore"
      module_type            = "Terraform"
      cte_version_id         = "108744"
    },
    var.tags
  )
}
