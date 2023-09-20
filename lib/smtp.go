package lib

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"html/template"
	"net"
	"net/mail"
	"net/smtp"

	"go_grpc/lib/logger"
)

type SMTPClient struct {
	Host     string
	Port     string
	Password string
	Username string
	Sender   string
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown from server")
		}
	}

	return nil, nil
}

func (client *SMTPClient) Send(ctx context.Context, req SMTPRequest) error {
	address := fmt.Sprintf("%s:%s", client.Host, client.Port)
	payload := req.GetPayload(client)

	if client.Password != "" {
		host, _, _ := net.SplitHostPort(address)

		// TLS config
		tlsconfig := &tls.Config{
			ServerName: host,
		}

		c, err := smtp.Dial(address)
		if err != nil {
			logger.Error(ctx, "error smtp dialing", map[string]interface{}{
				"error": err,
				"tags":  []string{"smtp"},
			})

			return err
		}

		c.StartTLS(tlsconfig)

		// Auth
		if err = c.Auth(LoginAuth(client.Username, client.Password)); err != nil {
			logger.Error(ctx, "error smtp auth", map[string]interface{}{
				"error": err,
				"tags":  []string{"smtp"},
			})

			return err
		}

		// To && From
		if err = c.Mail(client.Sender); err != nil {
			logger.Error(ctx, "error issue MAIL command", map[string]interface{}{
				"error": err,
				"tags":  []string{"smtp"},
			})

			return err
		}

		if err = c.Rcpt(req.To); err != nil {
			logger.Error(ctx, "error issue RCPT command", map[string]interface{}{
				"error": err,
				"tags":  []string{"smtp"},
			})

			return err
		}

		// Data
		w, err := c.Data()
		if err != nil {
			logger.Error(ctx, "error issue DATA command", map[string]interface{}{
				"error": err,
				"tags":  []string{"smtp"},
			})

			return err
		}

		_, err = w.Write(payload)
		if err != nil {
			logger.Error(ctx, "error write payload", map[string]interface{}{
				"error": err,
				"tags":  []string{"smtp"},
			})

			return err
		}

		err = w.Close()
		if err != nil {
			logger.Error(ctx, "error closing", map[string]interface{}{
				"error": err,
				"tags":  []string{"smtp"},
			})

			return err
		}

		c.Quit()
	} else {
		err := smtp.SendMail(address, nil, client.Sender, []string{req.To}, payload)
		if err != nil {
			logger.Error(ctx, "error send smtp", map[string]interface{}{
				"error": err,
				"tags":  []string{"smtp"},
			})

			return err
		}
	}

	return nil
}

type SMTPRequest struct {
	To       string
	Subject  string
	Template string
	Data     interface{}
}

func (req SMTPRequest) GetPayload(client *SMTPClient) []byte {
	tmpl, err := template.ParseFiles(req.Template)
	if err != nil {
		return []byte{}
	}

	buf := new(bytes.Buffer)
	if err = tmpl.Execute(buf, req.Data); err != nil {
		return []byte{}
	}

	content := buf.String()

	from := mail.Address{"", client.Sender}
	to := mail.Address{"", req.To}
	subj := req.Subject

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj
	headers["MIME-version"] = "1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	return []byte(message + "\r\n" + content)
}
