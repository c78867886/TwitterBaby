import { Component, OnInit, Inject } from '@angular/core';

@Component({
  selector: 'app-nav-bar',
  templateUrl: './nav-bar.component.html',
  styleUrls: ['./nav-bar.component.css']
})
export class NavBarComponent implements OnInit {
  username: string = "";
  loginName: string = "";
  shouldBeShowed: boolean;
  constructor(@Inject('data') private data, @Inject('auth') private auth) { }

  ngOnInit() {
    if(this.auth.isLoggedIn()) {
      this.shouldBeShowed = true;
      let userinfo = JSON.parse(localStorage.getItem("user_info_object"));
      this.loginName = userinfo.firstname + ' ' + userinfo.lastname;
    } else {
      this.shouldBeShowed = false;
    }
  }

  onSubmit(): void {
    this.username = '';
    var input = document.getElementById("input");
    input.blur();
  }
}
