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
  userInfo: object = null;
  username: string;
  constructor(private route: ActivatedRoute,
  @Inject('data') private data) { }

  ngOnInit() {
    this.route.params.subscribe(params => {
      
      this.data.getUserInfo(params["id"])
      .then(userinfo => 
        {
          console.log(userinfo);
          this.userInfo = userinfo;
          this.username = userinfo.userinfo.username;
        }
      );

      this.data.getTweetList(params["id"])
        .then(list => 
          {
            console.log(list);
            this.list = list;
          }
       );

    });
  }

}
