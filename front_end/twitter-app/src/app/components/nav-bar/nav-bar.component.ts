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
  hostName: string = "";
  shouldBeShowed: boolean;
  subscription: Subscription;
  constructor(@Inject('data') private data, @Inject('auth') private auth, private route: Router) { }

  ngOnInit() {
    this.subscription = this.auth.isLoggedIn().subscribe( bol => {
      if(bol) {
        this.shouldBeShowed = true;
        let userinfo = JSON.parse(localStorage.getItem("user_info_object"));
        this.loginName = userinfo.firstname + ' ' + userinfo.lastname;
        this.hostName = userinfo.username;
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

  refresh(): void {
    if (this.route.url !== '/home') {
      this.route.navigateByUrl('/home');
    } else {
      this.data.getTweetListTimeLine(this.hostName, 1);
    }
    

  }

  /**
  * Go to user profile
  */
  goToUserProfile(): void{
    console.log("Navigate to the userprofile webpage");
    this.auth.isLoggedIn();
    this.route.navigateByUrl('/userprofile');
  }
}
