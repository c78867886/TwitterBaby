import { Component, OnInit, Inject } from '@angular/core';
import { ActivatedRoute } from "@angular/router";
import { Follower } from '../../models/follower.model';
@Component({
  selector: 'app-followerlist',
  templateUrl: './followerlist.component.html',
  styleUrls: ['./followerlist.component.css']
})
export class FollowerlistComponent implements OnInit {
  followerList: Follower[];
  constructor(
    private route: ActivatedRoute
    ,@Inject('data') private data) { }

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.data.showFollower(params['id'])
        .then((list) => {
          this.followerList = list;
        });
    });
  }

}
