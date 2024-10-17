package layer2

import (
	accountdomain "digital-bank/internal/account/domain"
	"digital-bank/internal/system/infrastructure/service_credentials"
	"digital-bank/pkg"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
)

type (
	Layer2Application struct {
		c *Layer2
	}
)

func NewLayer2Application(credentials credentials.Layer2Credentials) *Layer2Application {
	return &Layer2Application{
		c: NewLayer2(credentials),
	}
}

func (l *Layer2Application) CreateApplication(a *accountdomain.Account) error {
	l.c.payload = ApplicationPayloadPrepare(a)
	l.c.endPointURL = "v1/applications"

	bodyResp, err := l.c.Post()
	if err != nil {
		log.Println("Error on create application", err)
		return err
	}

	var res applicationResponse
	if err := json.Unmarshal(bodyResp, &res); err != nil {
		return err
	}

	a.SetApplicationID(res.data.id)

	generateContract(a)

	if a.GetAccountHolder().GetType() == accountdomain.COMPANY_CLIENT {
		l.sendPartner(a)
	}

	applicationStatus, err := l.lookApplicationStatus(a.GetApplicationID())

	if err != nil {
		return err
	}

	if len(applicationStatus.applicationDocumentErrors) > 0 {
		err = l.sendDocument(a.GetAccountHolder().GetDocuments(), applicationStatus.applicationDocumentErrors)
		if err != nil {
			return err
		}
	}

	if len(applicationStatus.applicationIndividualErrors) > 0 && a.GetAccountHolder().GetType() == accountdomain.INDIVIDUAL_CLIENT {
		err = l.sendDocument(a.GetAccountHolder().GetDocuments(), applicationStatus.applicationIndividualErrors[0].documentErrors)
		if err != nil {
			return err
		}
	}

	if a.GetType() == accountdomain.COMPANY_CLIENT {
		holder := a.GetAccountHolder().(accountdomain.CompanyAccountHolder)
		err = l.sendDocumentsPartners(applicationStatus.applicationIndividualErrors, holder.GetPartners())
		if err != nil {
			return err
		}
	}

	err = l.submitApplication(a)
	if err != nil {
		return err
	}

	return nil
}

func (l *Layer2Application) submitApplication(a *accountdomain.Account) error {

	l.c.endPointURL = "v1/applications/" + a.GetApplicationID() + "/submit"
	_, err := l.c.Post()
	if err != nil {
		log.Println("error submitApplication", err)
		return err
	}

	a.Processing()

	return nil
}

func (l *Layer2Application) sendDocumentsPartners(applicationErrors []applicationIndividualErrors, partners []accountdomain.Individual) error {
	for _, partner := range partners {
		for _, applicationError := range applicationErrors {
			if applicationError.individualID == partner.PartnerApplicationId {
				err := l.sendDocument(partner.Documents, applicationError.documentErrors)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (l *Layer2Application) sendPartner(a *accountdomain.Account) {
	holder := a.GetAccountHolder().(accountdomain.CompanyAccountHolder)

	partners := holder.GetPartners()

	for _, partner := range partners {
		l.c.endPointURL = "v1/applications/" + a.GetApplicationID() + "/individual"
		l.c.payload = PartnerPayloadPrepare(partner)

		bodyResp, err := l.c.Post()
		if err != nil {
			log.Println("error send partner", partner.ToMap())
			continue

		}

		var res applicationResponse
		if err := json.Unmarshal(bodyResp, &res); err != nil {
			continue
		}

		partner.PartnerApplicationId = res.data.id

		holder.EditPartner(partner.DNI, partner)

	}
}

func (l *Layer2Application) lookApplicationStatus(applicationID string) (layer2ApplicationStatusResponse, error) {
	l.c.endPointURL = "v1/applications/" + applicationID + "/status"
	bodyResp, err := l.c.Get()
	if err != nil {
		log.Println("error lookApplicationStatus: ", err)
		return layer2ApplicationStatusResponse{}, err
	}

	var res layer2ApplicationStatusResponse
	if err := json.Unmarshal(bodyResp, &res); err != nil {
		return layer2ApplicationStatusResponse{}, err

	}

	return res, nil
}

func (l *Layer2Application) sendDocument(docs []accountdomain.Document, docErrors []ApplicationDocumentError) error {

	client := resty.New()

	for _, document := range docs {
		layer2DocConverted := convertDocumentTypeToLayer2(document.DocumentType, document.GetDocumentSide())

		if layer2DocConverted == "" {
			continue
		}

		Layer2DocumentID := lookLayer2DocumentID(docErrors, layer2DocConverted)

		if Layer2DocumentID == nil {
			continue
		}

		url, err := pkg.NewAWSS3Storage(os.Getenv("AWS_S3_BUCKET_DOCUMENT")).DownloadFile(document.Patch)
		if err != nil {
			return err
		}

		resp, err := client.R().Get(url)
		if err != nil {
			return fmt.Errorf("failed to download file: %w", err)
		}

		if resp.IsError() {
			return fmt.Errorf("error response from server: %s", resp.String())
		}

		err = l.c.SendDocument(resp.Body(), *Layer2DocumentID)
		if err != nil {
			return err

		}

	}

	return nil
}
