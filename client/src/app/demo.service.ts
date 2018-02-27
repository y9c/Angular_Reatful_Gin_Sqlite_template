import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable()
export class DemoService {

  constructor(private http: HttpClient) {
  }

  // NOTE: all API calls in this file use simple endpoints served by
  // an Express app in the file app.js in the repo root. See that file
  // for all back-end code.
  // Uses http.get() to load data from a single API endpoint
  getMembers() {
    return this.http.get('/api/v1/member/');
  }

  // send a POST request to the API to create a new data object
  createMember(member) {
    let body = JSON.stringify(member);
    return this.http.post('/api/v1/member/', body, httpOptions);
  }

  // send a PUT request to the API to update a data object
  updateMember(member) {
    let body = JSON.stringify(member);
    return this.http.put('/api/v1/member/' + member.id, body, httpOptions);
  }

  // send a DELETE request to the API to delete a data object
  deleteMember(member) {
    return this.http.delete('/api/v1/member/' + member.id);
  }

}
