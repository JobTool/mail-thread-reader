package models

import (
  "time"
)

type Mail struct {
  From string
  To string
  Body []string
  Subject string
  Date time.Time
}

func (m *Mail) SetDate(epochms int64) {
  m.Date = time.Unix(epochms / 1000, 0)
}

func (m *Mail) ParseBody() {
  // The body must to be parsed from base64URL
}

