import { Component, OnInit, Inject } from '@angular/core';
import { UserprofileComponent } from '../userprofile/userprofile.component';
import { MatDialogRef } from '@angular/material';
import { MAT_DIALOG_DATA } from '@angular/material';


@Component({
  selector: 'app-edit-user-profile-dialog',
  templateUrl: './edit-user-profile-dialog.component.html',
  styleUrls: ['./edit-user-profile-dialog.component.css']
})
export class EditUserProfileDialogComponent implements OnInit {

  constructor(public thisDialogRef:MatDialogRef<EditUserProfileDialogComponent>,
              @Inject(MAT_DIALOG_DATA) public data:string ) { }

  ngOnInit() {
  }

  onCloseConfirm(){
      this.thisDialogRef.close('Confirm');
  }

  onCloseCancel(){
      this.thisDialogRef.close('Cancel');
  }
}
