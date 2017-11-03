# SG

A simple library for sending transactional mails through [SendGrid] or
[SparkPost].

## Usage

The simplest way to send a mail is through `sg.Send` which uses a global client:

```go
package main

import "github.com/gsamokovarov/sg"

func main() {
	// If you have your key in the SG_API_KEY environment variable, you may
	// skip this step.
	sg.Setup(sg.NewSendGridClient("API_KEY"))

	sg.Send(&sg.Mail{
		TemplateId:    "c2723c5e-b693-4086-968f-bd9057cc6ae4",
		From:          "from@example.com",
		To:            "to@example.com",
		Substitutions: sg.H{"SUB": "value"},
	})
}
```

The global client is using [SendGrid], by default. If you don't want to use a
global client, you can build your own.

```go
package main

import "github.com/gsamokovarov/sg"

func main() {
	// If you have your key in the SG_API_KEY environment variable, you may
	// skip this step.
	client := sg.NewSparkPostClient("API_KEY")

	client.Send(&sg.Mail{
		TemplateId:    "c2723c5e-b693-4086-968f-bd9057cc6ae4",
		From:          "from@example.com",
		To:            "to@example.com",
		Substitutions: sg.H{"SUB": "value"},
	})
}
```

[SendGrid]: https://sendgrid.com
[SparkPost]: https://sparkpost.com
