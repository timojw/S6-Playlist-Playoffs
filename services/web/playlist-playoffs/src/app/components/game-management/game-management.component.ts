import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, RouterModule } from '@angular/router';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { CommonModule } from '@angular/common';
import { DragDropModule } from '@angular/cdk/drag-drop';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

interface Song {
  id: number;
  title: string;
  artist: string;
  points: number;
}

interface Vote {
  game_id: number;
  song_id: number;
  points: number;
}

interface VoteResult {
  song_id: number;
  points: number;
}

interface VotesResponse {
  votes: VoteResult[];
}

@Component({
  selector: 'app-game-management',
  templateUrl: './game-management.component.html',
  styleUrls: ['./game-management.component.css'],
  standalone: true,
  imports: [CommonModule, HttpClientModule, RouterModule, DragDropModule]
})
export class GameManagementComponent implements OnInit {
  gameId: string | null;
  game: any;
  error: string | null = null;
  songs: Song[] = [
    { id: 1, title: 'Bohemian Rhapsody', artist: 'Queen', points: 0 },
    { id: 2, title: 'Stairway to Heaven', artist: 'Led Zeppelin', points: 0 },
    { id: 3, title: 'Hotel California', artist: 'Eagles', points: 0 },
    { id: 4, title: 'Imagine', artist: 'John Lennon', points: 0 },
    { id: 5, title: 'Hey Jude', artist: 'The Beatles', points: 0 },
    { id: 6, title: 'Like a Rolling Stone', artist: 'Bob Dylan', points: 0 },
    { id: 7, title: 'Smells Like Teen Spirit', artist: 'Nirvana', points: 0 },
    { id: 8, title: 'Sweet Child O\' Mine', artist: 'Guns N\' Roses', points: 0 },
    { id: 9, title: 'Billie Jean', artist: 'Michael Jackson', points: 0 },
    { id: 10, title: 'Purple Haze', artist: 'Jimi Hendrix', points: 0 },
  ];

  constructor(private route: ActivatedRoute, private http: HttpClient) { }

  ngOnInit(): void {
    this.gameId = this.route.snapshot.paramMap.get('id');
    if (this.gameId) {
      this.fetchGameDetails();
      this.loadVotes();
    } else {
      this.error = 'No game ID provided';
    }
  }

  fetchGameDetails(): void {
    this.http.get(`http://4.182.99.25:8080/api/game/${this.gameId}`).subscribe(
      (data: any) => {
        this.game = data;
        console.log('Game data:', this.game);
      },
      (error: any) => {
        console.error('Error fetching game details:', error);
        this.error = `Error fetching game details: ${error.message}`;
      }
    );
  }

  drop(event: CdkDragDrop<Song[]>): void {
    moveItemInArray(this.songs, event.previousIndex, event.currentIndex);
    this.updatePoints();
  }

  updatePoints(): void {
    this.songs.forEach((song, index) => {
      song.points = 10 - index;
    });
  }

  saveVotes(): void {
    if (!this.gameId) {
      console.error('Game ID is not available');
      return;
    }

    const votes = this.songs.map(song => ({
      game_id: parseInt(this.gameId!, 10),
      song_id: song.id,
      points: song.points
    }));

    const headers = { 'Content-Type': 'application/json' };

    this.http.post('http://4.182.99.25:8080/api/votes', { votes }, { headers }).subscribe(
      () => {
        console.log('Votes saved successfully');
      },
      (error: any) => {
        console.error('Error saving votes:', error);
      }
    );
  }

  loadVotes(): void {
    if (!this.gameId) {
      console.error('Game ID is not available');
      return;
    }

    this.http.get<VotesResponse>(`http://4.182.99.25:8080/api/votes/${this.gameId}`).subscribe(
      (response: any) => {
        console.log('Response from API:', response); // Log the entire response

        const votes = response.votes;
        if (Array.isArray(votes)) {
          votes.forEach((vote: VoteResult) => {
            const song = this.songs.find(s => s.id === vote.song_id);
            if (song) {
              song.points = vote.points;
            }
          });
          this.songs.sort((a, b) => b.points - a.points);
        } else {
          console.error('Error: Invalid response format', response);
        }
      },
      (error: any) => {
        console.error('Error loading votes:', error);
      }
    );
  }
}
