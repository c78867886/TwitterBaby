import { Component, OnInit, Input } from '@angular/core';
import { Tweet } from '../../models/tweet.model';
@Component({
  selector: 'app-tweetlist',
  templateUrl: './tweetlist.component.html',
  styleUrls: ['./tweetlist.component.css']
})
export class TweetlistComponent implements OnInit {
  @Input() tweetlist: Tweet[];
  @Input() username: string;
  constructor() { }

  ngOnInit() {
  }

}
