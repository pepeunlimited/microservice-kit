package mail

import (
	"bytes"
	"fmt"
	"github.com/aymerick/douceur/inliner"
	"html/template"
)

const (
	raw = head + `{{range .Raw}}{{ . }}{{end}}` + footer
)


type Columns struct {
	Raw 				[]template.HTML
}

func (columns Columns) Footer(lang string) (string, error) {

	begin := `<tr>
			<td class="one-column">
				<table width="100%">
					<tr>
						<td class="inner contents">
							<span class="db">PIIIAshop</span>
							<span class="db">PIIIA.com • 2950462-4 • Finland • info@piiia.com</span>
							<br/>`
	end := 				`<br/>
						</td>
					</tr>
				</table>
			</td>
		</tr>`

	columns.Raw = append(columns.Raw,template.HTML(begin))

	if lang == "fi" {
		columns.Raw = append(columns.Raw,template.HTML("<i class=\"db\">Tämä on automaattinen viesti, ethän vastaa tähän viestiin.</i>"))
	} else {
		columns.Raw = append(columns.Raw,template.HTML("<i class=\"db\">This is an automatic message, please do not reply.</i>"))
	}

	columns.Raw = append(columns.Raw,template.HTML(end))
	return columns.BuildColumns()
}

func (columns Columns) FullWidthImage(src string) ColumnStep {
	begin := `<tr>
                <td class="full-width-image">`
	end := `	<br/><br/>	
				</td>
            </tr>`
	columns.Raw = append(columns.Raw,template.HTML(begin))
	columns.Raw = append(columns.Raw, template.HTML(fmt.Sprintf(`<img src="%v" width="600" alt="" />`, src)))
	columns.Raw = append(columns.Raw,template.HTML(end))
	return columns
}

func (columns Columns) OneColumn(raw []template.HTML) ColumnStep {
	begin := `<tr>
				<td class="one-column">
					<table width="100%">
						<tr>
							<td class="inner contents">`

	end := `        		</td>
						</tr>
					</table>
				</td>
			</tr>`
	columns.Raw = append(columns.Raw,template.HTML(begin))
	for _, v := range raw {
		columns.Raw = append(columns.Raw,v)
	}
	columns.Raw = append(columns.Raw,template.HTML(end))
	return columns
}

func (columns Columns) TwoColumn(left []template.HTML, right []template.HTML) ColumnStep {
	twoRowBegin := `<div class="column">
						<table width="100%">
							<tr>
								<td class="inner">
									<table class="contents">
										<tr>
											<td class="text">`

	twoRowEnd := `			     			</td>
										</tr>
									</table>
								</td>
							</tr>
						</table>
		 			</div>`
	begin := `<tr>
				<td class="two-column">`
	end := `
	   			</td>
			</tr>`
	columns.Raw = append(columns.Raw,template.HTML(begin))
	columns.Raw = append(columns.Raw, template.HTML(twoRowBegin))
	for _, l := range left {
		columns.Raw = append(columns.Raw, l)
	}
	columns.Raw = append(columns.Raw, template.HTML(twoRowEnd))
	columns.Raw = append(columns.Raw, template.HTML(twoRowBegin))
	for _, r := range right {
		columns.Raw = append(columns.Raw, r)
	}
	columns.Raw = append(columns.Raw, template.HTML(twoRowEnd))
	columns.Raw = append(columns.Raw,template.HTML(end))
	return columns
}

func (columns Columns) BuildColumns() (string, error) {
	parse, err := template.New("raw-html").Parse(raw)
	if err != nil {
		return "", err
	}
	output := new(bytes.Buffer)
	err = parse.Execute(output, columns)
	if err != nil {
		return "", err
	}
	inlined, err := columns.cssInliner(output.String())
	if err != nil {
		return "", err
	}
	return inlined, nil
}



func (Columns) cssInliner(src string) (string, error) {
	html, err := inliner.Inline(src)
	if err != nil {
		return "", err
	}
	return html, nil
}

type ColumnStep interface {
	OneColumn(raw []template.HTML) ColumnStep
	TwoColumn(left []template.HTML, right []template.HTML) ColumnStep
	FullWidthImage(src string) ColumnStep
	BuildColumns() (string, error)
	Footer(lang string) (string, error)
}


func NewColumnBuilder() ColumnStep {
	return Columns{
		Raw:make([]template.HTML, 0)}
}