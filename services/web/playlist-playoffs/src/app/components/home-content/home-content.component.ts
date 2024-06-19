import { Component } from '@angular/core';
import { Router, RouterModule } from '@angular/router';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { faLink } from '@fortawesome/free-solid-svg-icons';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-home-content',
  templateUrl: './home-content.component.html',
  styleUrls: ['./home-content.component.css'],
  standalone: true,
  imports: [FontAwesomeModule, RouterModule, FormsModule]  // Import FormsModule here
})
export class HomeContentComponent {
  faLink = faLink;
  gameId: string;

  constructor(private router: Router) {}

  searchGame() {
    if (this.gameId) {
      this.router.navigate(['/game', this.gameId]);
    } else {
      alert('Please enter a valid game ID');
    }
  }
}
