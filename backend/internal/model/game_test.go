package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame_DefineWinner(t *testing.T) {
	type args struct {
		game Game
	}
	type wanted struct {
		winnerName string
	}
	tests := []struct {
		name   string
		args   args
		wanted wanted
	}{
		{
			name: "Winner by points 1",
			args: args{
				game: Game{
					players: []Player{
						{
							Name:       "player1",
							lowerBound: 4,
							upperBound: 6,
							points:     14,
						},
						{
							Name:       "player2",
							lowerBound: 2,
							upperBound: 7,
							points:     13,
						},
						{
							Name:       "player3",
							lowerBound: 1,
							upperBound: 5,
							points:     6,
						},
					},
				},
			},
			wanted: wanted{
				winnerName: "player1",
			},
		},
		{
			name: "Winner by points 2",
			args: args{
				game: Game{
					players: []Player{
						{
							Name:       "player1",
							lowerBound: 4,
							upperBound: 6,
							points:     -6,
						},
						{
							Name:       "player2",
							lowerBound: 2,
							upperBound: 7,
							points:     -3,
						},
						{
							Name:       "player3",
							lowerBound: 1,
							upperBound: 5,
							points:     -8,
						},
					},
				},
			},
			wanted: wanted{
				winnerName: "player2",
			},
		},
		{
			name: "Higher upper bound",
			args: args{
				game: Game{
					players: []Player{
						{
							Name:       "player1",
							lowerBound: 4,
							upperBound: 6,
							points:     10,
						},
						{
							Name:       "player2",
							lowerBound: 2,
							upperBound: 7,
							points:     10,
						},
						{
							Name:       "player3",
							lowerBound: 1,
							upperBound: 5,
							points:     10,
						},
					},
				},
			},
			wanted: wanted{
				winnerName: "player2",
			},
		},
		{
			name: "Higher lower bound",
			args: args{
				game: Game{
					players: []Player{
						{
							Name:       "player1",
							lowerBound: 4,
							upperBound: 6,
							points:     10,
						},
						{
							Name:       "player2",
							lowerBound: 2,
							upperBound: 7,
							points:     10,
						},
						{
							Name:       "player3",
							lowerBound: 1,
							upperBound: 7,
							points:     10,
						},
					},
				},
			},
			wanted: wanted{
				winnerName: "player1",
			},
		},
		{
			name: "Alphabetically",
			args: args{
				game: Game{
					players: []Player{
						{
							Name:       "player1",
							lowerBound: 4,
							upperBound: 6,
							points:     10,
						},
						{
							Name:       "player2",
							lowerBound: 2,
							upperBound: 7,
							points:     10,
						},
						{
							Name:       "player3",
							lowerBound: 1,
							upperBound: 7,
							points:     10,
						},
						{
							Name:       "player4",
							lowerBound: 4,
							upperBound: 7,
							points:     10,
						},
					},
				},
			},
			wanted: wanted{
				winnerName: "player1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.game.SearchCandidates()
			tt.args.game.DefineWinner()
			assert.Equal(t, tt.wanted.winnerName, tt.args.game.GetWinner().Name)
		})
	}
}
