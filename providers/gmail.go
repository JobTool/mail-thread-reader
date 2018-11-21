package providers

import (
  "net/http"
  "log"
  "strings"
  "encoding/base64"
  . "github.com/JobTool/mail-thread-reader/models"
  "google.golang.org/api/gmail/v1"
)

func FetchMailsFrom(oauthClient *http.Client, query string) []Mail {
  gmailApi, err := gmail.New(oauthClient)
  if err != nil {
    log.Fatalf("Unable to retrieve Gmail client: %v", err)
  }

  mails := []Mail {}
  user := "me"
  pageToken := ""
  for {
    request := gmailApi.Users.Messages.List(user).Q(query)
    if pageToken != "" {
      request.PageToken(pageToken)
    }
    response, err := request.Do()
    if err != nil {
      log.Fatalf("Unable to retrieve messages: %v", err)
    }
    for _, message := range response.Messages {
      messageBody, err := gmailApi.Users.Messages.Get(user, message.Id).Do()
      if err != nil {
        log.Fatalf("Unable to retrieve message %v: %v", message.Id, err)
      }

      // Fetch all data about the messages: From, To and Date
      mail := Mail {}
      for _, messageHeader := range messageBody.Payload.Headers {
        value := messageHeader.Value
        if messageHeader.Name == "From" {
          mail.From = value
        }

        if messageHeader.Name == "To" {
          mail.To = value
        }

        if messageHeader.Name == "Subject" {
          mail.Subject = value
        }
      }

      payload := messageBody.Payload
      mimeType := payload.MimeType
      if (strings.Contains(mimeType, "multipart")) {
        mail.Body = parseMultPartBody(payload.Parts)
      } else {
        mail.Body = append(mail.Body, MailBody {
          mimeType,
          payload.Body.Data,
        })
      }
      mail.SetDate(messageBody.InternalDate)
      mails = append(mails, mail)
    }

    if response.NextPageToken == "" {
      break
    }
    pageToken = response.NextPageToken
  }

  return mails
}

func parseMultPartBody(parts []*gmail.MessagePart) []MailBody {
  var body []MailBody
  for _, message := range parts {
    mimeType := message.MimeType
    bodyData := message.Body.Data

    // TODO: Add support to different types of body data
    if (mimeType == HtmlType || mimeType == TextType) {
      body = append(body, MailBody {
        Type: mimeType,
        Content: parseStringBody(bodyData),
      })
    }
  }

  return body
}

func parseStringBody(bodyData string) string {
  decodedContent, _ := base64.URLEncoding.DecodeString(bodyData)
  return string(decodedContent)
}


