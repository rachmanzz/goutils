package mailwrap

import (
	"fmt"

	mail "github.com/xhit/go-simple-mail/v2"
)

var mailServer *mail.SMTPServer

type MailWrap struct {
	email *mail.Email
}

// Set the global mail server
func NewMailServer(m *mail.SMTPServer) {
	mailServer = m
}

// MailClient returns a connected SMTP client
func MailClient() (*mail.SMTPClient, error) {
	if mailServer == nil {
		return nil, fmt.Errorf("mail server not initialized")
	}
	return mailServer.Connect()
}

// NewMail creates a new MailWrap (email instance)
func NewMail() *MailWrap {
	return &MailWrap{email: mail.NewMSG()}
}

// From sets the sender
func (m *MailWrap) From(email string, name ...string) *MailWrap {
	if len(name) > 0 {
		m.email.SetFrom(fmt.Sprintf("%s <%s>", name[0], email))
	} else {
		m.email.SetFrom(email)
	}
	return m
}

// To sets the recipient
func (m *MailWrap) To(email string, name ...string) *MailWrap {
	if len(name) > 0 {
		m.email.AddTo(fmt.Sprintf("%s <%s>", name[0], email))
	} else {
		m.email.AddTo(email)
	}
	return m
}

// CC sets CC recipients
func (m *MailWrap) CC(email string, name ...string) *MailWrap {
	if len(name) > 0 {
		m.email.AddCc(fmt.Sprintf("%s <%s>", name[0], email))
	} else {
		m.email.AddCc(email)
	}
	return m
}

// BCC sets BCC recipients
func (m *MailWrap) BCC(email string, name ...string) *MailWrap {
	if len(name) > 0 {
		m.email.AddBcc(fmt.Sprintf("%s <%s>", name[0], email))
	} else {
		m.email.AddBcc(email)
	}
	return m
}

// Subject sets the subject
func (m *MailWrap) Subject(subject string) *MailWrap {
	m.email.SetSubject(subject)
	return m
}

// TextContent sets plain text body (overwrite)
func (m *MailWrap) TextContent(body string) *MailWrap {
	m.email.SetBody(mail.TextPlain, body)
	return m
}

// HTMLContent sets HTML body (overwrite)
func (m *MailWrap) HTMLContent(body string) *MailWrap {
	m.email.SetBody(mail.TextHTML, body)
	return m
}

// TextAndHTML sets multi-part email (text + HTML)
func (m *MailWrap) TextAndHTML(text, html string) *MailWrap {
	m.email.SetBody(mail.TextPlain, text)
	m.email.AddAlternative(mail.TextHTML, html)
	return m
}

// Send sends the email via global mailServer or provided client
func (m *MailWrap) Send(client *mail.SMTPClient) error {
	if mailServer == nil {
		return fmt.Errorf("mail server not initialized")
	}

	if client == nil {
		var err error
		client, err = MailClient()
		if err != nil {
			return fmt.Errorf("failed to connect to SMTP server: %w", err)
		}
	}

	return m.email.Send(client)
}
