package systemusecase

import (
	systemdomain "digital-bank/internal/system/domain"
	systemreq "digital-bank/internal/system/infrastructure/requests"
	"fmt"
)

type (
	OnboardingAppClient struct {
		repoAppClient       systemdomain.AppClientRepository
		repoSystemParameter systemdomain.SystemParametersRepository
		clintID             systemdomain.EntityID
	}
)

func NewOnboardingAppClient(clintID systemdomain.EntityID, repoAppClient systemdomain.AppClientRepository, repoSystemParameter systemdomain.SystemParametersRepository) *OnboardingAppClient {
	return &OnboardingAppClient{
		repoAppClient:       repoAppClient,
		repoSystemParameter: repoSystemParameter,
		clintID:             clintID,
	}
}

func (o *OnboardingAppClient) Run(appClientRequest systemreq.AppClientRequest) systemdomain.Result[*systemdomain.AppClient] {
	systemParameters, err := o.repoSystemParameter.GetSystemParameters()

	if err != nil {
		return systemdomain.NewResult[*systemdomain.AppClient](nil, systemdomain.NewError(500, err.Error()))
	}

	domesticUSA := systemdomain.DomesticUSA{
		ACH: struct {
			IN  float64 `json:"in"`
			OUT float64 `json:"out"`
		}{
			IN:  appClientRequest.Commissions.DomesticUSA.ACH.IN,
			OUT: appClientRequest.Commissions.DomesticUSA.ACH.OUT,
		},
		FedWire: struct {
			IN  float64 `json:"in"`
			OUT float64 `json:"out"`
		}{
			IN:  appClientRequest.Commissions.DomesticUSA.FedWire.IN,
			OUT: appClientRequest.Commissions.DomesticUSA.FedWire.OUT,
		},
	}

	swift := systemdomain.SwiftUSA{
		IN:  appClientRequest.Commissions.SwiftUSA.IN,
		OUT: appClientRequest.Commissions.SwiftUSA.OUT,
	}

	swap := systemdomain.Swap{
		Buy:  appClientRequest.Commissions.Swap.Buy,
		Sell: appClientRequest.Commissions.Swap.Sell,
	}

	commissionsChargedClient := systemdomain.NewTransactionFee(domesticUSA, swift, swap)

	appClient := systemdomain.NewAppClient(
		o.clintID, appClientRequest.Name, appClientRequest.PhoneNumber, appClientRequest.Email, commissionsChargedClient,
		systemParameters.GetCommissions(),
	)

	err = o.repoAppClient.Upsert(appClient)

	if err != nil {
		fmt.Println(`OnboardingAppClient Error:`, err)
		return systemdomain.NewResult[*systemdomain.AppClient](nil, systemdomain.NewError(500, err.Error()))
	}

	return systemdomain.NewResult[*systemdomain.AppClient](appClient, nil)
}
