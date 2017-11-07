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
  @Input() commentlist;

  public commentsList: Comment[];
  subscriptComments: Subscription;
  showCommentOrNot:string;
  toggle: boolean;
  constructor( @Inject('data') private data,) { }

  ngOnInit() {
    this.toggle = false;
    this.showCommentOrNot = "Show More comments...";
    // this.getComments(this.tweet.id);
  }

  /**
   * Get Comments for the tweet
   */
  public getComments(tweetid){
    this.subscriptComments = this.data.fetchComment(tweetid)
    .subscribe(list => 
      {
        this.commentsList = list.commentlist;
      }
    );
  }
  
  showAllComments(){
    this.toggle = !this.toggle;
    if (this.toggle){
      this.showCommentOrNot = "Show less...";
    } else {
      this.showCommentOrNot = "Show More Comments...";
    }
  }

}
