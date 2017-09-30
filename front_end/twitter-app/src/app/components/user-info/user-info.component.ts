import { Component, OnInit, Input,Inject } from '@angular/core';
import { UserInfo } from '../../models/userinfo.model';
@Component({
  selector: 'app-user-info',
  templateUrl: './user-info.component.html',
  styleUrls: ['./user-info.component.css']
})
export class UserInfoComponent implements OnInit {
  @Input() isHost: boolean;
  @Input() userInfo: UserInfo;
  url: string = 'http://s7d2.scene7.com/is/image/PetSmart/PB1201_STORY_CARO-Authority-HealthyOutside-DOG-20160818?$PB1201$';
  constructor(@Inject('data') private data) { }
  ngOnInit() { }

  followFunc(): void {
    if(!this.userInfo.followed) {
      this.data.followUser(this.userInfo.userinfo.id);
      this.userInfo.followed = true;
      this.userInfo.followercount++;
    } else {
      this.data.unfollowUser(this.userInfo.userinfo.id);
      this.userInfo.followed = false;
      this.userInfo.followercount--;
    }
      
  }

}
