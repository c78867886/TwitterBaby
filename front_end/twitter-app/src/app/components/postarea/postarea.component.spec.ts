import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PostareaComponent } from './postarea.component';

describe('PostareaComponent', () => {
  let component: PostareaComponent;
  let fixture: ComponentFixture<PostareaComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PostareaComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PostareaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
