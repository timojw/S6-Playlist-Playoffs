import { Component, NgModule } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { AuthButtonComponent } from './auth-button-component/auth-button-component.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    RouterOutlet,
    AuthButtonComponent
  ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})

export class AppComponent {

  title = 'playlist-playoffs';
}
