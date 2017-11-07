import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EditCommentsDialogComponent } from './edit-comments-dialog.component';

describe('EditCommentsDialogComponent', () => {
  let component: EditCommentsDialogComponent;
  let fixture: ComponentFixture<EditCommentsDialogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EditCommentsDialogComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EditCommentsDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
