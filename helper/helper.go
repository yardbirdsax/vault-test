/* This package contains code that makes it easy to test code that interacts with Hashicorp Vault, by enabling the creation of an in-memory Vault
cluster that is easily disposed of once testing is complete.
*/
package helper

import (
	"net"
	"testing"

	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/http"
	"github.com/hashicorp/vault/vault"

	"github.com/yardbirdsax/vault-test/client"
)

const (
	ColorRed = "\033[31m"
	ColorReset = "\033[0m"
)

/* 
CreateTestCluster creates an in-memory Vault cluster that can be used for testing code that interacts with Vault. 
*/
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

/*
MountVaultSecretEngine configures a new mount within a Vault cluster. It can be used to configure the in-memory Vault client
with other mounts beyond the default KV at the `secret/` path.
*/
func MountVaultSecretEngine(t *testing.T, client client.VaultSysClient, path string, mountType string, options map[string]string) {
	mountInput := &api.MountInput{
		Type: mountType,
		Options: options,
	}
	err := client.Mount(path, mountInput)
	if err != nil {
		t.Fatal(err)
	}
}

/*
ReadVaultSecretE reads a secret at the given path from Vault. It is basically a thin wrapper around the
built in [Read](https://pkg.go.dev/github.com/hashicorp/vault/api#Logical.Read) method, done for the purpose
of making testing of this library easier.
*/
func ReadVaultSecretE(client client.VaultLogicalClient, path string) (*api.Secret, error) {
	secret, err := client.Read(path)
	return secret, err
}

func GetVaultClientE(config *api.Config, token string) (client.VaultLogicalClient, error){
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