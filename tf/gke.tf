resource "google_container_cluster" "book_store" {
  name               = "marcellus-wallace"
  location           = var.Region
  enable_autopilot   = true
}
