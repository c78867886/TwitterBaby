import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FollowerlistComponent } from './followerlist.component';

describe('FollowerlistComponent', () => {
  let component: FollowerlistComponent;
  let fixture: ComponentFixture<FollowerlistComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FollowerlistComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FollowerlistComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
