package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hashicorp/vault/api"
)

const version = "0.1.0"

var (
	path      = flag.String("path", "", "Vault Dynamic Secret Path")
	vaultAddr = flag.String("addr", "http://127.0.0.1:8200", "Vault Address")
	token     = os.Getenv("VAULT_TOKEN")
)

// Credentials is the creds object for AWS credentials_process
type Credentials struct {
	Version         int
	AccessKeyId     string
	SecretAccessKey string
	SessionToken    string
	Expiration      string
}

func main() {
	flag.Parse()
	config := api.DefaultConfig()
	config.Address = *vaultAddr

	cli, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create vault client: %s", err)
	}
	cli.SetToken(token)
	c := cli.Logical()

	secret, err := c.Read(*path)
	if err != nil {
		log.Fatal(err)
	}

	exp := time.Now().Add(time.Second * time.Duration(secret.LeaseDuration)).Format(time.RFC3339)

	creds := Credentials{
		Version:         1,
		AccessKeyId:     secret.Data["access_key"].(string),
		SecretAccessKey: secret.Data["secret_key"].(string),
		SessionToken:    secret.Data["security_token"].(string),
		Expiration:      exp,
	}

	output, _ := json.Marshal(creds)
	fmt.Println(string(output))
}
