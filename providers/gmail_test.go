package providers

import (
  "fmt"
  . "github.com/JobTool/mail-thread-reader/models"
  "github.com/apibillme/vcr"
  "net/http"
  "testing"
)

func TestGmailProvider(test *testing.T) {
  assertCorrectMessage := func(t *testing.T, got, want []Mail) {
    t.Helper()
  }

  test.Run("Must to fetch all emails from spotify", func(t *testing.T) {
    vcr.Start("gmail_provider_from_spotify", nil)
    client := &http.Client {}
    mails := FetchMailsFrom(client, "from:@spotify")
    defer vcr.Stop()

    fmt.Println(mails)
    want := []Mail {}
    assertCorrectMessage(t, mails, want)
  })

  test.Run("Must to fetch all emails to spotify", func(t *testing.T) {
    vcr.Start("gmail_provider_to_spotify", nil)
    client := &http.Client {}
    mails := FetchMailsFrom(client, "to:@spotify")
    defer vcr.Stop()
    want := []Mail {}
    assertCorrectMessage(t, mails, want)
  })
}
