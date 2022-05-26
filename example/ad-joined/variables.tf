variable "ad_connection_task" {
  type = object({
    hostname_attribute = string
    access_address_attribute = string
    operating_system_attribute = string
    bastion_attribute = string
    alt_names_attributes = list(string)
  })

  default = {
    hostname_attribute = "dNSHostName"
    access_address_attribute = "dNSHostName"
    operating_system_attribute = "operatingSystem"
  }
}