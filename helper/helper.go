package helper

import (
	"net"
	"testing"



	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/http"
	"github.com/hashicorp/vault/vault"
)

const (
	colorRed = "\033[31m"
	colorReset = "\033[0m"
)

func CreateTestCluster(t *testing.T) (net.Listener, *api.Client) {
	t.Helper()

	core, _, rootToken := vault.TestCoreUnsealed(t)

	listener, address := http.TestServer(t, core)

	config := api.DefaultConfig()
	config.Address = address

	client, err := api.NewClient(config)
	if (err != nil) {
		t.Fatal(err)
	}
	client.SetToken(rootToken)

	return listener, client
}

func MountVaultSecretEngine(t *testing.T, client *api.Sys, path string) {
	mountInput := &api.MountInput{
		Type: "kv",
		Options: map[string]string{ "version": "2"},
	}
	err := client.Mount(path, mountInput)
	if err != nil {
		t.Fatal(err)
	}
}

type VaultLogicalClient interface {
	Read(string) (*api.Secret, error)
	Delete(string) (*api.Secret, error)
}

// AssertVaultSecretExists asserts that a secret exists at the given path with the given key.
func AssertVaultSecretExists(t *testing.T, client VaultLogicalClient, path string, key string) {
	secret, err := ReadVaultSecretE(client, path)
 		if err != nil {
		t.Log(err)
		t.Fail()
	}
	if secret == nil {
		t.Logf("%sSecret not found at path '%s'.%s", colorRed, path, colorReset)
		t.Fail()
		return
	}
	
	if _, exists := secret.Data[key]; !exists {
		t.Logf("%sSecret at path '%s' does not contain a key with the name '%s'.%s", colorRed, path, key, colorReset)
		t.Fail()
	}
}

func DeleteVaultSecret(t *testing.T, client VaultLogicalClient, path string) {
	_, err := client.Delete(path)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func ReadVaultSecretE(client VaultLogicalClient, path string) (*api.Secret, error) {
	secret, err := client.Read(path)
	return secret, err
}

func GetVaultClientE(config *api.Config, token string) (VaultLogicalClient, error){
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	if token != "" {
		client.SetToken(token)
	}
	logicalClient := client.Logical()
	return logicalClient, nil
}