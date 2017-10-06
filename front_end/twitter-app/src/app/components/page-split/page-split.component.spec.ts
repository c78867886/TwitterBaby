import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PageSplitComponent } from './page-split.component';

describe('PageSplitComponent', () => {
  let component: PageSplitComponent;
  let fixture: ComponentFixture<PageSplitComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PageSplitComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PageSplitComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
