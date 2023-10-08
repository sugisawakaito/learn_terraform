# -----------------------------------------------------
# Terraform configuration file for AWS
# -----------------------------------------------------

terraform {
  required_version = ">= 0.13.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 3.0.0"
    }
  }
  backend "s3" {
    bucket  = "tastylog-tfstate-bucket-kaito"
    key     = "tastylog-dev.tfstate"
    region  = "ap-northeast-1"
    profile = "terraform"
  }
}

# -----------------------------------------------------
# Provider configuration
# -----------------------------------------------------

provider "aws" {
  profile = "terraform"
  region  = "ap-northeast-1"
}

# -----------------------------------------------------
# Variables
# -----------------------------------------------------

variable "project" {
  type = string
}

variable "environment" {
  type = string
}


