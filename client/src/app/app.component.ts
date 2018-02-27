import { Component } from '@angular/core';
import { DemoService } from './demo.service';
import { Observable } from 'rxjs/Rx';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})


export class AppComponent {

  title = 'test!';

  public members;
  public member_name;

  constructor(private _demoService: DemoService) { }

  ngOnInit() {
    this.getMembers();
  }

  getMembers() {
    this._demoService.getMembers().subscribe(
      // the first argument is a function which runs on success
      data => { this.members = data },
      // the second argument is a function which runs on error
      err => console.error(err),
      // the third argument is a function which runs on completion
      () => console.log('done loading members')
    );
  }

  createMember(name) {
    let member = { name: name };
    this._demoService.createMember(member).subscribe(
      data => {
        // refresh the list
        this.getMembers();
        return true;
      },
      error => {
        console.error("Error saving member!");
        return Observable.throw(error);
      }
    );
  }

  updateMember(member) {
    this._demoService.updateMember(member).subscribe(
      data => {
        // refresh the list
        this.getMembers();
        return true;
      },
      error => {
        console.error("Error saving member!");
        return Observable.throw(error);
      }
    );
  }

  deleteMember(member) {
    if (confirm("Are you sure you want to delete " + member.name + "?")) {
      this._demoService.deleteMember(member).subscribe(
        data => {
          // refresh the list
          this.getMembers();
          return true;
        },
        error => {
          console.error("Error deleting member!");
          return Observable.throw(error);
        }
      );
    }
  }
}

