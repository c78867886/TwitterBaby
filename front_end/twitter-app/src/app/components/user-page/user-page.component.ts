import { Component, OnInit, Inject } from '@angular/core';
import { Tweet } from '../../models/tweet.model';
import { ActivatedRoute } from "@angular/router";
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
  constructor(private route: ActivatedRoute,
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
