"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const readline = __importStar(require("readline"));
const ws_1 = __importDefault(require("ws"));
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});
let tickets = [];
// Set up WebSocket connection to server
const ws = new ws_1.default('ws://localhost:8080/ws');
ws.on('open', () => {
    console.log('Client gestartet');
    // Prompt for client ID
    rl.question("Client-ID eingeben: ", clientId => {
        sendClientId(clientId);
        interact(clientId);
    });
});
// Receive data (list of all ticket information) from server
ws.on('message', message => {
    const data = JSON.parse(message.toString());
    tickets = data.tickets;
});
const sendClientId = (clientId) => {
    ws.send(JSON.stringify({
        message_type: "set_client_id",
        client_id: clientId
    }));
};
const interact = (clientId) => {
    console.log(`\nClient: ${clientId}`);
    displayTickets();
    rl.question("Nummer eingeben zur Selbstzuweisung, 'q' zum Beenden\n", input => {
        if (input === 'q') {
            ws.close();
            rl.close();
            process.exit();
        }
        const ticketId = parseInt(input);
        if (isNaN(ticketId)) {
            console.log("Bitte eine Nummer eingeben. Versuchen Sie nochmal.");
        }
        else if (ticketIsValid(ticketId)) {
            assignTicket(ticketId, clientId);
        }
        interact(clientId);
    });
};
const ticketIsValid = (ticketId) => {
    for (const ticket of tickets) {
        if (ticket.id === ticketId) {
            if (ticket.clientId === "") {
                return true;
            }
            else {
                console.log(`Ticket mit der ID: ${ticketId} ist schon zugewiesen.`);
                return false;
            }
        }
    }
    console.log(`Ticket mit der ID: ${ticketId} ist nicht gefunden.`);
    return false;
};
const assignTicket = (ticketId, clientId) => {
    ws.send(JSON.stringify({
        message_type: "assign_ticket",
        ticket_id: ticketId,
        client_id: clientId
    }));
};
const displayTickets = () => {
    if (tickets.length === 0) {
        console.log("Keine Tickets");
    }
    else {
        for (const ticket of tickets) {
            console.log(`Ticket ID: ${ticket.id}, zugewiesen von: ${ticket.clientId}`);
        }
    }
};
