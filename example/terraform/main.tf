variable "vault_address" {
  type = string
}

variable "vault_path" {
  type = string
}

variable "vault_secret_data" {
  type = map(string)
}

provider "vault" {
  address = var.vault_address  
}

resource "vault_generic_secret" "test" {
  path = var.vault_path
  data_json = jsonencode(var.vault_secret_data)
}