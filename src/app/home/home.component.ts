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

  private socket: WebSocket
  private allUsers = [];
  private chatUsers = [];
  private loggedInUser: any;
  private selectedUser;
  private allChats = [];
  private currentChat = [];

  constructor(private authService: AuthService, private ref: ChangeDetectorRef) { }

  ngOnInit() {

    this.loggedInUser = this.authService.getLoggedInUser();

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

    this.socket = new WebSocket('ws://127.0.0.1:8000/ws');

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
        from: user.ID,
        to: '',
        username: user.Username,
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
      this.currentChat = this.allChats.filter(chat => chat.to == this.selectedUser.ID || chat.from == this.selectedUser.ID);
      console.log(msg);
      this.ref.markForCheck();
    });
  }

  sendMessage(val) {
    const user: any = this.authService.getLoggedInUser();
    const msg: any = {
      timeStamp: moment().format(),
      from: user.ID,
      to: this.selectedUser.ID,
      username: user.Username,
      message: val,
      is_read: false
    }
    this.socket.send(JSON.stringify(msg));
    this.allChats.push(msg);
    this.currentChat = this.allChats.filter(chat => chat.to == this.selectedUser.ID || chat.from == this.selectedUser.ID);
    this.ref.markForCheck();
  }

  selectUser(currUser) {
    this.chatUsers.filter(user => {
      if (user.ID != currUser.ID) {
        user.active = false
      }
      return;
    });
    this.currentChat = this.allChats.filter(chat => chat.to == currUser.ID || chat.from == currUser.ID);
    this.selectedUser = currUser;
    setTimeout( () => {
      this.inputEl.nativeElement.focus();
      this.ref.markForCheck();
    }, 0)    
  }

  formatTimestamp(val) {
    return moment(val).format('LT') + ' | ' + moment(val).format('MMMM');
  }

}
