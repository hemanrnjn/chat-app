import { Component, OnInit, ChangeDetectorRef, ChangeDetectionStrategy } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';

declare var $: any;

@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.css'],
  changeDetection: ChangeDetectionStrategy.OnPush
})

export class AuthComponent implements OnInit {

  reg_firstname: string;
  reg_lastname: string;
  reg_username: string;
  reg_email: string;
  reg_password: string;
  reg_re_password: string;
  login_email: string;
  login_password: string;
  isRegister = false;
  warning = {
    title: '',
    text: ''
  }

  constructor(private authService: AuthService,
    private router: Router,
    private ref: ChangeDetectorRef) { }

  ngOnInit() {
  }

  login() {
    const obj = {
      email: this.login_email,
      password: this.login_password
    }
    console.log(obj);
    this.authService.login(obj).subscribe((res: any) => {
      console.log(res);
      if(res.status) {
        localStorage.setItem('loggedInUser', JSON.stringify(res.account));
        this.authService.setSession(res);
        this.router.navigate(['/home']);
        this.ref.markForCheck();
      } else {
        this.warning.title = "Warning";
        this.warning.text = "Invalid Username or password";
        $('.toast').toast('show')
      }
    });
  }

  register() {
    if (this.reg_password == this.reg_re_password) {
      const obj = {
        firstname: this.reg_firstname,
        lastname: this.reg_lastname,
        username: this.reg_username,
        email: this.reg_email,
        password: this.reg_password
      }
      console.log(obj);
      this.authService.register(obj).subscribe((res: any) => {
        console.log(res);
        if(res.status) {
          this.router.navigate(['/home']);
          this.ref.markForCheck();
        }
      });
    } else {
      this.warning.title = "Warning";
      this.warning.text = "Passwords do not match";
      $('.toast').toast('show')
    }
  }

}
