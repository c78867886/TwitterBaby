import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { FormGroup, FormControl, Validators} from '@angular/forms';
import { SignUpComponent } from './sign-up.component';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthService } from '../../services/auth.service';
import { Observable } from 'rxjs/Observable';
import { Http, Response, Headers, RequestOptions } from '@angular/http';
import  'rxjs/add/Observable/from';
import  'rxjs/add/Observable/empty';
import  'rxjs/add/Observable/throw';

describe('SignUpComponent', () => {
  let component: SignUpComponent;
  // let router: Router;
  let authService: AuthService;
  let fixture: ComponentFixture<SignUpComponent>;
  let auth;

  beforeEach(async(() => {
    // TestBed.configureTestingModule({
    //   imports: [ ReactiveFormsModule],
    //   declarations: [ SignUpComponent ],
    //   providers: [AuthService],
    // })
    // .compileComponents();
  }));

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [ ReactiveFormsModule],
      declarations: [ SignUpComponent ],
      providers: [AuthService],
    })
    .compileComponents();
    
    // auth = TestBed.get(AuthService);
    // fixture = TestBed.createComponent(SignUpComponent);
    authService = new AuthService(null);
    component = new SignUpComponent(authService, null);

  });

  it ('should be sign up in successful if signup() successfully', () => {
    let userInputInfo = "username";
    let backAns = {"username": "a", "userid": 1}; 
    spyOn(authService, 'signUp').and.callFake(() => {
      return Observable.from([backAns]);
    })
     
    component.signUp(userInputInfo);

    expect(component.successMessage).toContain("Thank you");
  })

  it ('should be signup failure if signup() is failed', () => {
    let userInputInfo = "username";
    let backAns = {"username": "a", "userid": 1}; 
    spyOn(authService, 'signUp').and.callFake(() => {
      return Observable.throw(Error);
    })
     
    component.signUp(userInputInfo);

    expect(component.successMessage).toBeNull;
    expect(component.errorMessage).toContain("Something is wrong");
  })

  
});

