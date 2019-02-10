import { Component, OnInit } from '@angular/core';
import { AuthService } from './auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

  private isLoggedIn = false;

  constructor(private authService: AuthService, private router: Router) {
  }

  ngOnInit() {
    this.authService.getLoggedInStatus().subscribe(mess => {
      this.isLoggedIn = mess.loggedIn;
    });
  }

  logout() {
    this.authService.logout();
    this.router.navigate(['/auth']);
  }
}
