# ticketSim
> A simulated ticketing system, along with a server and client(s).

## About the project
This project is written to mainly learn about concurrency through Goroutines, channels and WebSockets.

The server is written in Go. It can handle multiple clients and serves these functions:
- Serves a HTTP Server
- Creates new, unassigned tickets
- Broadcasts new updates to the client

![image](https://github.com/brythnl/ticketSim/assets/98424627/0511c6d2-caa5-4d83-a06e-9d52893b94c8)


The client is written in TypeScript. It runs in Node and serves these functions:
- Assigns itself into an existing ticket of choice
- Sends updates back to the server

![image](https://github.com/brythnl/ticketSim/assets/98424627/cc5154fc-4a20-48db-9fd9-87ebda0e330f)
