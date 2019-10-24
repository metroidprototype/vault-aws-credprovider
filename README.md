# vault-aws-credprovider

Hashicorp Vault AWS dynamic credentials credential_process helper

## Installation

[Download](https://github.com/d3sw/metroidprototype/vault-aws-credprovider) the binary based on your OS.  Once uncompressed copy it into your system PATH.

## Usage

Once installed the following can be added to your `~/.aws/credentials` file and your profile will pull dynamic credentials from the vault path specified (*Note*: the VAULT_TOKEN environment variable needs to be set to a functional vault token on the server you are pointing to):

```
[profile-name]
region=us-west-2
credential_process=/path/to/vault-cred-helper -path aws/creds/roleName -addr http://127.0.0.1:8200
```
