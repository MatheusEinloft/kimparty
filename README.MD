# Kimparty
**Project made with the purpose of learning [Go](https://golang.org/) + WebSockets + [SolidJS](https://solidjs.com/).**

It allows you to watch streaming content synchronously across multiple browsers, along with a chat feature for communication.

**Note: This project is a work in progress and it is not ready for use yet.**

## Technologies Used
- **Server**: [Go](https://golang.org/) + [gws](https://github.com/lxzan/gws)
- **Client (Browser Extension)**: [SolidJS](https://solidjs.com/) + [CRXjs](https://crxjs.dev/vite-plugin)

## Roadmap
- [x] Synchronized stream. Play, pause and seek are synchronized across all browser.
    - [X] YouTube support.
    - [ ] HBO Max support (planned).
    - [ ] Netflix support (planned).
- [ ] In-webpage chat for communication between members.

## Requirements
- [Go >=1.22](https://golang.org/)
- [Node.js >=21.6](https://nodejs.org/)
- [pnpm](https://pnpm.io/) (or npm or yarn)

## Setup for Development
### Install the dependencies:
- For the server, navigate to the **server** directory and run `go mod tidy`.
- For the browser extension, navigate to the **browser-extension** directory and run `pnpm install`.

### Run the server:
- Navigate to the **server** directory and run `go run cmd/main.go`.

### Load the browser extension:
- Navigate to the **browser-extension** directory and run `pnpm run dev`.
- Load the extension in your browser by following the instructions in the [CRXjs documentation](https://crxjs.dev/vite-plugin/getting-started/solid/dev-basics).

## Building for Production
### Build the server:
- Navigate to the **server** directory and run `go build -o kimparty cmd/main.go`.
### Build the browser extension:
- Navigate to the **browser-extension** directory and run `pnpm run build`.

*Dockerfile will be added soon*

## Usage
*Coming soon*

## Server Endpoints
| METHOD   |                      ENDPOINT                            |                              DESCRIPTION                           |
|----------|:---------------------------------------------------------|:-------------------------------------------------------------------|
| POST     |`/party?url=SOME_URL&capacity=SOME_CAPACITY`              | Creates a new party with specified URL and capacity.               |
| GET      |`/party?find=SOME_PARTY_ID`                               | Retrieves details of a party by its ID.                            |
| WS       |`/ws/join?party_id=SOME_PARTY_ID&username=SOME_USERNAME`  | WebSocket endpoint to join a party with given ID using a username. |

## Contributing
Contributions are welcome! If you'd like to contribute to Kimparty, feel free to fork the repository and submit a pull request.
