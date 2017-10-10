import { RouterModule, Routes } from '@angular/router';
import { UserInfoComponent } from './components/user-info/user-info.component';
import { LoginComponent } from './components/login/login.component';
import { TweetlistComponent } from './components/tweetlist/tweetlist.component';
import { UserPageComponent } from './components/user-page/user-page.component';
import { FollowerlistComponent } from './components/followerlist/followerlist.component';
import { FollowlistComponent } from './components/followlist/followlist.component';
import { UserloginComponent } from './components/userlogin/userlogin.component';
import { SignUpComponent } from './components/sign-up/sign-up.component';

const appRoutes: Routes = [
    { path: 'home', component:  LoginComponent},
    { path: 'user/:id', component: UserPageComponent},
    { path: 'user/follower/:id', component: FollowerlistComponent},
    { path: 'user/following/:id', component: FollowlistComponent},
    { path: 'login', component: UserloginComponent},
    { path: 'signup', component: SignUpComponent},
    { path: '**', redirectTo: '/login' },
];

export const rooting = RouterModule.forRoot(appRoutes);
