package mail

import (
	"fmt"
	"html/template"
)

type builder3 struct {
	htmls []template.HTML
}

func (builder *builder3) ParagraphWithClass(paragraph string, class string) Html {
	builder.htmls = append(builder.htmls, template.HTML(fmt.Sprintf("<p class=\"%v\">%v</p>",class, paragraph)))
	return builder
}

func (builder *builder3) Hr() Html {
	builder.htmls = append(builder.htmls, template.HTML("<hr/>"))
	return builder
}

func (builder *builder3) Br() Html {
	builder.htmls = append(builder.htmls, template.HTML("<br/>"))
	return builder
}

type Html interface {
	Paragraph(paragraph string) Html
	ParagraphWithClass(paragraph string, class string) Html
	H1(h1 string, class string) Html
	Image(src string, width int) Html
	Span(text string, class string) Html
	Br() Html
	Hr() Html
	Build() []template.HTML
}

func NewHtmlBuilder() Html {
	return &builder3{htmls:make([]template.HTML, 0)}
}

func (builder *builder3) Span(text string, class string) Html {
	builder.htmls = append(builder.htmls, template.HTML(fmt.Sprintf("<span class=\"%v\">%v</span>", class, text)))
	return builder
}

func (builder *builder3) Paragraph(paragraph string) Html {
	builder.htmls = append(builder.htmls, template.HTML(fmt.Sprintf("<p>%v</p>", paragraph)))
	return builder
}

func (builder *builder3) H1(h1 string, class string) Html {
	builder.htmls = append(builder.htmls, template.HTML(fmt.Sprintf("<h1 class=\"%v\">%v</h1>", class, h1)))
	return builder
}

func (builder *builder3) Image(src string, width int) Html {
	builder.htmls = append(builder.htmls, template.HTML(fmt.Sprintf("<img src=\"%v\" width=\"%v\" alt=\"\" />", src, width)))
	return builder
}

func (builder builder3) Build() []template.HTML {
	return builder.htmls
}