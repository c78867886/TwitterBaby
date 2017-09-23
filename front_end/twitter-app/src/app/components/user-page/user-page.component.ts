import { Component, OnInit } from '@angular/core';
import { Tweet } from '../../models/tweet.model';
@Component({
  selector: 'app-user-page',
  templateUrl: './user-page.component.html',
  styleUrls: ['./user-page.component.css']
})
export class UserPageComponent implements OnInit {
  list: Tweet[] = [
    {
      content: "This is a test contentThis is a test contentThis is a test content",
      timestamp: "2012-12-02"
    },
    {
      content: "This is a test contentThis is a test contentThis is a test content",
      timestamp: "2012-12-03"
    },
    {
      content: "This is a test contentThis is a test content",
      timestamp: "2012-12-04"
    }
  ]
  constructor() { }

  ngOnInit() {
  }

}
