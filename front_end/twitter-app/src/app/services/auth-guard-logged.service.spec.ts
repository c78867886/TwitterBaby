import { TestBed, inject } from '@angular/core/testing';

import { AuthGuardLoggedService } from './auth-guard-logged.service';

describe('AuthGuardLoggedService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [AuthGuardLoggedService]
    });
  });

  it('should be created', inject([AuthGuardLoggedService], (service: AuthGuardLoggedService) => {
    expect(service).toBeTruthy();
  }));
});
