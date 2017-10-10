import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DianeNavBarComponent } from './diane-nav-bar.component';

describe('DianeNavBarComponent', () => {
  let component: DianeNavBarComponent;
  let fixture: ComponentFixture<DianeNavBarComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DianeNavBarComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DianeNavBarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
