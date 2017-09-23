import { Component, OnInit, Inject } from '@angular/core';

@Component({
  selector: 'app-nav-bar',
  templateUrl: './nav-bar.component.html',
  styleUrls: ['./nav-bar.component.css']
})
export class NavBarComponent implements OnInit {
  username: string = "";
  constructor(@Inject('data') private data) { }

  ngOnInit() {
  }

  onSubmit(): void {
    this.username = '';
  }
}
