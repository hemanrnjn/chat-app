import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  private socket: WebSocket

  constructor() { }

  ngOnInit() {
    this.socket = new WebSocket('ws://127.0.0.1:8000/ws');

    // on websocket error
    this.socket.addEventListener('error', (event) => {
      console.log(event);
    });

    // Connection opened
    this.socket.addEventListener('open', (event) => {
      console.log(event);
      var msg = { "type": "hello" };
      this.socket.send(JSON.stringify(msg));
    });

    // Listen for messages
    this.socket.addEventListener('message', (event) => {
      console.log(event);
      var msg = JSON.parse(event.data);
    });
  }

}
