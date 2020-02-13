package server

import (
	"gogw/statistics"
	"time"
)

var WINLENGTH = time.Second

type SpeedMonitor struct {
	Upload *statistics.WindowStat
	Download *statistics.WindowStat
}

func NewSpeedMonitor() *SpeedMonitor {
	return & SpeedMonitor{
		Upload: statistics.NewWindowStat(WINLENGTH, statistics.Sum, 100),
		Download: statistics.NewWindowStat(WINLENGTH, statistics.Sum, 100),
	}
}

func (sm *SpeedMonitor) Add(uploadSize int64, downloadSize int64) {
	if uploadSize > 0 {
		sm.Upload.Add(uploadSize)
	}

	if downloadSize > 0 {
		sm.Download.Add(downloadSize)
	}
}

func (sm *SpeedMonitor) GetUploadSpeed() (*statistics.ValueWithTime, error) {
	return sm.Upload.GetLatest()
}

func (sm *SpeedMonitor) GetDownloadSpeed() (*statistics.ValueWithTime, error) {
	return sm.Download.GetLatest()
}




