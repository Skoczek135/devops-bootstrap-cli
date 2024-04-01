variable "account_id" {
  type        = string
  description = "The AWS account ID"
}

variable "env" {
  type        = string
  description = "The environment name"
}

variable "region" {
  type        = string
  description = "The AWS region"
}

variable "vpc_id" {
  type        = string
  description = "The VPC ID"
}

variable "eks_cluster_name" {
  type        = string
  description = "Name of the eks cluster; used for provider data"
}
