package main

import "testing"

func TestWithdraw(t *testing.T) {
	tests := []struct {
		name            string
		balance, amount int
		want            int
		wantErr         bool
	}{
		{"Normal withdraw", 1000, 750, 250, false},
		{"Invalid amount", 1000, -20, 0, true},
		{"Insufficient funds", 500, 600, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Withdraw(tt.balance, tt.amount)

			if !tt.wantErr && err != nil {
				t.Errorf("ไม่ควร error แต่ดันเจอ: %v", err)
			}
			if tt.wantErr && err == nil {
				t.Errorf("ควรร้อง error แต่เงียบ")
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Withdraw(%d, %d) = %d;want %d", tt.balance, tt.amount, got, tt.want)
			}
		})
	}
}
