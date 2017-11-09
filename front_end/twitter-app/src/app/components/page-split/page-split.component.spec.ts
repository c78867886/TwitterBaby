import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { DataService } from '../../services/data.service';
import { PageSplitComponent } from './page-split.component';
import { HttpModule } from '@angular/http';
describe('PageSplitComponent', () => {
  let component: PageSplitComponent;
  let fixture: ComponentFixture<PageSplitComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [
        HttpModule
      ],
      declarations: [ PageSplitComponent ],
      providers: [ {provide: 'data', useClass: DataService} ],
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
