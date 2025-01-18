package main

import "time"

func switchTime(responseTime time.Duration) {
	switch {
	case responseTime < 100*time.Millisecond:
		less100ms++
	case responseTime >= 100*time.Millisecond && responseTime < 200*time.Millisecond:
		between100and200ms++
	case responseTime >= 200*time.Millisecond && responseTime < 300*time.Millisecond:
		between200and300ms++
	case responseTime >= 300*time.Millisecond && responseTime < 400*time.Millisecond:
		between300and400ms++
	case responseTime >= 400*time.Millisecond && responseTime < 500*time.Millisecond:
		between400and500ms++
	case responseTime >= 500*time.Millisecond && responseTime < 1*time.Second:
		between500and1000ms++
	case responseTime >= 1*time.Second && responseTime < 2*time.Second:
		between1000and2000ms++
	case responseTime >= 2*time.Second && responseTime < 3*time.Second:
		between2000and3000ms++
	case responseTime >= 3*time.Second && responseTime < 4*time.Second:
		between3000and4000ms++
	case responseTime >= 4*time.Second && responseTime < 5*time.Second:
		between4000and5000ms++
	case responseTime >= 5*time.Second && responseTime < 6*time.Second:
		between5000and6000ms++
	case responseTime >= 6*time.Second && responseTime < 7*time.Second:
		between6000and7000ms++
	case responseTime >= 7*time.Second && responseTime < 8*time.Second:
		between7000and8000ms++
	case responseTime >= 8*time.Second && responseTime < 9*time.Second:
		between8000and9000ms++
	case responseTime >= 9*time.Second && responseTime < 10*time.Second:
		between9000and10000ms++
	case responseTime >= 10*time.Second && responseTime < 11*time.Second:
		between10000and11000ms++
	case responseTime >= 11*time.Second && responseTime < 12*time.Second:
		between11000and12000ms++
	case responseTime >= 12*time.Second && responseTime < 13*time.Second:
		between12000and13000ms++
	case responseTime >= 13*time.Second && responseTime < 14*time.Second:
		between13000and14000ms++
	case responseTime >= 14*time.Second && responseTime < 15*time.Second:
		between14000and15000ms++
	case responseTime >= 15*time.Second && responseTime < 16*time.Second:
		between15000and16000ms++
	case responseTime >= 16*time.Second && responseTime < 17*time.Second:
		between16000and17000ms++
	case responseTime >= 17*time.Second && responseTime < 18*time.Second:
		between17000and18000ms++
	case responseTime >= 18*time.Second && responseTime < 19*time.Second:
		between18000and19000ms++
	case responseTime >= 19*time.Second && responseTime < 20*time.Second:
		between19000and20000ms++
	default:
		morethan20000ms++
	}
}
