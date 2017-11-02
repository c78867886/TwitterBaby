import { Component, OnInit, Input } from '@angular/core';
import { Tweet } from '../../models/tweet.model';
import { TweetlistComponent } from '../tweetlist/tweetlist.component';
import { MatDialog } from '@angular/material'
import { EditUserProfileDialogComponent } from '../edit-user-profile-dialog/edit-user-profile-dialog.component';

@Component({
  selector: 'app-userprofile',
  templateUrl: './userprofile.component.html',
  styleUrls: ['./userprofile.component.css']
})
export class UserprofileComponent implements OnInit {
  // @Input() userTweetList: Tweet[];
  // @Input() username: string;
  dialogResult = "";

  constructor(public dialog : MatDialog) { }

  ngOnInit() {
  }

  openEditUserProfileDiag() {
    let dialogRef = this.dialog.open(EditUserProfileDialogComponent, {
      width: '600px',
      data: 'The dialog data',

    })

    dialogRef.afterClosed().subscribe(result => {
      console.log('Dialog is closed: ${result}');
      this.dialogResult = result;
    })
  }

}
