package assert

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yardbirdsax/vault-test/helper"
)

func TestAssertVaultSecretExists(t *testing.T) {
	// Setup
	fakeTestSecretDoesNotExist := &testing.T{}
	fakePath := "secret/somesecret"
	fakeKey := "mysecret"
	fakeMissingKey := "myothersecret"
	fakeValue := "myvalue"
	listener, client := helper.CreateTestCluster(t)
	defer listener.Close()

	// Test fails if secret itself does not exist
	AssertVaultSecretExists(fakeTestSecretDoesNotExist, client.Logical(), fakePath, "")
	assert.True(t, fakeTestSecretDoesNotExist.Failed())

	// Test fails if secret exists but key does not exist
	fakeTestKeyDoesNotExist := &testing.T{}
	_, err := client.Logical().Write(fakePath, map[string]interface{}{fakeKey: fakeValue})
	if err != nil {
		t.Fatal(err)
	}
	AssertVaultSecretExists(fakeTestKeyDoesNotExist, client.Logical(), fakePath, fakeMissingKey)
	assert.True(t, fakeTestKeyDoesNotExist.Failed())

}