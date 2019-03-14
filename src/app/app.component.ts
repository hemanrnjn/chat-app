import { Component, OnInit } from '@angular/core';
import { AuthService } from './auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

  isLoggedIn = false;
  username: string;

  constructor(private authService: AuthService, private router: Router) {
    if(localStorage.getItem('id_token')) {
      this.isLoggedIn = true;
      const user = JSON.parse(localStorage.getItem('loggedInUser'));
        this.username = user.username
    }  
  }

  ngOnInit() {
    this.authService.getLoggedInStatus().subscribe(mess => {
      this.isLoggedIn = mess.loggedIn;
      if(this.isLoggedIn) {
        const user = JSON.parse(localStorage.getItem('loggedInUser'));
        this.username = user.Username
      }
    });
  }

  logout() {
    this.authService.logout();
    this.router.navigate(['/auth']);
  }
}
