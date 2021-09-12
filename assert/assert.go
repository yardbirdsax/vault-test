package assert

import (
	"testing"

	"github.com/yardbirdsax/vault-test/client"
	"github.com/yardbirdsax/vault-test/helper"
)

// AssertVaultSecretExists asserts that a secret exists at the given path with the given key.
func AssertVaultSecretExists(t *testing.T, client client.VaultLogicalClient, path string, key string) {
	secret, err := helper.ReadVaultSecretE(client, path)
 		if err != nil {
		t.Log(err)
		t.Fail()
	}
	if secret == nil {
		t.Logf("%sSecret not found at path '%s'.%s", helper.ColorRed, path, helper.ColorReset)
		t.Fail()
		return
	}
	
	if _, exists := secret.Data[key]; !exists {
		t.Logf("%sSecret at path '%s' does not contain a key with the name '%s'.%s", helper.ColorRed, path, key, helper.ColorReset)
		t.Fail()
	}
}