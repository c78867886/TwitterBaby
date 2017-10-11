import { RouterModule, Routes, CanActivate } from '@angular/router';
import { UserInfoComponent } from './components/user-info/user-info.component';
import { LoginComponent } from './components/login/login.component';
import { TweetlistComponent } from './components/tweetlist/tweetlist.component';
import { UserPageComponent } from './components/user-page/user-page.component';
import { FollowerlistComponent } from './components/followerlist/followerlist.component';
import { FollowlistComponent } from './components/followlist/followlist.component';
import { UserloginComponent } from './components/userlogin/userlogin.component';
import { SignUpComponent } from './components/sign-up/sign-up.component';
import { AuthGuardService } from './services/auth-guard.service';
import { AuthGuardLoggedService } from './services/auth-guard-logged.service';

const appRoutes: Routes = [
    { path: 'home', component:  LoginComponent, canActivate: [AuthGuardService]},
    { path: 'user/:id', component: UserPageComponent, canActivate: [AuthGuardService]},
    { path: 'user/follower/:id', component: FollowerlistComponent, canActivate: [AuthGuardService]},
    { path: 'user/following/:id', component: FollowlistComponent, canActivate: [AuthGuardService]},
    { path: 'login', component: UserloginComponent, canActivate: [AuthGuardLoggedService]},
    { path: 'signup', component: SignUpComponent, canActivate: [AuthGuardLoggedService]},
    { path: '**', redirectTo: '/home' },
];

export const rooting = RouterModule.forRoot(appRoutes);
