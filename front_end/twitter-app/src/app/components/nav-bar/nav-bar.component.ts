import { Component, OnInit, Inject } from '@angular/core';

@Component({
  selector: 'app-nav-bar',
  templateUrl: './nav-bar.component.html',
  styleUrls: ['./nav-bar.component.css']
})
export class NavBarComponent implements OnInit {
  username: string = "";
  loginName: string = "";
  constructor(@Inject('data') private data) { }

  ngOnInit() {
    let userinfo = JSON.parse(localStorage.getItem("user_info_object"));
    this.loginName = userinfo.firstname + ' ' + userinfo.lastname;
  }

  onSubmit(): void {
    this.username = '';
  }
}
