package txt

import "time"

type sms_raw struct {
	text   string
	status string
	timer  *time.Timer
}

func NewSMS(text string) *sms_raw {
	s := &sms_raw{
		text:   text,
		status: "unverified",
	}
	s.timer = time.NewTimer(3 * time.Second)
	go func() {
		<-s.timer.C
		if s.status == "unverified" {
			s = nil
		}
	}()
	return s
}

func (s *sms_raw) Verify() {
	s.status = "verified"
	if !s.timer.Stop() && s.timer != nil {
		<-s.timer.C
	}
}
