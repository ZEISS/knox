terraform {
  backend "http" {
    address        = "http://localhost:8080/client/zeiss/demo/dev/state"
    lock_address   = "http://localhost:8080/client/zeiss/demo/dev/lock"
    unlock_address = "http://localhost:8080/client/zeiss/demo/dev/unlock"
    lock_method    = "POST"
    unlock_method  = "POST"
  }

  required_providers {
    random = {
      source  = "hashicorp/random"
      version = "~>3.0"
    }
  }
}

resource "local_file" "foo" {
  content  = "foo!"
  filename = "${path.module}/foo.bar"
}
