import { Component, OnInit } from '@angular/core';
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  private socket: WebSocket
  private loggedInUser: any;

  constructor(private authService: AuthService) { }

  ngOnInit() {
    this.socket = new WebSocket('ws://127.0.0.1:8000/ws');

    // on websocket error
    this.socket.addEventListener('error', (event) => {
      console.log(event);
    });

    // Connection opened
    this.socket.addEventListener('open', (event) => {
      console.log(event);
      const user: any = this.authService.getLoggedInUser();
      var msg = {
        email: user.email,
        username: user.Username,
        message: "Hello World"
      }
      this.socket.send(JSON.stringify(msg));
    });

    // Listen for messages
    this.socket.addEventListener('message', (event) => {
      var msg = JSON.parse(event.data);
      console.log(msg);
    });
  }

}
