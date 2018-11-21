package models

import (
  "time"
)

const (
  HtmlType = "text/html"
  TextType = "text/plain"
  MultipartType = "multipart/alternative"
  PdfType = "application/pdf"
)

type MailBody struct {
  Type string
  Content string
}

type Mail struct {
  From string
  To string
  Body []MailBody
  Subject string
  Date time.Time
}

func (m *Mail) SetDate(epochms int64) {
  m.Date = time.Unix(epochms / 1000, 0)
}

func (m *Mail) ParseBody() {
  // The body must to be parsed from base64URL
}

