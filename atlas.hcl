variable "token" {
  type    = string
  default = getenv("TURSO_TOKEN")
}

variable "url" {
  type = string
  default = getenv("TURSO_URL")
}

env "turso" {
  url     = "${var.url}?authToken=${var.token}"
}

