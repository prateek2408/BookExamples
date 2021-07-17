resource "google_sql_database_instance" "mysql" {
  name                = var.MysqlName
  database_version    = "MYSQL_8_0"
  region              = var.Region
  deletion_protection = false

  settings {
    # Second-generation instance tiers are based on the machine
    # type. See argument reference below.
    tier = "db-f1-micro"
  }
}
