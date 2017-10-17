import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { PostareaComponent } from '../postarea/postarea.component';
import { LoginComponent } from './login.component';
import { UserInfoComponent } from '../user-info/user-info.component';
import { TweetlistComponent } from '../tweetlist/tweetlist.component';
import { PageSplitComponent } from '../page-split/page-split.component';
import { DataService } from '../../services/data.service';
import { HttpModule } from '@angular/http';
import { FormsModule } from '@angular/forms';
import { RouterTestingModule } from '@angular/router/testing';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
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

describe('LoginComponent', () => {
  let component: LoginComponent;
  let fixture: ComponentFixture<LoginComponent>;
  let spy: jasmine.Spy;
  let dataService: DataService;
  let userName: string;
  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ 
        LoginComponent, 
        UserInfoComponent, 
        TweetlistComponent, 
        PageSplitComponent,    
        PostareaComponent ],
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
      providers: [ {provide: 'data', useClass: DataService} ],
    })
    .compileComponents();
  }));


  beforeEach(() => {
    fixture = TestBed.createComponent(LoginComponent);
    component = fixture.componentInstance;
    component.userInfo = {
      followingcount: 1,
      followercount: 2,
      followed: false,
      userinfo: {
        bio: 'test',
        email: 'test1',
        firstname: "string;",
        id: "string;",
        lastname: "string;",
        username: "string;"
      }};
    // dataService = fixture.debugElement.injector.get(DataService);
    // spy = spyOn(dataService, 'getUserInfo')
    //       .and.returnValue(Promise.resolve('test'));
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
