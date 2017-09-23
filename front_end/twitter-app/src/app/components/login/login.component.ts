import { Component, OnInit } from '@angular/core';
import { Tweet } from '../../models/tweet.model';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  list: Tweet[] = [
    {
      content: "This is a test content",
      timestamp: "2012-12-01"
    },
    {
      content: "This is a test content",
      timestamp: "2012-12-01"
    },
    {
      content: "This is a test content",
      timestamp: "2012-12-01"
    }
  ]
  constructor() { }

  ngOnInit() {
  }

}
