package never18

import (
	"testing"
	"time"
)

func getTime(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
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
			birth:      getTime(1962, 8, 7),
			moment:     getTime(2112, 9, 3),
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
			birth:      getTime(1962, 8, 7),
			moment:     getTime(2023, 7, 17),
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
			birth:      getTime(1962, 8, 7),
			moment:     getTime(1962, 8, 7),
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
			birth:      getTime(1962, 8, 7),
			moment:     getTime(1952, 8, 7), // 20 years before the (almost) 10 years old.
			limitYears: 12,
			fail:       true,
		},
		{
			name:       "許容されている年齢までは、戸籍上と同様に振る舞う",
			birth:      getTime(1962, 8, 7),
			moment:     getTime(1972, 5, 29),
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
		{
			name:       "#3 - same day in different month",
			birth:      getTime(2012, 9, 24),
			moment:     getTime(2023, 7, 24),
			limitYears: 17,
			nominally: Report{
				Years:  10,
				Months: 10,
				Days:   0,
			},
			truth: Report{
				Years:  10,
				Months: 10,
				Days:   0,
			},
		}, {
			name:       "#3 - before day in different month",
			birth:      getTime(2012, 9, 23),
			moment:     getTime(2023, 7, 24),
			limitYears: 17,
			nominally: Report{
				Years:  10,
				Months: 10,
				Days:   1,
			},
			truth: Report{
				Years:  10,
				Months: 10,
				Days:   1,
			},
		}, {
			name:       "#3 - next day in different month",
			birth:      getTime(2012, 9, 25),
			moment:     getTime(2023, 7, 24),
			limitYears: 17,
			nominally: Report{
				Years:  10,
				Months: 9,
				Days:   29,
			},
			truth: Report{
				Years:  10,
				Months: 9,
				Days:   29,
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
