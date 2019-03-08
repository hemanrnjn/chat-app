import { Component, OnInit } from '@angular/core';
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  private socket: WebSocket
  private allUsers = [];
  private chatUsers = [];
  private loggedInUser: any;

  constructor(private authService: AuthService) { }

  ngOnInit() {

    this.loggedInUser = this.authService.getLoggedInUser();

    this.authService.getAllUsers().subscribe((res: any) => {
      console.log(res, this.loggedInUser)
      this.allUsers = res;
      const chatUsers = this.allUsers.filter(user => user.ID != this.loggedInUser.ID);
      for (let user in chatUsers) {
        user.active = false;
      }
    });

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
        from: user.email,
        to: "abc@abc.com",
        username: user.Username,
        message: "Hello World"
      }
      console.log(msg);
      this.socket.send(JSON.stringify(msg));
    });

    // Listen for messages
    this.socket.addEventListener('message', (event) => {
      var msg = JSON.parse(event.data);
      console.log(msg);
    });
  }

  sendMessage(val) {
    const user: any = this.authService.getLoggedInUser();
    const msg: any = {
      from: user.email,
      to: "abc@abc.com",
      username: user.Username,
      message: val
    }
    this.socket.send(JSON.stringify(msg));
  }

  selectUser(user) {
    console.log(user);
  }

}
