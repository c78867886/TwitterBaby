import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { TweetlistComponent } from './tweetlist.component';

describe('TweetlistComponent', () => {
  let component: TweetlistComponent;
  let fixture: ComponentFixture<TweetlistComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ TweetlistComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TweetlistComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
