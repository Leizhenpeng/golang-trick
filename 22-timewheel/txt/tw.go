package txt

import (
	"time"

	"github.com/rfyiamcool/go-timewheel"
)

type sms_tw struct {
	text   string
	status string
	timer  *timewheel.Task
}

/*
tw, err := timewheel.NewTimeWheel(1*time.Second, 20)
tw.Start()
defer tw.Stop()
*/
func NewSMS_(text string, tw *timewheel.TimeWheel) *sms_tw {
	s := &sms_tw{
		text:   text,
		status: "unverified",
	}
	task := tw.Add(3*time.Second, func() {
		if s.status == "unverified" {
			s = nil
		}
	})
	s.timer = task
	return s
}

func (s *sms_tw) Verify_(tw *timewheel.TimeWheel) {
	s.status = "verified"
	tw.Remove(s.timer)
}
