import { Injectable } from '@angular/core';
import { Tweet } from '../models/tweet.model';
import { Http, Response, Headers, RequestOptions } from '@angular/http';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';
import { Observable } from 'rxjs/Rx';
import 'rxjs/add/operator/toPromise';
@Injectable()
export class DataService {
  localhost = "http://localhost:1323";
  private timelineSource = new BehaviorSubject<object>([]);
  constructor(private http: Http) { }

  // Create the header for http request
  getHeader(): RequestOptions {
    let access_token: string = localStorage.getItem("access_token");
    let headers: Headers = new Headers();
    headers.append('Authorization', 'Bearer ' + access_token);
    return new RequestOptions({ headers: headers });
  }


  // Get the tweetlist. This method will be changed to return Observable
  getTweetList(id: string): Promise<object> {
    let options: RequestOptions = this.getHeader();
    return this.http.get(this.localhost +`/api/v1/tweetlist/${id}?perpage=100&page=1`, options)
                      .toPromise()
                      .then((res: Response) => res.json())
                      .catch(this.handleError);
  }
  // Get the timeline of host.
  getTweetListTimeLine(id: string, page: number): Observable<object> {
    let options: RequestOptions = this.getHeader();
    this.timelineSource.next([]);
    this.http.get(this.localhost +`/api/v1/tweettimeline/${id}?perpage=15&page=${page}`, options)
                      .toPromise()
                      .then((res: Response) => this.timelineSource.next(res.json()))
                      .catch(this.handleError);
    return this.timelineSource.asObservable();
  }

  getUserInfo(id: string): Promise<Object> {
    let options: RequestOptions = this.getHeader();
    return this.http.get(this.localhost + `/api/v1/userInfo?username=${id}`, options)
                      .toPromise()
                      .then((res: Response) => res.json())
                      .catch(this.handleError);
  }
  
  // Follow and unfollow part
  followUser(mongoid: string): Promise<Object> {
    let post = {};
    let options: RequestOptions = this.getHeader();
    return this.http.post(`http://127.0.0.1:1323/api/v1/follow/${mongoid}`, post, options)
      .toPromise()
      .catch(this.handleError);
  }

  unfollowUser(mongoid: string): void {
    
  }

  showFollower(mongoid: string): Promise<object> {
    let options: RequestOptions = this.getHeader();
    return this.http.get(this.localhost + `/api/v1/showFollower/${mongoid}`, options)
                      .toPromise()
                      .then((res: Response) => res.json())
                      .catch(this.handleError);
  }

  showFollowing(mongoid: string): Promise<object> {
    let options: RequestOptions = this.getHeader();
    return this.http.get(this.localhost + `/api/v1/showFollowing/${mongoid}`, options)
                      .toPromise()
                      .then((res: Response) => res.json())
                      .catch(this.handleError);
  }

  //Create new tweet
  postTweet(mongoid: string, id: string, content: string): Promise<Object>{
    let options: RequestOptions = this.getHeader();
    let message: object = {message: content};
    //console.log(message);
    return this.http.post(`http://127.0.0.1:1323/api/v1/newTweet/${mongoid}`, message, options)
            .toPromise()
            .then((res: Response) => {
              this.getTweetListTimeLine(id, 1);
              return res.json();
            })
            .catch(this.handleError);

  }

  //MockLogin only for development
  mockLogin(): Promise<Object> {
    let loginfo: object = {email:"hojason117@gmail.com", password:"test1"};
    let headers: Headers = new Headers({ 'content-type': 'application/json'});
    let options: RequestOptions = new RequestOptions({ headers: headers });
    return this.http.post('http://127.0.0.1:1323/api/v1/login', loginfo, options)
      .toPromise()
      .then((res: Response) => {
        console.log(res.json());
        localStorage.setItem('access_token', res.json().token);
        localStorage.setItem('id', res.json().id);
        localStorage.setItem('user_info_object', JSON.stringify(res.json()));
      })
      .catch(this.handleError);
  }

  // ERROR handler
  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error);
    return Promise.reject(error.body || error);
  }
}
