import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { NgxEchartsService } from 'ngx-echarts';

@Component({
  selector: 'app-plot',
  templateUrl: './plot.component.html',
  styleUrls: ['./plot.component.css']
})

export class PlotComponent implements OnInit {
  // empty option before cellSeries loaded:
  options = {};

  constructor(private http: HttpClient, private es: NgxEchartsService) { }

  ngOnInit() {
    // fetch map geo JSON data from server
    this.http.get('/api/v1/cell/')
      .subscribe(cellSeries => {
        // update options:
        this.options = {
          title: {
            text: 'cell scatter plot',
          },
          xAxis: {},
          yAxis: {},
          series: cellSeries
        };
      });
  }

}
