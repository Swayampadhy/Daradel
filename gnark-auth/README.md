# gnark-auth

## Overview
The `gnark-auth` project implements an authentication system using zero-knowledge proofs with the Gnark library. It consists of an agent and a server that communicate securely to authenticate users without revealing sensitive information.

## Project Structure
```
gnark-auth
├── cmd
│   ├── agent
│   │   └── main.go        # Entry point for the agent application
│   └── server
│       └── main.go        # Entry point for the server application
├── internal
│   └── auth
│       ├── circuit.go     # Cryptographic circuit for authentication
│   ├── agent.go           # Agent's authentication logic
│   └── server.go          # Server's authentication logic
├── go.mod                  # Module definition and dependencies
└── README.md               # Project documentation
```

## Setup Instructions
1. **Clone the repository:**
   ```
   git clone https://github.com/yourusername/gnark-auth.git
   cd gnark-auth
   ```

2. **Install dependencies:**
   Ensure you have Go installed, then run:
   ```
   go mod tidy
   ```

3. **Build the applications:**
   To build the agent and server, navigate to their respective directories and run:
   ```
   go build -o agent ./cmd/agent
   go build -o server ./cmd/server
   ```

## Usage
1. **Start the server:**
   ```
   ./server
   ```

2. **Run the agent:**
   In a separate terminal, run:
   ```
   ./agent
   ```

## Authentication Process
- The agent generates an authentication request using the methods defined in `internal/auth/agent.go`.
- The server processes this request and verifies it using the cryptographic circuit defined in `internal/auth/circuit.go`.
- Upon successful verification, the server responds to the agent, completing the authentication process.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.