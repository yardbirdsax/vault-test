package assert

import (
	"testing"

	"github.com/yardbirdsax/vault-test/client"
	"github.com/yardbirdsax/vault-test/helper"
)

// AssertVaultSecretExists asserts that a secret exists at the given path with the given key and value. If the value input 
// is set to nil, then it will be ignored and the test will pass so long as a secret exists at the given
// path and contains the given key.
func AssertVaultSecretExists(t *testing.T, client client.VaultLogicalClient, path string, key string, value interface{}) {
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
	
	secretValue, exists := secret.Data[key]
	if !exists {
		t.Logf("%sSecret at path '%s' does not contain a key with the name '%s'.%s", helper.ColorRed, path, key, helper.ColorReset)
		t.Fail()
	}

	if value != nil && secretValue != value {
		t.Logf("%sSecret value at path '%s' and key '%s' does not match. Expected: %v; actual: %v.%s", helper.ColorRed, path, key, value, secretValue, helper.ColorReset)
		t.Fail()
	}
}