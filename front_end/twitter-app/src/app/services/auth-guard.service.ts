import { Injectable, Inject } from '@angular/core';
import { Router, CanActivate } from '@angular/router';

@Injectable()
export class AuthGuardService {

  constructor(@Inject('auth') private auth, private router: Router) { }

  canActivate() {
    if(this.auth.isLoggedInSimple()) {
      return true;
    } else {
      this.router.navigateByUrl('/login');
      return false;
    }
  }
}
