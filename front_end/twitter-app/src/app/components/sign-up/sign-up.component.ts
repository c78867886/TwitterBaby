import { Component, OnInit, Inject } from '@angular/core';
import { DomSanitizer} from '@angular/platform-browser';
import { MatIconRegistry} from '@angular/material';
import { FormGroup, FormControl, Validators} from '@angular/forms';
import { Observable } from 'rxjs/Observable';
import { Http, Response, Headers, RequestOptions } from '@angular/http';
import { rooting } from '../../app.router';
import { ActivatedRoute, Router } from '@angular/router';
import  'rxjs/add/operator/map';
import  'rxjs/add/operator/do';
import  'rxjs/add/operator/catch';
import 'rxjs/add/operator/toPromise';

@Component({
  selector: 'app-sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css']
})

export class SignUpComponent implements OnInit {

  successMessage: string = '';
  errorMessage: string = '';
  form;

  constructor(@Inject('auth') private service,
              private _route: Router) { }

  ngOnInit() {
    /**
     * Form control
     * handle with the username, password, first name's validations. 
     */
    this.form = new FormGroup({
      username: new FormControl('',Validators.compose([
        Validators.required,
        Validators.pattern('[\\w\\-\\s\\/]+$'),
      ])),
      firstname: new FormControl('',Validators.compose([
        Validators.required,
        Validators.pattern('[\\w\\-\\s\\/]+$'),
      ])),
      lastname: new FormControl(''),
      password: new FormControl('', Validators.compose([
        Validators.required,
        Validators.minLength(6),
        Validators.maxLength(16),
        ,
      ])),
      emailAddr: new FormControl('',Validators.required),
    });
  }

  /**
   * Sign Up
   */
  signUp(userInputInfo){
    this.service.signUp(userInputInfo)
        .subscribe(
        data => { //if success
          this.successMessage = "Thank you for your signup, You sign up successfully!! webpage will be nagivated to login webpage shortly!"
          console.log(this.successMessage);
          setTimeout((router: Router) => {
            this._route.navigate(['/login']);
          }, 3000);

          // this._route.navigate(['/login']);
        },
        
        err => { //if error
          this.errorMessage = "Something is wrong, please sign up again."
          console.log(err);
        }
      )
  }
  
  private handleError(err){
        let errMessage: string;
        
        if (err instanceof Response){
            let body = err.json() || '';
            let error = body.error || JSON.stringify(body);
            errMessage = `${err.status} - ${err.statusText || ''} ${error}`;

        } else {
            errMessage = err.message ? err.message : err.toString();
        }
        return Observable.throw(errMessage);
  }

  // constructor(iconRegistry: MatIconRegistry, sanitizer: DomSanitizer) { 
  //   iconRegistry.addSvgIcon(
  //     'account_box',
  //     sanitizer.bypassSecurityTrustResourceUrl('ic_account_box_black_24px.svg'));
  // }

}
