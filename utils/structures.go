package utils

import (
	"encoding/xml"
	"time"
)

var ApiKey string
var FolderPath string
var VatFolderPath string
var Url string = "https://demoefaktura.mfin.gov.rs/"
var UrlPath string = Url + "api/publicApi"
var UrlPathV2 string = Url + "api/v2/publicApi"

//var UrlPath1 string = "https://demoefaktura.mfin.gov.rs/api/publicApi"
//var UrlPath2 string = "https://demoefaktura.mfin.gov.rs/api/v2/publicApi"

var AttachmentPart string = `<cac:AdditionalDocumentReference>
							 	<cbc:ID>NAZIV</cbc:ID>
							 	<cac:Attachment>
								<cbc:EmbeddedDocumentBinaryObject mimeCode="application/pdf">PRILOG</cbc:EmbeddedDocumentBinaryObject>
								</cac:Attachment>
							 </cac:AdditionalDocumentReference>`

/*
	type Message struct {
		Message string `json:"Message"`
	}
*/
type ErrorInvoice struct {
	Message   string `json:"Message"`
	FieldName string `json:"FieldName"`
	ErrorCode string `json:"ErrorCode"`
}

type MiniInvoiceDto struct {
	InvoiceID         int `json:"invoiceId"`
	PurchaseInvoiceID int `json:"purchaseInvoiceId"`
	SalesInvoiceID    int `json:"salesInvoiceId"`
}

type SimpleInvoiceDto struct {
	Status                    string    `json:"status"`
	InvoiceID                 int       `json:"invoiceId"`
	GlobUniqID                string    `json:"globUniqId"`
	Comment                   string    `json:"comment"`
	CirStatus                 string    `json:"cirStatus"`
	CirInvoiceID              string    `json:"cirInvoiceId"`
	Version                   int       `json:"version"`
	LastModifiedUtc           time.Time `json:"lastModifiedUtc"`
	CirSettledAmount          float64   `json:"cirSettledAmount"`
	VatNumberFactoringCompany string    `json:"vatNumberFactoringCompany"`
	FactoringContractNumber   string    `json:"factoringContractNumber"`
	CancelComment             string    `json:"cancelComment"`
	StornoComment             string    `json:"stornoComment"`
}

type SalesInvoicesDto struct {
	SalesInvoiceIds []int `json:"salesInvoiceIds"`
}

type PurchaseInvoicesDto struct {
	PurchaseInvoiceIds []int `json:"purchaseInvoiceIds"`
}

type SalesInvoiceStatusChangeDto []struct {
	EventID             int    `json:"eventId"`
	Date                string `json:"date"`
	NewInvoiceStatus    string `json:"newInvoiceStatus"`
	SalesInvoiceID      int    `json:"salesInvoiceId"`
	Comment             string `json:"comment"`
	CirInvoiceID        string `json:"cirInvoiceId"`
	SubscriptionKey     string `json:"subscriptionKey"`
	StornoNumber        string `json:"stornoNumber"`
	CirAssignmentChange string `json:"cirAssignmentChange"`
	IsSigned            bool   `json:"isSigned"`
}

type PurchaseInvoiceStatusChangeDto []struct {
	EventID             int    `json:"eventId"`
	Date                string `json:"date"`
	NewInvoiceStatus    string `json:"newInvoiceStatus"`
	PurchaseInvoiceID   int    `json:"purchaseInvoiceId"`
	Comment             string `json:"comment"`
	CirInvoiceID        string `json:"cirInvoiceId"`
	SubscriptionKey     string `json:"subscriptionKey"`
	StornoNumber        string `json:"stornoNumber"`
	CirAssignmentChange string `json:"cirAssignmentChange"`
	IsSigned            bool   `json:"isSigned"`
}

type Status struct {
	EventId int
	Status  string
	Comment string
	Date    string
}

type Company struct {
	BugetCompanyNumber  string `json:"BugetCompanyNumber"`
	RegistrationCode    string `json:"RegistrationCode"`
	VatRegistrationCode string `json:"VatRegistrationCode"`
	Name                string `json:"Name"`
}

type UnitMeasures struct {
	Code          string `json:"Code"`
	Symbol        string `json:"Symbol"`
	NameEng       string `json:"NameEng"`
	NameSrbLtn    string `json:"NameSrbLtn"`
	NameSrbCyr    string `json:"NameSrbCyr"`
	IsOnShortList bool   `json:"IsOnShortList"`
}

type IndividualVatDto struct {
	IndividualVatID     int       `json:"individualVatId"`
	CompanyID           int       `json:"companyId"`
	Year                int       `json:"year"`
	DocumentNumber      string    `json:"documentNumber"`
	VatRecordingStatus  string    `json:"vatRecordingStatus"`
	SendDate            time.Time `json:"sendDate"`
	TurnoverDate        time.Time `json:"turnoverDate"`
	PaymentDate         time.Time `json:"paymentDate"`
	DocumentType        string    `json:"documentType"`
	TurnoverDescription string    `json:"turnoverDescription"`
	TurnoverAmount      float64   `json:"turnoverAmount"`
	VatBaseAmount20     float64   `json:"vatBaseAmount20"`
	VatBaseAmount10     float64   `json:"vatBaseAmount10"`
	VatAmount           float64   `json:"vatAmount"`
	VatAmount10         float64   `json:"vatAmount10"`
	VatAmount20         float64   `json:"vatAmount20"`
	TotalAmount         float64   `json:"totalAmount"`
	VatDeductionRight   string    `json:"vatDeductionRight"`
	RelatedDocuments    []struct {
		RelatedVatDocumentID int    `json:"relatedVatDocumentId"`
		DocumentNumber       string `json:"documentNumber"`
	} `json:"relatedDocuments"`
	BasisForPrepayment     string    `json:"basisForPrepayment"`
	DocumentDirection      string    `json:"documentDirection"`
	RelatedPartyIdentifier string    `json:"relatedPartyIdentifier"`
	ForeignDocument        bool      `json:"foreignDocument"`
	TurnoverDescription20  string    `json:"turnoverDescription20"`
	TurnoverDescription10  string    `json:"turnoverDescription10"`
	VatPeriod              string    `json:"vatPeriod"`
	InternalInvoiceOption  string    `json:"internalInvoiceOption"`
	CalculationNumber      string    `json:"calculationNumber"`
	VatRecordingVersion    string    `json:"vatRecordingVersion"`
	CreatedDateUtc         time.Time `json:"createdDateUtc"`
	StatusChangeDateUtc    time.Time `json:"statusChangeDateUtc"`
	IndividualVatHistory   []struct {
		CalculationNumber  string    `json:"calculationNumber"`
		VatRecordingStatus string    `json:"vatRecordingStatus"`
		StatusChangeDate   time.Time `json:"statusChangeDate"`
		TurnoverDate       time.Time `json:"turnoverDate"`
	} `json:"individualVatHistory"`
}

type IndividualVatAddDto struct {
	DocumentNumber      string  `json:"documentNumber"`
	TurnoverDate        string  `json:"turnoverDate"`
	PaymentDate         string  `json:"paymentDate"`
	DocumentType        string  `json:"documentType"`
	Year                int     `json:"year"`
	TurnoverDescription string  `json:"turnoverDescription"`
	TurnoverAmount      float64 `json:"turnoverAmount"`
	VatBaseAmount20     float64 `json:"vatBaseAmount20"`
	VatBaseAmount10     float64 `json:"vatBaseAmount10"`
	VatAmount           float64 `json:"vatAmount"`
	VatAmount10         float64 `json:"vatAmount10"`
	VatAmount20         float64 `json:"vatAmount20"`
	TotalAmount         float64 `json:"totalAmount"`
	VatDeductionRight   string  `json:"vatDeductionRight"`
	RelatedDocuments    []struct {
		DocumentNumber string `json:"documentNumber"`
	} `json:"relatedDocuments"`
	DocumentDirection      string `json:"documentDirection"`
	RelatedPartyIdentifier string `json:"relatedPartyIdentifier"`
	ForeignDocument        bool   `json:"foreignDocument"`
	TurnoverDescription20  string `json:"turnoverDescription20"`
	TurnoverDescription10  string `json:"turnoverDescription10"`
	VatPeriod              string `json:"vatPeriod"`
	InternalInvoiceOption  string `json:"internalInvoiceOption"`
	CalculationNumber      string `json:"calculationNumber"`
	BasisForPrepayment     string `json:"basisForPrepayment"`
}

type IndividualVatDtoV2 struct {
	IndividualVatID                     int       `json:"individualVatId"`
	VatRecordingStatus                  int       `json:"vatRecordingStatus"`
	StatusChangeDate                    time.Time `json:"statusChangeDate"`
	Year                                int       `json:"year"`
	CalculationNumber                   string    `json:"calculationNumber"`
	DocumentNumber                      string    `json:"documentNumber"`
	RecordingDate                       time.Time `json:"recordingDate"`
	VatPeriod                           int       `json:"vatPeriod"`
	DocumentDirection                   int       `json:"documentDirection"`
	DocumentType                        int       `json:"documentType"`
	InternalInvoiceOption               int       `json:"internalInvoiceOption"`
	RelatedPartyIdentifier              string    `json:"relatedPartyIdentifier"`
	RegistrationCode                    string    `json:"registrationCode"`
	VatRegistrationCode                 string    `json:"vatRegistrationCode"`
	NoRealEstateDescription10           string    `json:"noRealEstateDescription10"`
	NoRealEstateBaseAmount10            float64   `json:"noRealEstateBaseAmount10"`
	NoRealEstateCalculatedVat10         float64   `json:"noRealEstateCalculatedVat10"`
	NoRealEstateDescription20           string    `json:"noRealEstateDescription20"`
	NoRealEstateBaseAmount20            float64   `json:"noRealEstateBaseAmount20"`
	NoRealEstateCalculatedVat20         float64   `json:"noRealEstateCalculatedVat20"`
	NoChargeNoRealEstateDescription10   string    `json:"noChargeNoRealEstateDescription10"`
	NoChargeNoRealEstateBaseAmount10    float64   `json:"noChargeNoRealEstateBaseAmount10"`
	NoChargeNoRealEstateCalculatedVat10 float64   `json:"noChargeNoRealEstateCalculatedVat10"`
	NoChargeNoRealEstateDescription20   string    `json:"noChargeNoRealEstateDescription20"`
	NoChargeNoRealEstateBaseAmount20    float64   `json:"noChargeNoRealEstateBaseAmount20"`
	NoChargeNoRealEstateCalculatedVat20 float64   `json:"noChargeNoRealEstateCalculatedVat20"`
	RealEstateDescription10             string    `json:"realEstateDescription10"`
	RealEstateBaseAmount10              float64   `json:"realEstateBaseAmount10"`
	RealEstateCalculatedVat10           float64   `json:"realEstateCalculatedVat10"`
	RealEstateDescription20             string    `json:"realEstateDescription20"`
	RealEstateBaseAmount20              float64   `json:"realEstateBaseAmount20"`
	RealEstateCalculatedVat20           float64   `json:"realEstateCalculatedVat20"`
	NoChargeRealEstateDescription10     string    `json:"noChargeRealEstateDescription10"`
	NoChargeRealEstateBaseAmount10      float64   `json:"noChargeRealEstateBaseAmount10"`
	NoChargeRealEstateCalculatedVat10   float64   `json:"noChargeRealEstateCalculatedVat10"`
	NoChargeRealEstateDescription20     string    `json:"noChargeRealEstateDescription20"`
	NoChargeRealEstateBaseAmount20      float64   `json:"noChargeRealEstateBaseAmount20"`
	NoChargeRealEstateCalculatedVat20   float64   `json:"noChargeRealEstateCalculatedVat20"`
	TotalCalculatedVat                  float64   `json:"totalCalculatedVat"`
	PrepaymentDate                      time.Time `json:"prepaymentDate"`
	RelatedInvoiceOption                int       `json:"relatedInvoiceOption"`
	RelatedVatRecords                   []struct {
		RelatedIndividualVatRecordID int    `json:"relatedIndividualVatRecordId"`
		IndividualVatRecordID        int    `json:"individualVatRecordId"`
		DocumentNumber               string `json:"documentNumber"`
		InternalInvoiceNumber        string `json:"internalInvoiceNumber"`
	} `json:"relatedVatRecords"`
	RelatedVatRecordOutOfSystem  []string  `json:"relatedVatRecordOutOfSystem"`
	RelatedInvoiceIssueDateFrom  time.Time `json:"relatedInvoiceIssueDateFrom"`
	RelatedInvoiceIssueDateTo    time.Time `json:"relatedInvoiceIssueDateTo"`
	RelatedInternalInvoiceOption int       `json:"relatedInternalInvoiceOption"`
	InternalInvoiceNumber        string    `json:"internalInvoiceNumber"`
	BasisForPrepayment           string    `json:"basisForPrepayment"`
	CreatedUtc                   time.Time `json:"createdUtc"`
}

type IndividualVatAddDtoV2 struct {
	Year                                int     `json:"year"`
	CalculationNumber                   string  `json:"calculationNumber"`
	DocumentNumber                      string  `json:"documentNumber"`
	VatPeriod                           int     `json:"vatPeriod"`
	DocumentDirection                   int     `json:"documentDirection"`
	DocumentType                        int     `json:"documentType"`
	InternalInvoiceOption               int     `json:"internalInvoiceOption"`
	RelatedPartyIdentifier              string  `json:"relatedPartyIdentifier"`
	NoRealEstateDescription10           string  `json:"noRealEstateDescription10"`
	NoRealEstateBaseAmount10            float64 `json:"noRealEstateBaseAmount10"`
	NoRealEstateCalculatedVat10         float64 `json:"noRealEstateCalculatedVat10"`
	NoRealEstateDescription20           string  `json:"noRealEstateDescription20"`
	NoRealEstateBaseAmount20            float64 `json:"noRealEstateBaseAmount20"`
	NoRealEstateCalculatedVat20         float64 `json:"noRealEstateCalculatedVat20"`
	NoChargeNoRealEstateDescription10   string  `json:"noChargeNoRealEstateDescription10"`
	NoChargeNoRealEstateBaseAmount10    float64 `json:"noChargeNoRealEstateBaseAmount10"`
	NoChargeNoRealEstateCalculatedVat10 float64 `json:"noChargeNoRealEstateCalculatedVat10"`
	NoChargeNoRealEstateDescription20   string  `json:"noChargeNoRealEstateDescription20"`
	NoChargeNoRealEstateBaseAmount20    float64 `json:"noChargeNoRealEstateBaseAmount20"`
	NoChargeNoRealEstateCalculatedVat20 float64 `json:"noChargeNoRealEstateCalculatedVat20"`
	RealEstateDescription10             string  `json:"realEstateDescription10"`
	RealEstateBaseAmount10              float64 `json:"realEstateBaseAmount10"`
	RealEstateCalculatedVat10           float64 `json:"realEstateCalculatedVat10"`
	RealEstateDescription20             string  `json:"realEstateDescription20"`
	RealEstateBaseAmount20              float64 `json:"realEstateBaseAmount20"`
	RealEstateCalculatedVat20           float64 `json:"realEstateCalculatedVat20"`
	NoChargeRealEstateDescription10     string  `json:"noChargeRealEstateDescription10"`
	NoChargeRealEstateBaseAmount10      float64 `json:"noChargeRealEstateBaseAmount10"`
	NoChargeRealEstateCalculatedVat10   float64 `json:"noChargeRealEstateCalculatedVat10"`
	NoChargeRealEstateDescription20     string  `json:"noChargeRealEstateDescription20"`
	NoChargeRealEstateBaseAmount20      float64 `json:"noChargeRealEstateBaseAmount20"`
	NoChargeRealEstateCalculatedVat20   float64 `json:"noChargeRealEstateCalculatedVat20"`
	PrepaymentDate                      string  `json:"prepaymentDate"`
	RelatedInvoiceOption                int     `json:"relatedInvoiceOption"`
	RelatedVatRecords                   []struct {
		DocumentNumber        string `json:"documentNumber"`
		InternalInvoiceNumber string `json:"internalInvoiceNumber"`
	} `json:"relatedVatRecords"`
	RelatedInternalInvoiceOption int      `json:"relatedInternalInvoiceOption"`
	RelatedVatRecordOutOfSystem  []string `json:"relatedVatRecordOutOfSystem"`
	RelatedInvoiceIssueDateFrom  string   `json:"relatedInvoiceIssueDateFrom"`
	RelatedInvoiceIssueDateTo    string   `json:"relatedInvoiceIssueDateTo"`
	InternalInvoiceNumber        string   `json:"internalInvoiceNumber"`
	BasisForPrepayment           string   `json:"basisForPrepayment"`
}

type InvoiceCustomDto struct {
	XMLName        xml.Name `xml:"DocumentEnvelope"`
	DocumentHeader struct {
		SalesInvoiceID    string `xml:"SalesInvoiceId"`
		PurchaseInvoiceID string `xml:"PurchaseInvoiceId"`
		DocumentID        string `xml:"DocumentId"`
		CreationDate      string `xml:"CreationDate"`
		SendingDate       string `xml:"SendingDate"`
		//DocumentPdf       string `xml:"DocumentPdf"`
	} `xml:"DocumentHeader"`
	DocumentBody struct {
		Invoice struct {
			UBLExtensions struct {
				UBLExtension struct {
					ExtensionContent struct {
						SrbDtExt struct {
							InvoicedPrepaymentAmmount struct {
								ID       string `xml:"ID"`
								TaxTotal struct {
									TaxAmount   string `xml:"TaxAmount"`
									TaxSubtotal []struct {
										TaxableAmount string `xml:"TaxableAmount"`
										TaxAmount     string `xml:"TaxAmount"`
										TaxCategory   struct {
											ID                     string `xml:"ID"`
											Percent                string `xml:"Percent"`
											TaxExemptionReasonCode string `xml:"TaxExemptionReasonCode"`
											TaxExemptionReason     string `xml:"TaxExemptionReason"`
											TaxScheme              struct {
												ID string `xml:"ID"`
											} `xml:"TaxScheme"`
										} `xml:"TaxCategory"`
									} `xml:"TaxSubtotal"`
								} `xml:"TaxTotal"`
							} `xml:"InvoicedPrepaymentAmmount"`
							ReducedTotals struct {
								TaxTotal struct {
									TaxAmount   string `xml:"TaxAmount"`
									TaxSubtotal []struct {
										TaxableAmount string `xml:"TaxableAmount"`
										TaxAmount     string `xml:"TaxAmount"`
										TaxCategory   struct {
											ID                     string `xml:"ID"`
											Percent                string `xml:"Percent"`
											TaxExemptionReasonCode string `xml:"TaxExemptionReasonCode"`
											TaxExemptionReason     string `xml:"TaxExemptionReason"`
											TaxScheme              struct {
												ID string `xml:"ID"`
											} `xml:"TaxScheme"`
										} `xml:"TaxCategory"`
									} `xml:"TaxSubtotal"`
								} `xml:"TaxTotal"`
								LegalMonetaryTotal struct {
									LineExtensionAmount  string `xml:"LineExtensionAmount"`
									TaxExclusiveAmount   string `xml:"TaxExclusiveAmount"`
									TaxInclusiveAmount   string `xml:"TaxInclusiveAmount"`
									AllowanceTotalAmount string `xml:"AllowanceTotalAmount"`
									PrepaidAmount        string `xml:"PrepaidAmount"`
									PayableAmount        string `xml:"PayableAmount"`
								} `xml:"LegalMonetaryTotal"`
							} `xml:"ReducedTotals"`
						} `xml:"SrbDtExt"`
					} `xml:"ExtensionContent"`
				} `xml:"UBLExtension"`
			} `xml:"UBLExtensions"`
			CustomizationID      string `xml:"CustomizationID"`
			ID                   string `xml:"ID"`
			IssueDate            string `xml:"IssueDate"`
			DueDate              string `xml:"DueDate"`
			InvoiceTypeCode      string `xml:"InvoiceTypeCode"`
			CreditNoteTypeCode   string `xml:"CreditNoteTypeCode"`
			Note                 string `xml:"Note"`
			DocumentCurrencyCode string `xml:"DocumentCurrencyCode"`
			TaxCurrencyCode      string `xml:"TaxCurrencyCode"`
			BuyerReference       string `xml:"BuyerReference"`
			InvoicePeriod        struct {
				DescriptionCode string `xml:"DescriptionCode"`
			} `xml:"InvoicePeriod"`
			BillingReference struct {
				InvoiceDocumentReference struct {
					ID        string `xml:"ID"`
					IssueDate string `xml:"IssueDate"`
				} `xml:"InvoiceDocumentReference"`
			} `xml:"BillingReference"`
			OrderReference struct {
				ID string `xml:"ID"`
			} `xml:"OrderReference"`
			OriginatorDocumentReference struct {
				ID string `xml:"ID"`
			} `xml:"OriginatorDocumentReference"`
			ContractDocumentReference struct {
				ID string `xml:"ID"`
			} `xml:"ContractDocumentReference"`
			AdditionalDocumentReference []struct {
				ID               string `xml:"ID"`
				DocumentTypeCode string `xml:"DocumentTypeCode"`
				Attachment       struct {
					ExternalReference struct {
						URI string `xml:"URI"`
					} `xml:"ExternalReference"`
					EmbeddedDocumentBinaryObject string `xml:"EmbeddedDocumentBinaryObject"`
				} `xml:"Attachment"`
			} `xml:"AdditionalDocumentReference"`
			AccountingSupplierParty struct {
				Party struct {
					EndpointID string `xml:"EndpointID"`
					PartyName  struct {
						Name string `xml:"Name"`
					} `xml:"PartyName"`
					PostalAddress struct {
						StreetName string `xml:"StreetName"`
						CityName   string `xml:"CityName"`
						PostalZone string `xml:"PostalZone"`
						Country    struct {
							IdentificationCode string `xml:"IdentificationCode"`
						} `xml:"Country"`
					} `xml:"PostalAddress"`
					PartyTaxScheme struct {
						CompanyID string `xml:"CompanyID"`
						TaxScheme struct {
							ID string `xml:"ID"`
						} `xml:"TaxScheme"`
					} `xml:"PartyTaxScheme"`
					PartyLegalEntity struct {
						RegistrationName string `xml:"RegistrationName"`
						CompanyID        string `xml:"CompanyID"`
					} `xml:"PartyLegalEntity"`
					Contact struct {
						ElectronicMail string `xml:"ElectronicMail"`
					} `xml:"Contact"`
				} `xml:"Party"`
			} `xml:"AccountingSupplierParty"`
			AccountingCustomerParty struct {
				Party struct {
					EndpointID          string `xml:"EndpointID"`
					PartyIdentification struct {
						ID string `xml:"ID"`
					} `xml:"PartyIdentification"`
					PartyName struct {
						Name string `xml:"Name"`
					} `xml:"PartyName"`
					PostalAddress struct {
						StreetName string `xml:"StreetName"`
						CityName   string `xml:"CityName"`
						PostalZone string `xml:"PostalZone"`
						Country    struct {
							IdentificationCode string `xml:"IdentificationCode"`
						} `xml:"Country"`
					} `xml:"PostalAddress"`
					PartyTaxScheme struct {
						CompanyID string `xml:"CompanyID"`
						TaxScheme struct {
							ID string `xml:"ID"`
						} `xml:"TaxScheme"`
					} `xml:"PartyTaxScheme"`
					PartyLegalEntity struct {
						RegistrationName string `xml:"RegistrationName"`
						CompanyID        string `xml:"CompanyID"`
					} `xml:"PartyLegalEntity"`
					Contact struct {
						ElectronicMail string `xml:"ElectronicMail"`
					} `xml:"Contact"`
				} `xml:"Party"`
			} `xml:"AccountingCustomerParty"`
			Delivery struct {
				ActualDeliveryDate string `xml:"ActualDeliveryDate"`
			} `xml:"Delivery"`
			PaymentMeans struct {
				PaymentMeansCode      string `xml:"PaymentMeansCode"`
				PaymentID             string `xml:"PaymentID"`
				PayeeFinancialAccount struct {
					ID string `xml:"ID"`
				} `xml:"PayeeFinancialAccount"`
			} `xml:"PaymentMeans"`
			AllowanceCharge []struct {
				ChargeIndicator string `xml:"ChargeIndicator"`
				Amount          string `xml:"Amount"`
				TaxCategory     struct {
					ID        string `xml:"ID"`
					Percent   string `xml:"Percent"`
					TaxScheme struct {
						ID string `xml:"ID"`
					} `xml:"TaxScheme"`
				} `xml:"TaxCategory"`
			} `xml:"AllowanceCharge"`
			TaxTotal struct {
				TaxAmount   string `xml:"TaxAmount"`
				TaxSubtotal []struct {
					TaxableAmount string `xml:"TaxableAmount"`
					TaxAmount     string `xml:"TaxAmount"`
					TaxCategory   struct {
						ID                     string `xml:"ID"`
						Percent                string `xml:"Percent"`
						TaxExemptionReasonCode string `xml:"TaxExemptionReasonCode"`
						TaxExemptionReason     string `xml:"TaxExemptionReason"`
						TaxScheme              struct {
							ID string `xml:"ID"`
						} `xml:"TaxScheme"`
					} `xml:"TaxCategory"`
				} `xml:"TaxSubtotal"`
			} `xml:"TaxTotal"`
			LegalMonetaryTotal struct {
				LineExtensionAmount  string `xml:"LineExtensionAmount"`
				TaxExclusiveAmount   string `xml:"TaxExclusiveAmount"`
				TaxInclusiveAmount   string `xml:"TaxInclusiveAmount"`
				AllowanceTotalAmount string `xml:"AllowanceTotalAmount"`
				PrepaidAmount        string `xml:"PrepaidAmount"`
				PayableAmount        string `xml:"PayableAmount"`
			} `xml:"LegalMonetaryTotal"`
			InvoiceLine []struct {
				ID                  string `xml:"ID"`
				InvoicedQuantity    string `xml:"InvoicedQuantity"`
				LineExtensionAmount string `xml:"LineExtensionAmount"`
				AllowanceCharge     struct {
					ChargeIndicator         string `xml:"ChargeIndicator"`
					MultiplierFactorNumeric string `xml:"MultiplierFactorNumeric"`
					Amount                  string `xml:"Amount"`
				} `xml:"AllowanceCharge"`
				Item struct {
					Name                      string `xml:"Name"`
					SellersItemIdentification struct {
						ID string `xml:"ID"`
					} `xml:"SellersItemIdentification"`
					ClassifiedTaxCategory struct {
						ID        string `xml:"ID"`
						Percent   string `xml:"Percent"`
						TaxScheme struct {
							ID string `xml:"ID"`
						} `xml:"TaxScheme"`
					} `xml:"ClassifiedTaxCategory"`
				} `xml:"Item"`
				Price struct {
					PriceAmount string `xml:"PriceAmount"`
				} `xml:"Price"`
			} `xml:"InvoiceLine"`
		} `xml:"Invoice"`
	} `xml:"DocumentBody"`
}
