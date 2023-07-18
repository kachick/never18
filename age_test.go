package never18

import (
	"testing"
	"time"
)

func mustParseDate(date string) time.Time {
	then, err := time.Parse(time.DateOnly, date)
	if err != nil {
		panic(err)

	}

	return then
}

func TestTruthAge(t *testing.T) {
	testCases := []struct {
		name       string
		birth      time.Time
		moment     time.Time
		limitYears int
		fail       bool
		nominally  Report
		truth      Report
	}{
		{
			name:       "ドラえもんが生まれた日まで未来の技術で生き延びたのび太は何歳になっているのか",
			birth:      mustParseDate("1962-08-07"),
			moment:     mustParseDate("2112-09-03"),
			limitYears: 12,
			nominally: Report{
				Years:  150,
				Months: 0,
				Days:   27,
			},
			truth: Report{
				Years:  12,
				Months: 1656,
				Days:   27,
			},
		},
		{
			name:       "このコードを書き出した日の、のび太の年齢",
			birth:      mustParseDate("1962-08-07"),
			moment:     mustParseDate("2023-07-17"),
			limitYears: 12,
			nominally: Report{
				Years:  60,
				Months: 11,
				Days:   10,
			},
			truth: Report{
				Years:  12,
				Months: 587,
				Days:   10,
			},
		},
		{
			name:       "「ぼくの生まれた日」 ~ てんとう虫コミックス第2巻",
			birth:      mustParseDate("1962-08-07"),
			moment:     mustParseDate("1962-08-07"),
			limitYears: 12,
			nominally: Report{
				Years:  0,
				Months: 0,
				Days:   0,
			},
			truth: Report{
				Years:  0,
				Months: 0,
				Days:   0,
			},
		},
		{
			name:       "「のび太が消えちゃう?」 ~ てんとう虫コミックス第43巻",
			birth:      mustParseDate("1962-08-07"),
			moment:     mustParseDate("1952-08-07"), // 20 years before the (almost) 10 years old.
			limitYears: 12,
			fail:       true,
		},
		{
			name:       "許容されている年齢までは、戸籍上と同様に振る舞う",
			birth:      mustParseDate("1962-08-07"),
			moment:     mustParseDate("1972-05-29"),
			limitYears: 17,
			nominally: Report{
				Years:  9,
				Months: 9,
				Days:   22,
			},
			truth: Report{
				Years:  9,
				Months: 9,
				Days:   22,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			age := Age{
				Birth: tc.birth,
			}

			nominally, err := age.Nominally(tc.moment)
			if err != nil {
				if tc.fail {
					return
				} else {
					t.Errorf("unexpected error happned: %v", err)
					return
				}
			}

			if tc.fail {
				t.Errorf("expected error did not happen")
				return
			}

			truth, err := age.Truth(tc.moment, tc.limitYears)
			if err != nil {
				// Errors should be failed in nominally calling even if tc.fail is true.
				t.Errorf("unexpected error happned: %v", err)
			}

			if nominally != tc.nominally {
				t.Errorf("got %v, wanted %v", nominally, tc.nominally)
			}

			if truth != tc.truth {
				t.Errorf("got %v, wanted %v", truth, tc.truth)
			}
		})
	}
}
