
import { Component, OnInit, Input, Inject, ElementRef} from '@angular/core';
import { MatDialog } from '@angular/material'
import { EditUserProfileDialogComponent } from '../edit-user-profile-dialog/edit-user-profile-dialog.component';
import { Subscription } from 'rxjs/Subscription';
// import {ChangeDetectorRef} from '@angular/core';
import { Router } from '@angular/router';
// import { MediaService } from '../../services/media.service';

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
  private elem: ElementRef;

  constructor(public dialog : MatDialog,
              // private ref:ChangeDetectorRef,
              private route: Router,
              @Inject('media') private mediaService,
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
  
  uploadImg2(input): void {
    // this.elem.nativeElement.querySelector()
    console.log("Upload imaging!");
    if (input.files && input.files[0]){
      var reader = new FileReader();
      const formData = new FormData();
      formData.append("image", input.files[0]);
      this.mediaService.upload(formData).subscribe(res =>{
        console.log("Upload successfully!");
      })
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
    if (userinfo) {
        this.userName = userinfo.username;
        this.userFirstName = userinfo.firstname;
        this.userLastName = userinfo.lastname;
        this.userEmail = userinfo.email;
        this.userBio = userinfo.bio;
    }
    this.route.navigateByUrl('/userprofile');
    console.log(userinfo);
  })
  }

  uploadImg(event){
    var image = event.target.files[0];

    var pattern = /image-*/;
    var reader = new FileReader();

    if (!image.type.match(pattern)){
      console.error('File is not an image!');
      alert("The file is not an image, Please select an image");
      return;
    }

    

    const formData = new FormData();
    formData.append("image", image);
    this.mediaService.uploadImg(image)
      .then(res =>{
        console.log("Upload successfully!");
    })
    
  }
}
