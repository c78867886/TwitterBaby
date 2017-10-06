import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-page-split',
  templateUrl: './page-split.component.html',
  styleUrls: ['./page-split.component.css']
})
export class PageSplitComponent implements OnInit {
  index: number = 1;
  totalPage: number = 3;
  constructor() { }

  ngOnInit() {
  }
  
  nextPage(): void {
    this.index < this.totalPage ? this.index++ : this.index;
  }

  prePage(): void {
    this.index > 1 ? this.index-- : this.index;
  }

}
