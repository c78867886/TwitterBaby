import { Component, OnInit, Inject } from '@angular/core';
import { ActivatedRoute } from "@angular/router";
import { Follower } from '../../models/follower.model';
@Component({
  selector: 'app-followlist',
  templateUrl: './followlist.component.html',
  styleUrls: ['./followlist.component.css']
})
export class FollowlistComponent implements OnInit {
  followingList: Follower[];
  constructor(private route: ActivatedRoute
    ,@Inject('data') private data) { }

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.data.showFollowing(params['id'])
        .then((list) => {
          this.followingList = list;
        });
    });
  }

}
