variable "oktapam_key" {
  type = string
  description = "Service User Key"
}
variable "oktapam_secret" {
  type = string
  description = "Service user secret"
}
variable "oktapam_team" {
  type = string
  description = "PAM Team name"
}

variable "oktapam_api_host" {
  type = string
  description = "API Host name"
  default = "https://app.scaleft.com"
}


