provider "google" {
  project     = var.Project
  credentials = file(var.SvcAcc)
  region      = "us-west1"
}
