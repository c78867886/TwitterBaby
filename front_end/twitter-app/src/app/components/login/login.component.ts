import { Component, OnInit, Inject, Input } from '@angular/core';
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
  totalPage: number;

  constructor(
    private route: ActivatedRoute,
    @Inject('data') private data) { }

  ngOnInit() {
    // this.data.mockLogin()
    //       .then(() => {
    //         let userinfo = JSON.parse(localStorage.getItem("user_info_object"));
    //         this.username = userinfo.username;
    //         return this.username;
    //       }).then((username) => {
            
    //         this.data.getUserInfo(username)
    //         .then(userinfo => 
    //           {
    //             this.userInfo = userinfo;
    //           }
    //         );

    //         this.subscriptionTweets = this.data.getTweetListTimeLine(username, 1)
    //         .subscribe(list => 
    //           {
    //             this.list = list.tweetlist;
    //             this.totalPage = list.totalpage;
    //             //console.log(list);
    //           }
    //         );
    //       });
      let userinfo = JSON.parse(localStorage.getItem("user_info_object"));
      this.username = userinfo.username;
      this.data.getUserInfo(this.username)
      .then(userinfo => 
        {
          this.userInfo = userinfo;
        }
      );

      this.subscriptionTweets = this.data.getTweetListTimeLine(this.username, 1)
      .subscribe(list => 
        {
          this.list = list.tweetlist;
          this.totalPage = list.totalpage;
          //console.log(list);
        }
      );
    
  }

}
