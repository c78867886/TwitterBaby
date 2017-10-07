import { Component, OnInit, Inject } from '@angular/core';
import { Tweet } from '../../models/tweet.model';
import { ActivatedRoute } from "@angular/router";
import { Subscription } from 'rxjs/Subscription';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  list: Tweet[];
  username: string = "";
  userInfo: object = null;
  subscriptionTweets: Subscription;
  constructor(
    private route: ActivatedRoute,
    @Inject('data') private data) { }

  ngOnInit() {
    this.data.mockLogin()
          .then(() => {
            let userinfo = JSON.parse(localStorage.getItem("user_info_object"));
            this.username = userinfo.username;
            return this.username;
          }).then((username) => {
            
            this.data.getUserInfo(username)
            .then(userinfo => 
              {
                this.userInfo = userinfo;
              }
            );

            this.subscriptionTweets = this.data.getTweetListTimeLine(username)
            .subscribe(list => 
              {
                this.list = list.tweetlist;
                //console.log(list);
              }
            );
          });
    
    
  }

}
