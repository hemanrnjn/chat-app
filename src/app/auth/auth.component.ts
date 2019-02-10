import { Component, OnInit } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.css']
})
export class AuthComponent implements OnInit {

  private username: string;
  private email: string;
  private password: string;

  constructor(private authService: AuthService, private router: Router) { }

  ngOnInit() {
  }

  login() {
    const obj = {
      email: this.email,
      password: this.password
    }
    console.log(obj);
    this.authService.login(obj).subscribe((res: any) => {
      console.log(res);
      if(res.status) {
        this.authService.setSession(res);
        this.router.navigate(['/home']);
      }
    })
  }

}
