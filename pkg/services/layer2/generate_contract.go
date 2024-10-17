package layer2

import (
	"bytes"
	accountdomain "digital-bank/internal/account/domain"
	"digital-bank/pkg"
	"fmt"
	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"os"
	"strings"
	"time"
)

const (
	CONTRACT_ACCOUNT_AGREEMENT  ContractType = "ACCOUNT_AGREEMENT"
	CONTRACT_FORTRESS_AGREEMENT ContractType = "FORTRESS_AGREEMENT"
)

type ContractType string

func mapperContractValue(a *accountdomain.Account) []map[string]string {
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

func generatePDF(htmlContent string) ([]byte, error) {
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
		return nil, err
	}

	// Convert the PDF to a byte slice
	return pdfg.Bytes(), nil

}

func makeContract(contract ContractType, replacements []map[string]string) ([]byte, error) {

	// Read the HTML file
	htmlFilePath := "../templates_contract/account_agreement.html"
	if contract == CONTRACT_FORTRESS_AGREEMENT {
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

	return generatePDF(htmlString)
}

func generateContract(a *accountdomain.Account) {
	replacements := mapperContractValue(a)
	pdfByte, err := makeContract(CONTRACT_FORTRESS_AGREEMENT, replacements)
	if err != nil {
		log.Println("Error generating the contract: "+string(CONTRACT_FORTRESS_AGREEMENT), err)
		return
	}

	s3 := pkg.NewAWSS3Storage(os.Getenv("AWS_S3_BUCKET_DOCUMENT"))
	fileName, err := s3.UploadFile(pdfByte, string(CONTRACT_FORTRESS_AGREEMENT))

	if err != nil {
		log.Println("Error upload document " + string(CONTRACT_FORTRESS_AGREEMENT) + " " + string(accountdomain.FRONT))
	} else {

		a.GetAccountHolder().SetDocument(
			accountdomain.Document{
				AccountID:    a.GetAccountID(),
				Patch:        fileName,
				DocumentType: accountdomain.ACCOUNT_AGREEMENT,
				DocumentSide: accountdomain.FRONT,
			}, a.GetAccountHolder().GetIDNumber())
	}

	pdfByte, err = makeContract(CONTRACT_ACCOUNT_AGREEMENT, replacements)
	if err != nil {
		log.Println("Error generating the contract: "+string(CONTRACT_ACCOUNT_AGREEMENT)+" "+string(accountdomain.BACK), err)
		return
	} else {

		a.GetAccountHolder().SetDocument(
			accountdomain.Document{
				AccountID:    a.GetAccountID(),
				Patch:        fileName,
				DocumentType: accountdomain.ACCOUNT_AGREEMENT,
				DocumentSide: accountdomain.BACK,
			}, a.GetAccountHolder().GetIDNumber())
	}
}
