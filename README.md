# go-utils


## Setting Up `.gitconfig` for Private Repository
To enable seamless access to a private repository, configure Git to use either SSH or HTTPS for authentication:

### Configure SSH Access
For SSH, run the following command to replace `https://github.com/` with SSH access:
```bash
git config --global url."ssh://git@github.com/".insteadOf "https://github.com/"
```
### Configure HTTPS Access
For HTTPS access using a personal access token, replace `<github_token>` with your actual token:
```bash
git config --global url."https://<github_token>@github.com/".insteadOf "https://github.com/"
```
## Getting a GitHub Personal Access Token (`<github_token>`)
### 1. Go to GitHub Settings:

 - Navigate to GitHub and sign in to your account.
 - Click on your profile picture in the upper-right corner and select **Settings**.

### 2. Create a Personal Access Token:

 - In the left sidebar, click on **Developer settings**.
 - Click on **Personal access tokens**, then click the **Generate new token** button.
 - Select the required **scopes** (permissions) for your token:
    - **repo**: Full access to private repositories. This includes read and write access to code, issues, and pull requests.
    - **read**
(optional): To read organization membership.
    - **workflow** (optional): To interact with GitHub Actions workflows.
    - **delete_repo** (optional): If you want to delete repositories using the token.

### 3. Generate and Copy the Token:

 - After selecting the scopes, click **Generate token**.
 - Copy the token immediately, as you won't be able to see it again.

## Using `go-utils` in Your Project
To include `go-utils` from a private GitHub repository, follow these steps:

### 1. Set the GOPRIVATE Environment Variable

Configure Go to recognize the repository as private:

```bash
go env -w GOPRIVATE=github.com/repooooo
```

### 2. Install go-utils

Use the `go get` command to add `go-utils` to your project dependencies:

```bash
go get -u github.com/repooooo/go-utils
```