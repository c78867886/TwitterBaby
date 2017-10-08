import { Component, OnInit, Input, Inject } from '@angular/core';

@Component({
  selector: 'app-page-split',
  templateUrl: './page-split.component.html',
  styleUrls: ['./page-split.component.css']
})
export class PageSplitComponent implements OnInit {
  index: number = 1;
  @Input() totalPage: number;
  @Input() mongoid: string;
  constructor(@Inject('data') private data) { }

  ngOnInit() {
  }
  
  nextPage(): void {
    if (this.index != this.totalPage) {
      this.index < this.totalPage ? this.index++ : this.index;
      this.data.getTweetListTimeLine(this.mongoid, this.index);
      window.scrollTo(0,-10);
    }
  }

  prePage(): void {
    if (this.index != 1) {
      this.index > 1 ? this.index-- : this.index;
      this.data.getTweetListTimeLine(this.mongoid, this.index);
      window.scrollTo(0,-10);
    }
  }

}
