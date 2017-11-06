import { Component, OnInit, Input, Inject} from '@angular/core';
import { MatDialogRef } from '@angular/material';
import { MAT_DIALOG_DATA } from '@angular/material';
import { Subscription } from 'rxjs/Subscription';
import { Tweet } from '../../models/tweet.model';

@Component({
  selector: 'app-edit-comments-dialog',
  templateUrl: './edit-comments-dialog.component.html',
  styleUrls: ['./edit-comments-dialog.component.css']
})
export class EditCommentsDialogComponent implements OnInit {

  public commentContent: string;
  constructor(public thisDialogRef:MatDialogRef<EditCommentsDialogComponent>,
              @Inject(MAT_DIALOG_DATA) public tweet,
              @Inject('data') private data,) { }

  ngOnInit() {
  }

  onCloseSubmit(){
    console.log("------------------");
    console.log(this.tweet);
    this.data.addNewComment(this.commentContent, this.tweet.id)
      .then(
        console.log("Adding comment is done")
      )
    this.thisDialogRef.close('Confirm');
  }

  /**
   * Cancel updating user profile
   */
  onCloseCancel(){
    this.thisDialogRef.close('Cancel');
  }

}
