import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClient } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { environment } from 'src/environments/environment';

@Component({
  selector: 'app-create-game-form',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './create-game-form.component.html',
  styleUrl: './create-game-form.component.css'
})
export class CreateGameFormComponent {
  gameName: string = '';
  gameDeadline: string; 
  minDate: string;
  constructor(private http: HttpClient) {
    this.initializeDate();
  }

  initializeDate() {
    const today = new Date();
    this.minDate = today.toISOString().slice(0, 10);
    this.setDeadline(today);
  }

  onDateChange(selectedDate: string) {
    const date = new Date(selectedDate);
    this.setDeadline(date);
  }

  setDeadline(date: Date) {
    date.setHours(23, 59, 0, 0);
    this.gameDeadline = date.toISOString().slice(0, 16);
  }

  startGame() {
    const postData = {
      name: this.gameName,
      deadline: this.gameDeadline
    };

    // Log the request data to the console
    console.log('Sending POST request to API:', `${environment.apiUrl}/api/game/add`);
    console.log('Request data:', postData);

    this.http.post(`http://4.182.99.127:8080/api/game/add`, postData).subscribe({
      next: (response) => console.log('Response from API:', response),
      error: (error) => console.error('There was an error!', error)
    });
  }
}
