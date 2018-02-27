import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';  // replaces previous Http service
import { FormsModule } from '@angular/forms';
import { DemoService } from './demo.service';   // our custom service, see below

import { AppComponent } from './app.component';
import { PlotComponent } from './plot/plot.component';

import { NgxEchartsModule } from 'ngx-echarts';

@NgModule({
  imports: [BrowserModule, FormsModule, HttpClientModule, NgxEchartsModule],
  declarations: [AppComponent, PlotComponent],
  providers: [DemoService],
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
  bootstrap: [AppComponent]

})
export class AppModule { }
