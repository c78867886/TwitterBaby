import { Component, OnInit, Inject } from '@angular/core';
import { Tweet } from '../../models/tweet.model';
import { ActivatedRoute } from "@angular/router";
@Component({
  selector: 'app-user-page',
  templateUrl: './user-page.component.html',
  styleUrls: ['./user-page.component.css']
})
export class UserPageComponent implements OnInit {
  list: Tweet[];
  username: string;
  bio: string;
  constructor(private route: ActivatedRoute,
    @Inject('data') private data) { }

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.data.getTweetList(params["id"])
        .then(list => 
          {
            this.list = list.tweets;
            this.username = list.firstname + ' ' + list.lastname;
            this.bio = list.bio;
          }
        );
    });
  }

}
