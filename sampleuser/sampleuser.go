package sampleuser

import "github.com/fnami0316/moq_sample/sample"

type SampleUser struct {
	sample.Sample
}

// TimeOfDay 時間帯を指す文字列を返す
func (u *SampleUser)TimeOfDay() string {
	hour := u.Now().Hour()

	if 5 < hour && hour <= 12 {
		return "morning"
	} else if 12 < hour && hour <= 17 {
		return "afternoon"
	} else if 17 < hour && hour <= 19 {
		return "evening"
	} else {
		return "night"
	}
}