# EzMail

**EzMail** is a secure, lightweight, and privacy-focused email and chat service built with **Go**. It provides a modern alternative to traditional platforms like Gmail, offering real-time messaging and email features with a unique twist — support for **anonymous access** via **OAuth SSO** (Single Sign-On).

---

## Features

- Full-featured **email system** (send, receive, filter, organize)
- Real-time **chat messaging**
- **OAuth SSO** support for easy sign-in (Google, GitHub, etc.)
- Anonymous usage without tying your identity to the platform
- High-performance backend built in **Golang**
- Mobile-friendly, responsive UI (optional frontend included)
- Persistent storage with support for multiple databases (e.g., PostgreSQL, SQLite)

---

## Getting Started

### Prerequisites

- Go 1.20+
- A supported database (PostgreSQL, SQLite, etc.)
- OAuth credentials (Google, GitHub, etc.)

### Clone and Build

```bash
git clone https://github.com/yourusername/ezmail.git
cd ezmail
go build -o ezmail ./cmd/server
```

---

## Configuration

Create a .env file or set environment variables for:

```
PORT=8080
DB_URL=postgres://user:pass@localhost:5432/ezmail
OAUTH_PROVIDER=google
OAUTH_CLIENT_ID=your-client-id
OAUTH_CLIENT_SECRET=your-client-secret
OAUTH_REDIRECT_URL=http://localhost:8080/auth/callback
ALLOW_ANONYMOUS=true
```

---

## Run the Server

```
./ezmail
```

Open http://localhost:8080 in your browser.

---

## Tech Stack

- Language: Go
- Frameworks: net/http, Gorilla Mux
- Auth: OAuth 2.0 (SSO via Google, GitHub, etc.)
- Database: GORM (supports PostgreSQL)
- Real-time Chat: WebSockets

---

## Anonymous Access via OAuth

EZMail supports anonymous access by allowing users to authenticate via OAuth providers without storing personally identifiable information (PII). Upon first login, a pseudonymous identity is generated and mapped to the OAuth token, enabling a balance between privacy and ease-of-use.

---

## API Overview

- POST /auth/login - Initiates OAuth flow
- GET /auth/callback - Handles OAuth callback
- GET /inbox - Fetches user emails
- POST /send - Sends an email
- GET /chat - Initializes WebSocket connection
- POST /chat/message - Sends a chat message

---

## Contributing

I welcome any contributions, please feel free to fork the repo, open issues and submit pull requests.

1. Fork the project
2. Create your feature branch (git checkout -b feature/foo)
3. Commit your changes (git commit -am 'Add foo')
4. Push to the branch (git push origin feature/foo)
5. Open a PR. View the [GitHub Documentation](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request) for guidance.

---

## License

EzMail is released under the [MIT License](/LICENSE)
