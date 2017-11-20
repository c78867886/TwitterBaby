import { Component, OnInit, Inject } from '@angular/core';
import { MAT_DIALOG_DATA } from '@angular/material';
import { MatDialogRef } from '@angular/material';
@Component({
  selector: 'app-retweet-dialog',
  templateUrl: './retweet-dialog.component.html',
  styleUrls: ['./retweet-dialog.component.css']
})
export class RetweetDialogComponent implements OnInit {

  constructor(@Inject(MAT_DIALOG_DATA) private tweet,
              private thisDialogRef: MatDialogRef<RetweetDialogComponent>) { }

  ngOnInit() {
  }

  onSubmit() {
    this.thisDialogRef.close("Share!");
  }

}
