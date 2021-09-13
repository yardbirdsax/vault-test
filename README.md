# vault-test

When writing [Terraform](https://terraform.io) modules that interact with Hashicorp [Vault](https://vaultproject.io), I found it time consuming to set up a Vault cluster simply for the purposes of running my automated tests. I began looking for ways to easily set up a local instance of Vault, and in the process discovered that Vault's library contains a very easy way to do this directly in Go code! This library makes the process of using that functionality simple and also contains some methods to assert that secrets exist. And although its original purpose was mainly in the context of testing Terraform code, it could certainly be used for other things, such as running integration tests for Golang code that interacts with Vault.

## General methods

When you need to set up your test, you can create the test cluster (and defer its destruction) like this:

```golang
package something

import (
	"testing"

	"github.com/yardbirdsax/vault-test/helper"
)

func TestAssertVaultSecretExists(t *testing.T) {
	// Setup
	listener, client := helper.CreateTestCluster(t)
	defer listener.Close()

    // Assert stuff here
}
```

## Use in concert with Terratest

If you want to test Terraform code that interacts with Vault, there are a couple of things you can do to make this easier.

* Make the vault URL an input variable, like this:
  ```hcl
  variable "vault_address" {
    type = string
  }

  provider "vault" {
    address = var.vault_address
  }
  ```
* When running the `plan` or `apply` steps through your Go test, make sure you specify this variable with a value given by the `listener` object returned from creating the Vault test cluster, and set the VAULT_TOKEN environment variable from the `client` object. Here's an example:
  
  ```golang
  vaultPath := "secret/mysecret"
  vaultKey := "mykey"
  vaultSecret := "mysecretvalue"
  vaultURL := fmt.Sprintf("http://%s", listener.Addr())

  terraformDir := "terraform/"
  terraformOptions := &terraform.Options{
    TerraformDir: terraformDir,
    Vars: map[string]interface{}{
      "vault_address": vaultURL,
      "vault_path": vaultPath,
      "vault_secret_data": map[string]string{
        vaultKey: vaultSecret,
      },
    },
    EnvVars: map[string]string{
      "VAULT_TOKEN": client.Token(),
    },
  }
  ```

For a complete example of how to do this, see the [`example` directory](example/) of this repo.