import { Injectable } from '@angular/core';
import { Tweet } from '../models/tweet.model';
import { Http, Response, Headers, RequestOptions } from '@angular/http';
import 'rxjs/add/operator/toPromise';
@Injectable()
export class DataService {

  constructor(private http: Http) { }
  getTweetList(id: string): Promise<Tweet[]> {
    let auth: Object = {Authorization: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDY1OTA3MTUsImlkIjoiNTljYjY5Y2NhNTRkNzU3YThlMzljOTc0In0.8A8hQQtbZeYBT3LDmOu_-OnrmRsfSby-KZw0eAMJ06s"};
    return this.http.get(`http://localhost:1323/api/v1/tweetlist/${id}`, auth)
                      .toPromise()
                      .then((res: Response) => res.json())
                      .catch(this.handleError);
  }

  getUserInfo(id: string): Promise<Object> {
    let headers: Headers = new Headers();
    headers.append('Access-Control-Expose-Headers', 'Authorization');
    headers.append('Authorization', 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDY2NTc1MjYsImlkIjoiNTljYjY5Y2NhNTRkNzU3YThlMzljOTc0In0.RJjUkREgw-zpxUjVxc9-gn2gb5nRBb2IA0jud_GfByw');
    let options: RequestOptions = new RequestOptions({ headers: headers });
    return this.http.get(`http://127.0.0.1:1323/api/v1/userInfo?username=${id}`, options)
                      .toPromise()
                      .then((res: Response) => res.json())
                      .catch(this.handleError);
  }

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
