import { RouterModule, Routes } from '@angular/router';
import { UserInfoComponent } from './components/user-info/user-info.component';
import { LoginComponent } from './components/login/login.component';
import { TweetlistComponent } from './components/tweetlist/tweetlist.component';
import { UserPageComponent } from './components/user-page/user-page.component';

//Diane:
import { UserloginComponent } from './components/userlogin/userlogin.component';

const appRoutes: Routes = [
    { path: 'home', component:  LoginComponent},
    { path: 'user/:id', component: UserPageComponent},
    //{ path: '**', redirectTo: 'home' },
    {
        path: "",
        redirectTo: "UserloginComponent",
        pathMatch: "full"
    },
    
    {
        path: "userlogin",
        component: UserloginComponent
    },
    {
        path: "**",
        component: UserloginComponent
    },
    {
        path: "login",
        component: UserloginComponent
    },


];
export const routing = RouterModule.forRoot(appRoutes);
