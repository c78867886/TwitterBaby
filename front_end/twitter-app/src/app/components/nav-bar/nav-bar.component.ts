import { Component, OnInit, Inject } from '@angular/core';
import { Subscription } from 'rxjs/Subscription';
import { Router } from '@angular/router';
@Component({
  selector: 'app-nav-bar',
  templateUrl: './nav-bar.component.html',
  styleUrls: ['./nav-bar.component.css']
})
export class NavBarComponent implements OnInit {
  username: string = "";
  loginName: string = "";
  shouldBeShowed: boolean;
  subscription: Subscription;
  constructor(@Inject('data') private data, @Inject('auth') private auth, private route: Router) { }

  ngOnInit() {
    this.subscription = this.auth.isLoggedIn().subscribe( bol => {
      if(bol) {
        this.shouldBeShowed = true;
        let userinfo = JSON.parse(localStorage.getItem("user_info_object"));
        this.loginName = userinfo.firstname + ' ' + userinfo.lastname;
      } else {
        this.shouldBeShowed = false;
      }
    });
    
  }

  onSubmit(): void {
    this.username = '';
    var input = document.getElementById("input");
    input.blur();
  }

  logout(): void {
    localStorage.clear();
    this.auth.isLoggedIn();
    this.route.navigateByUrl('/login');
  }
}
