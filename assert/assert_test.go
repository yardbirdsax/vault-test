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
	fakeNotValue := "mybadvalue"
	listener, client := helper.CreateTestCluster(t)
	defer listener.Close()

	// Test fails if secret itself does not exist
	AssertVaultSecretExists(fakeTestSecretDoesNotExist, client.Logical(), fakePath, "", nil)
	assert.True(t, fakeTestSecretDoesNotExist.Failed())

	// Test fails if secret exists but key does not exist
	fakeTestKeyDoesNotExist := &testing.T{}
	_, err := client.Logical().Write(fakePath, map[string]interface{}{fakeKey: fakeValue})
	if err != nil {
		t.Fatal(err)
	}
	AssertVaultSecretExists(fakeTestKeyDoesNotExist, client.Logical(), fakePath, fakeMissingKey, nil)
	assert.True(t, fakeTestKeyDoesNotExist.Failed())

	// Test fails if value is provided and does not match
	fakeTestValueDoesNotMatch := &testing.T{}
	AssertVaultSecretExists(fakeTestValueDoesNotMatch, client.Logical(), fakePath, fakeKey, fakeNotValue)
	assert.True(t, fakeTestValueDoesNotMatch.Failed())

	// Test passes if value is not provided and does not match
	fakeTestValueNotProvided := &testing.T{}
	AssertVaultSecretExists(fakeTestValueNotProvided, client.Logical(), fakePath, fakeKey, nil)
	assert.False(t, fakeTestValueNotProvided.Failed())

	// Test passes if value is provided and does match
	fakeTestValueMatches := &testing.T{}
	AssertVaultSecretExists(fakeTestValueMatches, client.Logical(), fakePath, fakeKey, fakeValue)
	assert.False(t, fakeTestValueMatches.Failed())

}