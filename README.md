# ViChat

ViChat aims to be a lightweight video chat application with screen sharing and high quality audio built in Go. The project is inspired by tools like Discord but focuses on minimalism and ease of deployment.

## Architecture Overview

- **WebRTC** for peer-to-peer media transport.
- **WebSocket** for signalling between clients and the server.
- **Pion** Go library for WebRTC implementation on the server side.
- **STUN/TURN** servers for NAT traversal when peers cannot connect directly.

The server coordinates signalling and can act as a media relay if required. Clients (e.g. browser applications or native apps) establish WebRTC connections to exchange audio/video and screen sharing streams directly.

## Repository Structure

- `cmd/server` – entry point for the server application.
- `internal` – reusable packages (e.g. signalling logic).

## Getting Started

1. Install Go 1.20 or newer.
2. Run `go build ./cmd/server` to compile the signalling server.
3. Implement a WebRTC client (web or native) that communicates with the server via WebSocket.

This repository currently contains only a basic skeleton to get development started.
