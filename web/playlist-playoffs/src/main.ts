import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from './app/app.config';
import { AppComponent } from './app/app.component';
import { provideAuth0 } from '@auth0/auth0-angular';

bootstrapApplication(AppComponent, {
  providers: [
    provideAuth0({
      domain: '{dev-n4quk35t.us.auth0.com}',
      clientId: '{n6qfkjpBJnkTPaVT5K9J7ijxFCWCZEGJ}',
      authorizationParams: {
        redirect_uri: window.location.origin
      }
    }),
  ]
});