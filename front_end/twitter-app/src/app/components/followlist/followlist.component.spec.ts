import { FollowlistComponent } from './followlist.component';
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { DataService } from '../../services/data.service';
import { HttpModule } from '@angular/http';
import { FormsModule } from '@angular/forms';
import { RouterTestingModule } from '@angular/router/testing';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { DebugElement }    from '@angular/core';
import { By }              from '@angular/platform-browser';
import { MatToolbarModule, 
  MatInputModule, 
  MatMenuModule, 
  MatIconModule, 
  MatButtonModule, 
  MatCardModule, 
  MatExpansionModule,
  MatProgressSpinnerModule,
  MatChipsModule,
  } from '@angular/material';

<<<<<<< HEAD
// describe('FollowlistComponent', () => {
//   let component: FollowlistComponent;
//   let fixture: ComponentFixture<FollowlistComponent>;

=======
describe('FollowlistComponent', () => {
  let component: FollowlistComponent;
  let fixture: ComponentFixture<FollowlistComponent>;
  let spy: jasmine.Spy;
  let dataService: DataService;
  let de: DebugElement;
  let el: HTMLElement;
>>>>>>> 378f0b9effe6b76e9dc756c123d4775ef8d86004
  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [
        HttpModule,
        MatToolbarModule,
        MatInputModule,
        MatMenuModule,
        MatIconModule,
        MatButtonModule,
        MatCardModule,
        MatExpansionModule,
        MatProgressSpinnerModule,
        MatChipsModule,
        FormsModule,
        RouterTestingModule,
        BrowserAnimationsModule
      ],
      declarations: [ FollowlistComponent ],

      providers: [ {provide: 'data', useClass: DataService} ],
    })
    .compileComponents();
  }));

<<<<<<< HEAD
//   beforeEach(() => {
//     fixture = TestBed.createComponent(FollowlistComponent);
//     component = fixture.componentInstance;
//     fixture.detectChanges();
//   });

//   it('should be created', () => {
//     expect(component).toBeTruthy();
//   });
// });
=======
  beforeEach(() => {
    fixture = TestBed.createComponent(FollowlistComponent);
    component = fixture.componentInstance;
    dataService = fixture.debugElement.injector.get('data');
    spy = spyOn(dataService, 'showFollowing')
          .and.returnValue(Promise.resolve([]));
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });

  it('should have one following in the list', () => {
    component.followingList = [{bio: 'test', username: 'test', id: 'test'}];
    fixture.detectChanges();
    de = fixture.debugElement.query(By.css('.folBio'));
    el = de.nativeElement;
    expect(el.textContent).toContain('test');
  });
});
>>>>>>> 378f0b9effe6b76e9dc756c123d4775ef8d86004
