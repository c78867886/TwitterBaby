import { Component, OnInit, Input } from '@angular/core';
import { Tweet } from '../../models/tweet.model';
import { TweetlistComponent } from '../tweetlist/tweetlist.component';

@Component({
  selector: 'app-userprofile',
  templateUrl: './userprofile.component.html',
  styleUrls: ['./userprofile.component.css']
})
export class UserprofileComponent implements OnInit {
  // @Input() userTweetList: Tweet[];
  // @Input() username: string;
  constructor() { }

  ngOnInit() {
  }

}
