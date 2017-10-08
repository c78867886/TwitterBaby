import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { AppComponent } from './app.component';
import { NavBarComponent } from './components/nav-bar/nav-bar.component';

import { DataService } from './services/data.service';

import { rooting } from './app.router';

import { MatToolbarModule, 
         MatInputModule, 
         MatMenuModule, 
         MatIconModule, 
         MatButtonModule, 
         MatCardModule, 
         MatExpansionModule,
         MatProgressSpinnerModule,
         MatChipsModule,
         } from '@angular/material';

import { UserInfoComponent } from './components/user-info/user-info.component';
import { TweetlistComponent } from './components/tweetlist/tweetlist.component';
import { PostareaComponent } from './components/postarea/postarea.component';
import { LoginComponent } from './components/login/login.component';
import { UserPageComponent } from './components/user-page/user-page.component';
import { FollowlistComponent } from './components/followlist/followlist.component';
import { FollowerlistComponent } from './components/followerlist/followerlist.component';
import { PageSplitComponent } from './components/page-split/page-split.component';
import { UserloginComponent } from './components/userlogin/userlogin.component';
import { AuthService } from './services/auth.service';

@NgModule({
  declarations: [
    AppComponent,
    NavBarComponent,
    UserInfoComponent,
    TweetlistComponent,
    PostareaComponent,
    LoginComponent,
    UserPageComponent,
    FollowlistComponent,
    FollowerlistComponent,
    PageSplitComponent,
    UserloginComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    HttpModule,
    FormsModule,
    MatToolbarModule,
    MatInputModule,
    MatMenuModule,
    MatIconModule,
    MatButtonModule,
    MatCardModule,
    MatExpansionModule,
    MatProgressSpinnerModule,
    MatChipsModule,
    rooting
  ],
  providers: [
    { provide: 'data', useClass: DataService},
    AuthService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
