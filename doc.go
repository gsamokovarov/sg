// Package sg provides simple transactional mails clients for SendGrid and
// SparkPost.
//
// The simplest way to send a mail is through `sg.Send` which uses a global client:
//
//	func main() {
//		// If you have your key in the SG_API_KEY environment variable, you may
//		// skip this step.
//		sg.Setup(sg.NewClient("API_KEY"))
//
//		sg.Send(&sg.Mail{
//			TemplateId:    "c2723c5e-b693-4086-968f-bd9057cc6ae4",
//			From:          "from@example.com",
//			To:            "to@example.com",
//			Substitutions: sg.H{"SUB": "value"},
//		})
//	}
//
// If you don't want to use a global client, you can build your own:
//
//	func main() {
//		// If you have your key in the SG_API_KEY environment variable, you may
//		// skip this step.
//		client := sg.NewClient("API_KEY")
//
//		client.Send(&sg.Mail{
//			TemplateId:    "c2723c5e-b693-4086-968f-bd9057cc6ae4",
//			From:          "from@example.com",
//			To:            "to@example.com",
//			Substitutions: sg.H{"SUB": "value"},
//		})
//	}
package sg
