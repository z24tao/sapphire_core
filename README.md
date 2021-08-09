# Sapphire agent core codebase

## SSH Key Setup

Create Key:
```sh
mkdir -p ~/.ssh && ssh-keygen -t ed25519 -C "email@domain.com"
```
- Press enter twice to accept the default settings

Add key to SSH agent:
```shell
ssh-add ~/.ssh/id_ed25519
```

Setup local git config:

```shell
git config --global user.name "<First Last>"
git config --global user.email "email@domain.com"
```

Set up SSH to access Github:

```shell
cat ~/.ssh/id_ed25519.pub | pbcopy
```

- Go to: https://github.com/settings/ssh

- Click `New SSH Key`

- Paste Key and click `Add SSH Key`

```zsh
# Run the following to verify that your key works:
ssh git@github.com

# Successful key authentication:  

# Hi first-last! You've successfully authenticated, but GitHub does not provide shell access.
```

## Automatic Git Commit Signing

Install gpg
```shell
brew install gpg
```

Generate gpg key with empty passphrase (recommend)
```shell
# Generate key
gpg --default-new-key-algo rsa4096 --quick-generate-key --batch --passphrase "" email@domain.com
# Export and copy to clipboard
gpg --armor --export email@domain.com | pbcopy
```
- Go to https://github.com/settings/gpg/new
- Paste your gpg key and `Add GPG Key`

Set up git to automatically sign commits
```zsh
# Set signing key to your email
git config --global user.signingkey email@domain.com
# Set gpg signing to true
git config --global commit.gpgsign true 
```

## Troubleshooting Git Commit Signature

#### Shows up as `unverified`

```zsh
gpg --list-secret-keys --keyid-format=long

# Example output:
# sec  rsa4096/12345678910
#     created: 2021-06-19  expires: 2023-06-19  usage: SC  
#     trust: ultimate      validity: ultimate
# [ultimate] (1). email@domain.com

gpg --edit-key 12345678910

# gpg> adduid to add UserID Details
adduid

# Follow the commands and add name/email/comment
# Press `O` to confirm selections

# gpg> save to save the changes
save

# Run the following command to export and copy to clipboard
gpg --armor --export 12345678910 | pbcopy

# Follow above instructions to add key to github
# Ensure that the following command displays your intended email
git config --global user.email

```