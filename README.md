# Repo Reaper

`Send unwanted repos to hell 😈🔥`

## Overview
Repo Reaper is a CLI tool designed to simplify the process of deleting multiple GitHub repositories. With an intuitive interface and interactive prompts, you can quickly and securely manage your repositories.

## Features
- Easy to Use: Simple CLI interface for quick repository deletion.
- Secure: Uses GitHub App authentication for fine-grained permissions.
- Interactive: Provides real-time feedback with emojis for success and failure.

## Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/marpit19/RepoReaper.git
    cd RepoReaper
    ```

2. **Install dependencies:**

    ```bash
    go get github.com/google/go-github/v62/github
    go get github.com/urfave/cli/v2
    go get github.com/dgrijalva/jwt-go
    ```

3. **Create a configuration folder (`Config`) and file (`config.json`):**

    ```json
    {
        "app_id": "YOUR_APP_ID",
        "installation_id": "YOUR_INSTALLATION_ID",
        "private_key_path": "path/to/your/private-key.pem"
    }
    ```

    or

    ```json
    {
        "github_token": "YOUR_GITHUB_TOKEN",
        "github_username": "YOUR_USERNAME"
    }
    ```

    Replace the placeholder values with your actual GitHub App ID, installation ID, and the path to your private key file.

4. **Add the `config.json` file to `.gitignore`:**

    ```bash
    echo "config.json" >> .gitignore
    ```

5. **Build and run the CLI:**

    ```bash
    go run cmd/main.go
    ```

