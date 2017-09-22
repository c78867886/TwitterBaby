import { RouterModule, Routes } from '@angular/router';
import { UserInfoComponent } from './components/user-info/user-info.component';
import { TweetlistComponent } from './components/tweetlist/tweetlist.component';
const appRoutes: Routes = [
    { path: 'test', component: UserInfoComponent },
    { path: 'test/list', component: TweetlistComponent},
    { path: '**', redirectTo: 'test' }
];
export const rooting = RouterModule.forRoot(appRoutes);
