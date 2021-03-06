package sg

type o map[string]interface{}

// H is a shortcut for map[string]string. In Go 1.9 this will become a type
// alias.
type H map[string]string

// Mail represents a SendGrid transactional mailer.
type Mail struct {
	From           string
	FromName       string
	To             string
	ToName         string
	TemplateID     string
	TemplateInline string
	Subject        string
	Substitutions  H
	Attachments    []H
}
