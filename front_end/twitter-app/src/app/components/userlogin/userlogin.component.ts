import { Component, OnInit } from '@angular/core';
import { AuthService } from '../../services/auth.service';
import { rooting } from '../../app.router';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-userlogin',
  templateUrl: './userlogin.component.html',
  styleUrls: ['./userlogin.component.css']
})

export class UserloginComponent implements OnInit {
  credentials = { username:'', password:'' };
  successMessage: string = '';
  errorMessage: string = '';

  constructor(private service: AuthService,
              private _route: Router) { }


  ngOnInit() {
  }

  /**
   * Login a user
   */
  login(){
    let tag = 'loginpwd';
    this.service.login(this.credentials.username, this.credentials.password)
        .subscribe(
          //if success
          data => {
            console.log(data);
            console.log("Login success! Navigating to your home webpage");
            //this._route.navigateByUrl("https://www.google.com/" );
            this._route.navigate(['/home']);
          },
          //if error
          err => {
            this.errorMessage = 'That email address/password combination is not in our records. Forgot Your Password?  Click the "Forgot Your Password?"'
            console.log(err);
          }

        )
  }

  // resetPassword(){
  //   this.loginForm.setValue({
  //     passwordForm: '';
  //   })
  // }

}


