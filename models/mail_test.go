package models

import (
  "testing"
  "time"
)

func TestMail(test *testing.T) {
  assertCorrectMessage := func(t *testing.T, got, want time.Time) {
    t.Helper()
    if got != want {
      t.Errorf("got '%s' want '%s'", got, want)
    }
  }

  test.Run("The date must to be parsed from epoch ms to time", func(t *testing.T) {
    mail := Mail {}
    mail.SetDate(1351700038292)
    want := time.Unix(1351700038, 0)
    assertCorrectMessage(t, mail.Date, want)
  })
}
