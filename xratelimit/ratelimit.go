package xratelimit

import (
	"io"
	"time"
)

type RateLimitReader struct {
	RatePreSecs int64
	Reader      io.Reader
}

type RateLimitWriter struct {
	RatePreSecs int64
	Writer      io.Writer
}

func (r *RateLimitReader) Read(p []byte) (n int, err error) {
	startTime := time.Now().Nanosecond()
	n, err = r.Reader.Read(p)
	secs := int64(n) * int64(time.Second) / r.RatePreSecs
	elapsed := time.Now().Nanosecond() - startTime
	remaining := secs - int64(elapsed)
	if remaining > 0 {
		time.Sleep(time.Duration(remaining))
	}
	return
}

func (w *RateLimitWriter) Write(p []byte) (n int, err error) {
	startTime := time.Now().Nanosecond()
	n, err = w.Writer.Write(p)
	secs := int64(n) * int64(time.Second) / w.RatePreSecs
	elapsed := time.Now().Nanosecond() - startTime
	remaining := secs - int64(elapsed)
	if remaining > 0 {
		time.Sleep(time.Duration(remaining))
	}
	return
}
