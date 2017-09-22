import { Component, OnInit } from '@angular/core';
import { Tweet } from '../../models/tweet.model';
const list: Tweet[] = [
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

@Component({
  selector: 'app-tweetlist',
  templateUrl: './tweetlist.component.html',
  styleUrls: ['./tweetlist.component.css']
})
export class TweetlistComponent implements OnInit {
  tweetlist: Tweet[];
  constructor() { }

  ngOnInit() {
    this.tweetlist = list;
  }

}
