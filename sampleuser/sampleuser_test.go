package sampleuser

import (
	"testing"
	"time"

	"github.com/fnami0316/moq_sample/sample"
	"github.com/fnami0316/moq_sample/sample/mock"
)

func TestSampleUser_TimeOfDay(t *testing.T) {
	type fields struct {
		Sample sample.Sample
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		wantCallCnt int
	}{
		{
			name: "12:00 ピッタリの場合、昼",
			fields: fields{
				Sample: &mock.SampleMock{
					NowFunc: func() time.Time {
						return time.Date(2021, 1, 1, 12, 0, 0, 0, time.UTC)
					},
				},
			},
			want: "afternoon",
		},
		{
			name: "17:00 ピッタリの場合、夕方",
			fields: fields{
				Sample: &mock.SampleMock{
					NowFunc: func() time.Time {
						return time.Date(2021, 1, 1, 17, 0, 0, 0, time.UTC)
					},
				},
			},
			want: "evening",
		},
		{
			name: "19:00 ピッタリの場合、夜",
			fields: fields{
				Sample: &mock.SampleMock{
					NowFunc: func() time.Time {
						return time.Date(2021, 1, 1, 19, 0, 0, 0, time.UTC)
					},
				},
			},
			want: "night",
		},
		{
			name: "5:00 ピッタリの場合、朝",
			fields: fields{
				Sample: &mock.SampleMock{
					NowFunc: func() time.Time {
						return time.Date(2021, 1, 1, 5, 0, 0, 0, time.UTC)
					},
				},
			},
			want: "morning",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &SampleUser{
				Sample: tt.fields.Sample,
			}
			if got := u.TimeOfDay(); got != tt.want {
				t.Errorf("SampleUser.TimeOfDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
