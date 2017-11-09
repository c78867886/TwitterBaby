import { Component, OnInit, Input } from '@angular/core';
import { Tweet } from '../../models/tweet.model';
import { EditCommentsDialogComponent } from '../edit-comments-dialog/edit-comments-dialog.component';
import { MatDialog } from '@angular/material'
import { Subscription } from 'rxjs/Subscription';
import { CommentlistComponent } from '../commentlist/commentlist.component';

@Component({
  selector: 'app-tweetlist',
  templateUrl: './tweetlist.component.html',
  styleUrls: ['./tweetlist.component.css']
})
export class TweetlistComponent implements OnInit {
  @Input() tweetlist: Tweet[];
  @Input() username: string;
  yesShow: boolean;
  url: string = 'http://s7d2.scene7.com/is/image/PetSmart/PB1201_STORY_CARO-Authority-HealthyOutside-DOG-20160818?$PB1201$';
  constructor(public dialog : MatDialog) { }

  ngOnInit() {
  }

  editCommentDialog(tweet){
    console.log("Open a dialog");
    let dialogRef = this.dialog.open(EditCommentsDialogComponent, {
      width: '600px',
      data: tweet,
    })

    // After close the dialog
    dialogRef.afterClosed().subscribe(result => {
      console.log('Dialog is closed: ${result}');
    })

  }

  // showComments(tweet){
  //   this.yesShow = !this.yesShow;
  //   CommentlistComponent.showTweetComments(tweet);
  //   console.log("show comments in webpage");
  // }
}
