package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer_round(t *testing.T) {
	type args struct {
		num uint8
		plr Player
	}
	type wanted struct {
		points int16
	}
	tests := []struct {
		name   string
		args   args
		wanted wanted
	}{
		{
			name: "Exact match lower bound",
			args: args{
				num: 4,
				plr: Player{
					lowerBound: 4,
					upperBound: 6,
					points:     0,
				},
			},
			wanted: wanted{
				points: 5,
			},
		},
		{
			name: "Exact match upper bound",
			args: args{
				num: 6,
				plr: Player{
					lowerBound: 4,
					upperBound: 6,
					points:     0,
				},
			},
			wanted: wanted{
				points: 5,
			},
		},
		{
			name: "Inside bounds 1",
			args: args{
				num: 7,
				plr: Player{
					lowerBound: 3,
					upperBound: 8,
					points:     0,
				},
			},
			wanted: wanted{
				points: 0,
			},
		},
		{
			name: "Inside bounds 2",
			args: args{
				num: 7,
				plr: Player{
					lowerBound: 5,
					upperBound: 9,
					points:     0,
				},
			},
			wanted: wanted{
				points: 1,
			},
		},
		{
			name: "Inside bounds 3",
			args: args{
				num: 7,
				plr: Player{
					lowerBound: 2,
					upperBound: 9,
					points:     0,
				},
			},
			wanted: wanted{
				points: -2,
			},
		},
		{
			name: "Out bounds lower",
			args: args{
				num: 3,
				plr: Player{
					lowerBound: 4,
					upperBound: 6,
					points:     4,
				},
			},
			wanted: wanted{
				points: 3,
			},
		},
		{
			name: "Out bounds upper",
			args: args{
				num: 8,
				plr: Player{
					lowerBound: 4,
					upperBound: 6,
					points:     2,
				},
			},
			wanted: wanted{
				points: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.plr.round(tt.args.num)
			assert.Equal(t, tt.wanted.points, tt.args.plr.points)
		})
	}
}
