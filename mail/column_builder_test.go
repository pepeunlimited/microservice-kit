package mail

import (
	"testing"
)

func TestBuilder3_Build(t *testing.T) {

	createdAt := "27 June, 2017"
	firstName := "Piia"
	orderNumber := "442121"

	output,err := NewColumnBuilder().
		OneColumn(NewHtmlBuilder().
			H1("WE'VE GOT YOUR ORDER", "tc m0"). // Translate
			Span(createdAt, "tc db").
			Br().
			Br().
			Br().
			Br().
			Paragraph("Hi "+firstName+","). // Translate & OrderService
			Paragraph("Thank you for your order! We hope you enjoyed shopping with us."). // Translate
			Paragraph("While we get your order ready, please just double check the details below and let us know if anything needs changing."). // Translate
			Paragraph("We'll send you another email as soon as we ship your order."). // Translate
			Br().
			Hr().
			Br().
			Span("ORDER NO. "+orderNumber, "tc db").Br().
			Build()).
		//OneColumn(products.Build()).
		OneColumn(NewHtmlBuilder().Hr().Build()).
		OneColumn(NewHtmlBuilder().Span("Subtotal", "tl").Span("1", "tr db fr").Build()).
		OneColumn(NewHtmlBuilder().Span("Shipping cost", "tl").Span("0", " tr db fr").Build()).
		OneColumn(NewHtmlBuilder().Span("Taxes", "tl").Span("1", "tr db fr").Build()).
		OneColumn(NewHtmlBuilder().Span("Total", "tl").Span("1", "tr db fr").Build()).
		OneColumn(NewHtmlBuilder().
			Paragraph("Payment").
			Br().
			Br().
			Hr().
			Span("PayPal", "db").
			Span("1", "db").Br().
			Build()).
		OneColumn(NewHtmlBuilder().Span("BILLING ADDRESS", "tl").
			Span("DELIVERY ADDRESS", "tr db fr").Build()).
		OneColumn(NewHtmlBuilder().Span("Simo Al", "tl").
			Span("Simo Al", "tr db fr").Build()).
		OneColumn(NewHtmlBuilder().Span("Sörnäistenkatu 15 B 27", "tl").
			Span("Sörnäistenkatu 15 B 27", "tr db fr").Build()).
		OneColumn(NewHtmlBuilder().Span("00580, Helsinki", "tl").
			Span("00580, Helsinki", "tr db fr").Build()).
		OneColumn(NewHtmlBuilder().Span("fi", "tl").
			Span("fi", "tr db fr").Build()).
		OneColumn(NewHtmlBuilder().
			Br().
			Br().
			Build()).
		Footer("fi")

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	client := NewBuilder("", "").
		From(PiiiaCom, "PIIIA.com").
		To([]string{"medall@gmail.com"}).
		Subject("Subject!").
		Content(output).
		Build(Hotmail)
	err = client.Send()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
