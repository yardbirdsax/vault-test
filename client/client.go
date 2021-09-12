/*
The Client package contains interfaces that mimic numerous Vault API types.
*/
package client

import (
	"github.com/hashicorp/vault/api"
)

// VaultLogicalClient is an interface that matches with the Vault [Logical Client](https://pkg.go.dev/github.com/hashicorp/vault/api#Logical) API type.
type VaultLogicalClient interface {
	Read(string) (*api.Secret, error)
	Delete(string) (*api.Secret, error)
}

// VaultSysClient is an interface that matches with the Vault [Sys Client](https://pkg.go.dev/github.com/hashicorp/vault/api#Sys) API type.
type VaultSysClient interface {
	Mount(string, *api.MountInput) error
}