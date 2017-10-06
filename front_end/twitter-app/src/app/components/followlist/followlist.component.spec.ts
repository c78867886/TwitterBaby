import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FollowlistComponent } from './followlist.component';

describe('FollowlistComponent', () => {
  let component: FollowlistComponent;
  let fixture: ComponentFixture<FollowlistComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FollowlistComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FollowlistComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
