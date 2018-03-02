package main

import (
	"context"
	"fmt"
	"testing"

	vault "github.com/hashicorp/vault/api"

	"github.com/pachyderm/pachyderm/src/client"
	"github.com/pachyderm/pachyderm/src/client/auth"
)

const (
	vaultAddress = "http://127.0.0.1:8200"
	pachdAddress = "127.0.0.1:30650"
	pluginName   = "pachyderm"
)

func configurePlugin(t *testing.T, v *vault.Client) {

	c, err := client.NewFromAddress(pachdAddress)
	if err != nil {
		t.Errorf(err.Error())
	}
	resp, err := c.Authenticate(
		context.Background(),
		&auth.AuthenticateRequest{GitHubUsername: "admin", GitHubToken: "y"})

	fmt.Printf("login response: (%v)\n", resp)
	if err != nil {
		t.Errorf(err.Error())
	}

	vl := v.Logical()
	config := make(map[string]interface{})
	config["admin_token"] = resp.PachToken
	config["pachd_address"] = pachdAddress
	secret, err := vl.Write(
		fmt.Sprintf("/%v/config", pluginName),
		config,
	)

	fmt.Printf("config secret: %v\n", secret)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestLogin(t *testing.T) {
	c := vault.DefaultConfig()
	c.Address = vaultAddress
	v, err := vault.NewClient(c)
	if err != nil {
		t.Errorf(err.Error())
	}
	v.SetToken("root")
	// Hit login before admin token is set, expect err

	configurePlugin(t, v)

	// Now hit login endpoint w invalid vault token, expect err
	config := make(map[string]interface{})
	config["username"] = "daffyduck"
	vl := v.Logical()
	secret, err := vl.Write(
		fmt.Sprintf("/%v/login", pluginName),
		config,
	)

	fmt.Printf("config secret: %v\n", secret)
	if err != nil {
		t.Errorf(err.Error())
	}

	// Hit login w valid vault token, expect user token
	// Use client w that address / user token to list repos
}

func TestLoginTTL(t *testing.T) {
	// Same as above, validate that pach token expires after given TTL
}

func TestRenew(t *testing.T) {
	// Does login, issues renew request before TTL expires

	// Does login, issues renew request after TTL expires (expect err)
}

func TestRevoke(t *testing.T) {
	// Do normal login
	// Use user token to connect
	// Issue revoke
	// Now renewal should fail ... but that token should still work? AFAICT
}