import { Component, OnInit, Input, Inject} from '@angular/core';
import { Tweet } from '../../models/tweet.model';
import { Comment } from '../../models/comment.model';
import { Subscription } from 'rxjs/Subscription';

@Component({
  selector: 'app-commentlist',
  templateUrl: './commentlist.component.html',
  styleUrls: ['./commentlist.component.css']
})
export class CommentlistComponent implements OnInit {
  @Input() tweet;

  commentsList: Comment[];
  subscriptComments: Subscription;
  
  constructor( @Inject('data') private data) { }

  ngOnInit() {
    console.log("Tweet-------------------------");
    console.log(this.tweet);
    this.subscriptComments = this.data.fetchComment(this.tweet.id)
    .subscribe(list => 
      {
        this.commentsList = list.commentlist;
        console.log(list);
      }
    );
  }

}
