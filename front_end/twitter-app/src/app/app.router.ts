import { RouterModule, Routes } from '@angular/router';
import { UserInfoComponent } from './components/user-info/user-info.component';
const appRoutes: Routes = [
    { path: 'test', component: UserInfoComponent },
    { path: '**', redirectTo: 'test' }
];
export const rooting = RouterModule.forRoot(appRoutes);