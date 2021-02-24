[![GitHub release (latest by date)](https://img.shields.io/github/v/release/abergmeier/terraform-provider-dbussecretservice)](https://github.com/abergmeier/terraform-provider-dbussecretservice/releases/latest)
[![License](https://img.shields.io/github/license/abergmeier/terraform-provider-dbussecretservice)](https://github.com/abergmeier/terraform-provider-dbussecretservice/blob/master/LICENSE)

# terraform-provider-dbussecretservice

## Example use case

1. Store a Secret in a SecretService implementation
   
   Assigns identifiers, key `env` with value `dev` and key `foo` with value `bar`.

  ```bash
  echo -n MyPasswort | secret-tool store --label 'My Password' env dev foo bar
  ```
  
2. Declare provider in Terraform

   Add new provider `dbussecretservice` as requirement. 

  ```terraform
  terraform {
    required_providers {
      dbussecretservice = {
        source = "abergmeier/dbussecretservice"
    }
  }
  ```

3. Declare Secret DataSource

   Add a new secret declaration.

  ```terraform
  data "dbussecretservice_login" "my_password" {
    attributes = {
      "env" = "dev"
      "foo" = "bar"
    }
  }
  ```
  
4. Use Secret value

   Reference Secret value where necessary 

  ```terraform
  resource "hu_ha" "my_resource" {
    password = data.dbussecretservice_login.my_password.value
  }
  ```
  
5. Install provider in Terraform

  ```bash
  terraform init -upgrade
  ```
6. Deployment in Terraform

  ```bash
  terraform apply
  ```
