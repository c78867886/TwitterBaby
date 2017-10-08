import { RouterModule, Routes } from '@angular/router';
import { UserInfoComponent } from './components/user-info/user-info.component';
import { LoginComponent } from './components/login/login.component';
import { TweetlistComponent } from './components/tweetlist/tweetlist.component';
import { UserPageComponent } from './components/user-page/user-page.component';
import { FollowerlistComponent } from './components/followerlist/followerlist.component';
import { FollowlistComponent } from './components/followlist/followlist.component';
import { UserloginComponent } from './components/userlogin/userlogin.component';

const appRoutes: Routes = [
    { path: 'home', component:  LoginComponent},
    { path: 'user/:id', component: UserPageComponent},
    { path: 'user/follower/:id', component: FollowerlistComponent},
    { path: 'user/following/:id', component: FollowlistComponent},
    { path: '**', redirectTo: '/home' },
    { path:'login', component: UserloginComponent},
];

export const rooting = RouterModule.forRoot(appRoutes);
