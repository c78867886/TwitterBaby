import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { AppComponent } from './app.component';
import { NavBarComponent } from './components/nav-bar/nav-bar.component';

import { DataService } from './services/data.service';

import { rooting } from './app.router';

import { MdToolbarModule, MdInputModule } from '@angular/material';
import { UserInfoComponent } from './components/user-info/user-info.component';
@NgModule({
  declarations: [
    AppComponent,
    NavBarComponent,
    UserInfoComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MdToolbarModule,
    MdInputModule,
    rooting
  ],
  providers: [{provide: "data",
              useClass: DataService}],
  bootstrap: [AppComponent]
})
export class AppModule { }
