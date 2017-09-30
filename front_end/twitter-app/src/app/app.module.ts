import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { AppComponent } from './app.component';
import { NavBarComponent } from './components/nav-bar/nav-bar.component';

import { DataService } from './services/data.service';

import { routing } from './app.router';

import { MdToolbarModule, 
         MdInputModule, 
         MdMenuModule, 
         MdIconModule, 
         MdButtonModule, 
         MdCardModule, 
         MdExpansionModule } from '@angular/material';

import { UserInfoComponent } from './components/user-info/user-info.component';
import { TweetlistComponent } from './components/tweetlist/tweetlist.component';
import { PostareaComponent } from './components/postarea/postarea.component';
import { LoginComponent } from './components/login/login.component';
import { UserPageComponent } from './components/user-page/user-page.component';
import { UserloginComponent } from './components/userlogin/userlogin.component';
@NgModule({
  declarations: [
    AppComponent,
    NavBarComponent,
    UserInfoComponent,
    TweetlistComponent,
    PostareaComponent,
    LoginComponent,
    UserPageComponent,
    UserloginComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    HttpModule,
    FormsModule,
    MdToolbarModule,
    MdInputModule,
    MdMenuModule,
    MdIconModule,
    MdButtonModule,
    MdCardModule,
    MdExpansionModule,
    routing
  ],
  providers: [
    { provide: 'data', useClass: DataService}],
  bootstrap: [AppComponent]
})
export class AppModule { }
