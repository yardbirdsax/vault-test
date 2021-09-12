package example

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"

	"github.com/yardbirdsax/vault-test/assert"
	"github.com/yardbirdsax/vault-test/helper"
)

func TestVaultSecret (t *testing.T) {
	listener, client := helper.CreateTestCluster(t)
	defer test_structure.RunTestStage(t, "vault_destroy", func() {
		listener.Close()
	})
	
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

	defer test_structure.RunTestStage(t, "terraform_destroy", func() {
		_ = terraform.Destroy(t, terraformOptions)
	})
	test_structure.RunTestStage(t, "terraform_apply", func() {
		_ = terraform.InitAndApplyAndIdempotent(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "vault_test", func() {
		assert.AssertVaultSecretExists(t, client.Logical(), vaultPath, vaultKey, vaultSecret)
	})
}