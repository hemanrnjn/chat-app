import { Component, OnInit, ViewChild, ElementRef, ChangeDetectorRef, ChangeDetectionStrategy } from '@angular/core';
import { AuthService } from '../auth.service';

import * as moment from 'moment';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css'],
  changeDetection: ChangeDetectionStrategy.OnPush
})

export class HomeComponent implements OnInit {
  @ViewChild("messageBox") inputEl: ElementRef;

  socket: WebSocket
  allUsers = [];
  chatUsers = [];
  loggedInUser: any;
  selectedUser;
  allChats = [];
  currentChat = [];
  private host = "ws://" + window.location.host;

  constructor(private authService: AuthService, private ref: ChangeDetectorRef) { }

  ngOnInit() {

    this.loggedInUser = this.authService.getLoggedInUser();

    this.authService.getAllUserMessages(this.loggedInUser).subscribe((res: any) => {
      console.log(res);
      if (res.status) {
        this.allChats = res.messages;
      }
    });

    this.authService.getAllUsers().subscribe((res: any) => {
      console.log(res, this.loggedInUser)
      this.allUsers = res;
      const chatUsers = this.allUsers.filter(user => user.ID != this.loggedInUser.ID);
      chatUsers.forEach(user => {
        user.active = false;
      });
      this.chatUsers = chatUsers;
      console.log(this.chatUsers);
      this.ref.markForCheck();
    });

    this.socket = new WebSocket(this.host + '/ws');

    // on websocket error
    this.socket.addEventListener('error', (event) => {
      console.log(event);
      this.ref.markForCheck();
    });

    // Connection opened
    this.socket.addEventListener('open', (event) => {
      console.log("Connected!");
      const user: any = this.authService.getLoggedInUser();
      var msg = {
        from_user: user.ID,
        to_user: 0,
        username: user.username,
        message: "Connected!",
        is_read: false
      }
      console.log(msg);
      this.socket.send(JSON.stringify(msg));
      this.ref.markForCheck();
    });

    // Listen for messages
    this.socket.addEventListener('message', (event) => {
      var msg = JSON.parse(event.data);
      this.allChats.push(msg);
      this.currentChat = this.allChats.filter(chat => chat.to_user == this.selectedUser.ID || chat.from_user == this.selectedUser.ID);
      console.log(msg);
      this.ref.markForCheck();
    });
  }

  sendMessage(val) {
    const user: any = this.authService.getLoggedInUser();
    const msg: any = {
      timeStamp: moment().format(),
      from_user: user.ID,
      to_user: this.selectedUser.ID,
      username: user.username,
      message: val,
      is_read: false
    }
    this.socket.send(JSON.stringify(msg));
    this.allChats.push(msg);
    this.currentChat = this.allChats.filter(chat => chat.to_user == this.selectedUser.ID || chat.from_user == this.selectedUser.ID);
    this.ref.markForCheck();
  }

  selectUser(currUser) {
    this.chatUsers.filter(user => {
      if (user.ID != currUser.ID) {
        user.active = false
      }
      return;
    });
    this.currentChat = this.allChats.filter(chat => {
      return chat.to_user == currUser.ID || chat.from_user == currUser.ID
    });
    this.selectedUser = currUser;
    setTimeout(() => {
      this.inputEl.nativeElement.focus();
      this.ref.markForCheck();
    }, 0)
  }

  formatTimestamp(val) {
    return moment(val).format('LT') + ' | ' + moment(val).format('MMMM');
  }

}
