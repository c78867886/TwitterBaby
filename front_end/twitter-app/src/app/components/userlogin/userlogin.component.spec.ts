
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { NgModule } from '@angular/core';
import { UserloginComponent } from './userlogin.component';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthService } from '../../services/auth.service';
import { Observable } from 'rxjs/Observable';
import { Http, Response, Headers, RequestOptions } from '@angular/http';
import { FormsModule } from '@angular/forms';
import { rooting } from '../../app.router';
import  'rxjs/add/Observable/from';
import  'rxjs/add/Observable/empty';
import  'rxjs/add/Observable/throw';

describe('UserloginComponent', () => {
  let component: UserloginComponent;
  let fixture: ComponentFixture<UserloginComponent>;
  let authService: AuthService;
  let _route: Router;

  beforeEach(async(() => {

  }));

  beforeEach(() => {
    
    TestBed.configureTestingModule({
        imports: [ ReactiveFormsModule, FormsModule ],
        declarations: [ UserloginComponent ],
        providers: [AuthService ],
      })
      .compileComponents();
      // fixture = 
      authService = new AuthService(null);
      component = new UserloginComponent(authService, _route);
  });

  // it ('should be login successful if login() successfully', () => {
  //   let backAns = {"username": "a", "id": "1"}; 
  //   spyOn(authService, 'login').and.callFake(() => {
  //     return Observable.from([backAns]);
  //   })
     
  //   component.login();

  //   expect(component.successMessage).toContain("Login success");
  // })

  it ('should be signup failure if signup() is failed', () => {
    spyOn(authService, 'login').and.callFake(() => {
      return Observable.throw(Error);
    })
     
    component.login();

    expect(component.successMessage).toBeNull;
    expect(component.errorMessage).toContain("combination is not in our records");
  })
});

