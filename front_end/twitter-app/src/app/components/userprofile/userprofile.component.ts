import { Component, OnInit, Input, Inject} from '@angular/core';
import { Tweet } from '../../models/tweet.model';
import { TweetlistComponent } from '../tweetlist/tweetlist.component';
import { MatDialog } from '@angular/material'
import { EditUserProfileDialogComponent } from '../edit-user-profile-dialog/edit-user-profile-dialog.component';
import { Subscription } from 'rxjs/Subscription';

@Component({
  selector: 'app-userprofile',
  templateUrl: './userprofile.component.html',
  styleUrls: ['./userprofile.component.css']
})
export class UserprofileComponent implements OnInit {

  userName: string;
  userFirstName: string;
  userLastName: string;
  userEmail: string;
  userBio: string;
  dialogResult = "";

  constructor(public dialog : MatDialog,
              ) { }

  ngOnInit() {
    let userinfo = JSON.parse(localStorage.getItem("user_info_object"));
    if (userinfo) {
        this.userName = userinfo.username;
        this.userFirstName = userinfo.firstname;
        this.userLastName = userinfo.lastname;
        this.userEmail = userinfo.email;
        this.userBio = userinfo.bio;
    }
  }
  
  /**
   * Get User Info
   */
  getUserInfo() {
    let userinfo = JSON.parse(localStorage.getItem("user_info_object"));
    if (userinfo) {
        this.userName = userinfo.username;
        this.userFirstName = userinfo.firstname;
        this.userLastName = userinfo.lastname;
        this.userEmail = userinfo.email;
        this.userBio = userinfo.bio;
    }
  }

  /*
  * Edit profile
  * Open a dilog EditUserProfileDialogComponent
  */
  openEditUserProfileDiag() {
    let dialogRef = this.dialog.open(EditUserProfileDialogComponent, {
      width: '600px',
      data: 'The dialog data shows here',

    })

  /**
   *  After the dialog is closed
   */
  dialogRef.afterClosed().subscribe(result => {
    // Fresh user info 
    console.log('Dialog is closed: ${result}');
    this.dialogResult = result;
    let userinfo = JSON.parse(localStorage.getItem("user_info_object"));
    
    userinfo = JSON.parse(localStorage.getItem("user_info_object"));
    if (userinfo) {
        this.userName = userinfo.username;
        this.userFirstName = userinfo.firstname;
        this.userLastName = userinfo.lastname;
        this.userEmail = userinfo.email;
        this.userBio = userinfo.bio;
    }
  })
  }

}
