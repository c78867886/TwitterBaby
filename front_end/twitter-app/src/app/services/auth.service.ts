import { Injectable } from '@angular/core';
import { Http, Response, Headers, RequestOptions } from '@angular/http';
import { tokenNotExpired } from 'angular2-jwt';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';
import { Observable } from 'rxjs/Observable';
import  'rxjs/add/operator/map';
import  'rxjs/add/operator/do';
import  'rxjs/add/operator/catch';
import 'rxjs/add/operator/toPromise';

import { Tweet } from '../models/tweet.model';

@Injectable()
export class AuthService{
    private authUrl: String = "http://127.0.0.1:1323/api/v1";
    private loggedIn: boolean = false;
    localhost = "http://localhost:1323";
    private tempname: string;
    private loginSource = new BehaviorSubject<boolean>(false);

    constructor(private http: Http){
        this.loggedIn = !!localStorage.getItem('auth_token');
    }

    /**
     * Check if the user is logged in
     */

     isLoggedIn(): Observable<boolean> {
        this.loginSource.next(tokenNotExpired('access_token'));
        return this.loginSource.asObservable();
     }

     isLoggedInSimple(): boolean {
        return tokenNotExpired('access_token');
     }

     /**
      * Log the user in
      */ 
    login(username: string, password: string): Observable<string> {
            let loginfo: object = {email:username, password:password};
            let headers: Headers = new Headers({'content-type': 'application/json'});
            let options: RequestOptions = new RequestOptions({ headers: headers });
            return this.http.post(`${this.authUrl}/login`, loginfo, options)
                .map(res => res.json())
                .do(res => {
                    if(res.token) {
                        localStorage.setItem('auth_token', res.token);
                        localStorage.setItem('access_token', res.token);
                        localStorage.setItem('id', res.id);
                        localStorage.setItem('user_info_object', JSON.stringify(res));
                        this.loginSource.next(tokenNotExpired('access_token'));
                    }
                })
                .catch(this.handleError);
            
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
        
    /**
    * SignUp service
    * @param userInputInfo 
    */
    signUp(userInputInfo): Observable<string> {
        let signUpUserInfo: object = {username:userInputInfo.username, password:userInputInfo.password, 
                               firstname: userInputInfo.firstname, lastname:userInputInfo.lastname,
                               email: userInputInfo.emailAddr};
        let headers: Headers = new Headers({'content-type': 'application/json'});
        let options: RequestOptions = new RequestOptions({ headers: headers });
        console.log(`${this.authUrl}/signup`, signUpUserInfo, options);
        return this.http.post(`${this.authUrl}/signup`,signUpUserInfo, options)
            .map(res => res.json())
            .do(res => {
                if(res.token) {
                    console.log(res);
                }
            })
            .catch(this.handleError);
      }
}