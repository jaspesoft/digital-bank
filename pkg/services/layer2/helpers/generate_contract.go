package layer2helpers

import (
	"bytes"
	accountdomain "digital-bank/internal/account/domain"
	"fmt"
	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"os"
	"strings"
	"time"
)

const (
	ACCOUNT_AGREEMENT  ContractType = "ACCOUNT_AGREEMENT"
	FORTRESS_AGREEMENT ContractType = "FORTRESS_AGREEMENT"
)

type ContractType string

func mapperContractValue(a accountdomain.Account) []map[string]string {
	return []map[string]string{

		{
			"name":  "currentDate",
			"value": a.GetName(),
		},
		{
			"name":  "customerName",
			"value": a.GetName(),
		},
		{
			"name":  "customerOffices",
			"value": a.AccountHolder.GetAddress().City,
		},
		{
			"name":  "accountHolderName",
			"value": a.GetName(),
		},
		{
			"name":  "address",
			"value": a.AccountHolder.GetAddress().StreetOne + " " + a.AccountHolder.GetAddress().StreetTwo + " " + a.AccountHolder.GetAddress().City + " " + a.AccountHolder.GetAddress().Country,
		},
		{
			"name":  "by",
			"value": a.GetName(),
		},
		{
			"name":  "name",
			"value": a.GetName(),
		},
		{
			"name":  "title",
			"value": "Mr",
		},
		{
			"name":  "date",
			"value": time.Now().Format("2006-01-02"),
		},
	}
}

func generatePDF(htmlContent, fileName string) {
	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	pageReader := bytes.NewReader([]byte(htmlContent))

	// Añadir una nueva página con contenido HTML
	pdfg.AddPage(wkhtml.NewPageReader(pageReader))

	// Configura opciones del PDF
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtml.OrientationPortrait)
	pdfg.Grayscale.Set(false)

	// Genera el PDF
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Guardar el PDF en un archivo
	err = pdfg.WriteFile("./" + fileName + ".pdf")
	if err != nil {
		log.Fatal(err)
	}

}

func makeContract(a accountdomain.Account, contract ContractType, replacements []map[string]string) {

	// Read the HTML file
	htmlFilePath := "../templates_contract/account_agreement.html"
	if contract == FORTRESS_AGREEMENT {
		htmlFilePath = "../templates_contract/fortress_agreement.html"
	}

	htmlContent, err := os.ReadFile(htmlFilePath)
	if err != nil {
		log.Fatalf("Failed to read HTML file: %v", err)
	}
	// Convert HTML content to string
	htmlString := string(htmlContent)

	// Replace placeholders with actual values
	for _, replacement := range replacements {
		placeholder := fmt.Sprintf("{{%s}}", replacement["name"])
		htmlString = strings.ReplaceAll(htmlString, placeholder, replacement["value"])
	}

	generatePDF(htmlString, string(contract))
}

func GenerateContract(a accountdomain.Account) {
	replacements := mapperContractValue(a)
	makeContract(a, FORTRESS_AGREEMENT, replacements)
	makeContract(a, ACCOUNT_AGREEMENT, replacements)
}
