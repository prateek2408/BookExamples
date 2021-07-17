variable "Project" {
  type = string
}

variable "SvcAcc" {
  type        = string
  description = "GCP service account file"
}

variable "Region" {
  type        = string
  description = "GCP region where resources would be created"
}

variable "MysqlName" {
  type        = string
  description = "The database instance name"
}
