

variable "token" {
  type    = string
  default = getenv("TURSO_TOKEN")
}

env "turso" {
  url     = "libsql://tracker-aniketbiprojit.aws-eu-west-1.turso.io?authToken=${var.token}"
}

