package sample

import (
	"time"
)

type Sample interface {
	// Now 現在日時を返す
	Now() time.Time
}

type SampleImpl struct {}

var _ Sample = &SampleImpl{}

// NewSample Sampleのインスタンスを生成する
func NewSample() *SampleImpl {
	return &SampleImpl{}
}

// Now 現在日時を返す
func (s *SampleImpl) Now() time.Time {
	return time.Now()
}