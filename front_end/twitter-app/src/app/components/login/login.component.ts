import { Component, OnInit, Inject } from '@angular/core';
import { Tweet } from '../../models/tweet.model';
import { ActivatedRoute } from "@angular/router";
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
  constructor(
    private route: ActivatedRoute,
    @Inject('data') private data) { }

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.data.getTweetList(params["id"])
        .then(list => 
          {
            this.list = list;
          }
        );
    });
  }

}
