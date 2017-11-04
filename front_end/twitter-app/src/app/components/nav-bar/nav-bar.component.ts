import { Component, OnInit, OnDestroy, Inject } from '@angular/core';
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
  rcvnewTweet: boolean = false;
  //connected: boolean = false;
  constructor(@Inject('data') private data, @Inject('auth') private auth, 
              @Inject('notify') private notify, private route: Router) { }

  ngOnInit() {
    
    this.notify.getEventListener().subscribe(event => {
      console.log(this.notify.readyState());
      if (event.type === "open") {
        console.log("WS Connected!");
      } else if (event.type === "close") {
        console.log("WS Disconnected!");
      } else if (event.type === 'message') {
        console.log(event.data);
        if (event.data === "New tweets.") {
          this.rcvnewTweet = true;
        }
      }
    });

    this.subscription = this.auth.isLoggedIn().subscribe( bol => {
      if (bol) {
        this.shouldBeShowed = true;
        let userinfo = JSON.parse(localStorage.getItem("user_info_object"));
        this.loginName = userinfo.firstname + ' ' + userinfo.lastname;
        this.hostName = userinfo.username;
        setTimeout(() => {
          if ( !this.notify.readyState() ) {
            this.notify.connect(userinfo.username);
          }
        }, 2000);
        
      } else {
        this.shouldBeShowed = false;
        
        if (this.notify.readyState()) {
          this.notify.close();
        }
        
      }
    });
  }
  
  ngOnDestroy() {
    if (this.notify.readyState()) {
      this.notify.close();
    }
    
  }

  onSubmit(): void {
    this.username = '';
    var input = document.getElementById("input");
    input.blur();
  }

  logout(): void {
    localStorage.clear();
    this.auth.isLoggedIn();
    this.notify.close();
    this.route.navigateByUrl('/login');
  }

  refresh(): void {
    this.rcvnewTweet = false;
    if (this.route.url !== '/home') {
      this.route.navigateByUrl('/home');
    } else {
      this.data.getTweetListTimeLine(this.hostName, 1);
    }
    
  }
}
