import * as readline from 'readline'
import WebSocket from 'ws';

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
})

interface Ticket {
  id: number;
  clientId: string;
}

let clientId: string;
let tickets: Ticket[];

// Set up WebSocket connection to server
const ws = new WebSocket('ws://localhost:8080/ws');

ws.on('open', () => {
  console.log('Client gestartet');

  // Prompt and set client ID
  rl.question("Client-ID eingeben: ", clientId => {
    setClientId(clientId);
  })

  // Receive and display data (list of all ticket information) from server
  ws.on('message', message => {
    const data = JSON.parse(message.toString());
    displayTickets(data);
  })

  interact();
})

const setClientId = (clientId: string) => {
  ws.send(JSON.stringify({
    message_type: "set_client_id",
    client_id: "clientId"
  }));
}

const interact = () => {
  rl.question("Nummer eingeben zur Selbstzuweisung, 'q' zum Beenden", input => {
    const ticketId = parseInt(input);

    if (typeof ticketId === "number") {
      assignTicket(ticketId);
      interact();
    } else if (input === 'q') {
      ws.close();
      rl.close();
      process.exit();
    } else {
      console.log("Ungültige Eingabe. Versuchen Sie nochmal.");
      interact();
    }
  })
}

const assignTicket = (ticketId: number) => {
  ws.send(JSON.stringify({
    message_type: "assign_ticket",
    ticket_id: ticketId,
  }))
}

const displayTickets = (tickets: Ticket[]) => {
  if (tickets.length === 0) {
    console.log("Keine Tickets");
  } else {
    for (const ticket of tickets) {
      console.log(`Ticket ID: ${ticket.id}, zugewiesen von: ${ticket.clientId}\n`)
    }
  }
}
