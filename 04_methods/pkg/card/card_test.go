package card

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		cards []Card
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "no cards", args: args{cards: []Card{}}, want: 0},
		{name: "one card", args: args{cards: []Card{
			{Balance: 10},
		}}, want: 10},
		{name: "multiple cards", args: args{cards: []Card{
			{Balance: 10},
			{Balance: 100},
		}}, want: 110},
	}
	for _, tt := range tests {
		if got := Sum(tt.args.cards); got != tt.want {
			t.Errorf("Sum() = %v, want %v", got, tt.want)
		}
	}
}

func TestService_Sum(t *testing.T) {
	type fields struct {
		Cards []*Card
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "no cards", fields: fields{[]*Card{}}, want: 0},
		{name: "one card", fields: fields{[]*Card{
			{Balance: 10},
		}}, want: 10},
		{name: "multiple cards", fields: fields{[]*Card{
			{Balance: 10},
			{Balance: 100},
		}}, want: 110},
	}

	for _, tt := range tests {
		s := &Service{
			Cards: tt.fields.Cards,
		}
		if got := s.Sum(); got != tt.want {
			t.Errorf("Sum() = %v, want %v", got, tt.want)
		}
	}
}
