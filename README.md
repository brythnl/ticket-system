# ticketSim
> A simulated ticketing system, along with a server and client(s).

## About the project
This project is written to mainly learn about concurrency through Goroutines, channels and WebSockets.

The server is written in Go. It can handle multiple clients and serves these functions:
- Serves a HTTP Server
- Creates new, unassigned tickets
- Broadcasts new updates to the client

The client is written in TypeScript. It runs in Node and serves these functions:
- Assigns itself into an existing ticket of choice
- Sends updates back to the server

## Simulation
To test/simulate this yourself, ensure you have Node and Go installed.
- To run the server:
```go
go run server.go
```
- To start a client (this example uses npm, adjust accordingly):
```ts
npx ts-node client.ts
```
To start another client, run the same command in another window.
