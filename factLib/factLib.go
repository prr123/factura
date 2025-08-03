// https://www.facturae.gob.es/formato/Versiones/Esquema_castellano_v3_2_x_06_06_2017_unificado.pdf


package factLib

import (
	"fmt"
	"os"
	"time"
)

type date struct {
	date time.Time
}
type period struct {
	StartDate date `xml:"StartDate"`
	EndDate date `xml:"EndDate"`
}

type Factura struct {
	// 1
	FileHeader FileHd
	// 2
	Parties InvParties
	// 3
	Invoices Invoice
}

type FileHd struct {
	// 1.1
	SchemaVersion string
	// 1.2
	Modality string
	// 1.3
	InvoiceIssuerType string
	// 1.4
	ThirdParty *TParty
	// 1.5
	Batch *BatchList
	// 1.6
	FactoryAsssignmentData FactAssign	
}

// 2
type InvParties struct {
	// 2.1
	SellerParty Seller
	// 2.2
	BuyerParty Buyer

}

// 1.4
type TParty struct {

}

// 1.5
type BatchList struct {


}

// 1.6
type FactAssign struct {

}

// 2.1
type Buyer struct {
	TaxIdentification TaxId

}

// 2.2
type Seller struct {
	TaxIdentification TaxId

}

// 2.1.1 2.2.1
type TaxId struct {
	// F personal J company judica
	PersonTypeCode byte `xml:"PersonTypeCode"`
	// E outside EU R resident U EU resident
	ResidenceTypeCode byte `xml:"ResidenceTypeCode"`
	//up tp 30 char
	TaxIdentificationNumber string `xml:"TaxIdentificationNumber"`
	PartyIdentification int `xml:"PartyIdentification,omitempty"`
}

type BuyerName struct {
	// max 40 char
	Name string `xml:"Name"`
	FirstSurname string `xml:"FirstSurname,omitempty"`
	SecondSurname string `xml:"SecondSurname,omitempty"`
	AddressInSpain *AddrSpain
	OverseasAddress *AddrOver
	ContactDetails string `xml:"ContactDetails,omitempty"`
	Telephone string `xml:"Telephone,omitempty"`
	TeleFax string `xml:"TeleFax,omitempty"`
	WebAddress string `xml:"WebAddress,omitempty"`
	ElectronicMail string `xml:"ElectronicMail,omitempty"`
	ContactPersons string `xml:"ContactPersons,omitempty"`
	CnoCnae string `xml:"CnoCnae,omitempty"`
	INETownCode string `xml:"INETownCode,omitempty"`
	AdditionalContactDetails string `xml:"AdditionalContactDetails,omitempty"`
}


type AddrSpain struct {
	Address string `xml:"Address"`
	PostCode string `xml:"PostCode"`
	Town string `xml:"Town"`
	Province string `xml:"Province"`
	// ESP 3 let
	CountryCode string `xml:"CountryCode,omitempty"`
}

type AddrOver struct {
    Address string `xml:"Address"`
    PostCode string `xml:"PostCode"`
    Town string `xml:"Town"`
    Province string `xml:"Province"`
    // ESP 3 let
    CountryCode string `xml:"CountryCode,omitempty"`
}

// 3.1
type Invoice struct {
	// 3.1.1
	InvoiceHeader InvHd
	// 3.1.2
	InvoiceIssueData InvData
	// 3.1.3
	TaxesOutputs *TaxOut
	// 3.1.4
	TaxesWithheld *TaxWith
	// 3.1.5
	InvoiceTotals InvTot
	// 3.1.6
	Items InvItems
	// 3.1.7
	PaymentDetails *PayDet
	// 3.1.8
	LegalLiterals *LegalLit
	// 3.1.9
	AdditionalData *AddData
}

// 3.1.1
type InvHd struct {
	// 3.1.1.1
	InvoiceNumber string `xml:"InvoiceNumber"`
	// 3.1.1.2
	InvoiceSeriesCode string `xml:"InvoiceSeriesCode,omitempty"`
	// 3.1.1.3
	InvoiceDocumentType string `xml:"InvoiceDocumentType,omitempty"`
	// 3.1.1.4
	InvoiceClass string `xml:"InvoiceClass,omitempty"`
	// 3.1.1.5
	Corrective *InvCorr
}


// 3.1.2
type InvData struct {
	IssueDate date
	OperationDate *date
	PlaceOfIssue string
	InvoicePeriod period
}

// 3.1.3
type TaxOut struct {
	// 3.1.3.1
	Tax *taxDet
}
//3.1.3.1
type taxDet struct {
	TaxTypeCode string `xml:"TaxTypeCode,omitempty"`
	TaxRate	float64 `xml:"TaxRate,omitempty"`
	// 3.1.3.1.3
	TaxableBase *taxbase
	TaxAmount *taxbase
	// 3.1.3.1.5
	SpecialTaxableBase *taxbase
	SpecialTaxAmount *taxbase
	// 3.1.3.1.7
	EquivalenceSurcharge float64 `xml:"EquivalenceSurcharge,omitempty"`
	EquivalenceSurchargeAmount *taxbase
}

// 3.1.3.1.3
type taxbase struct {
	TotalAmount float64 `xml:"TotalAmount"`
	EquivalentInEuros float64 `xml:"EquivalentInEuros"`
}

// 3.1.4
type TaxWith struct {
	Tax *taxDet
}

// 3.1.4.1
type taxwithdet  struct {
	TaxTypeCode
	TaxRate


}

// 3.1.5
type InvTot struct {
	TotalGrossAmount float64
//	GeneralDiscounts *disc
//	Charge 
}


// 3.1.6
type InvItems struct {
	// 3.1.6.1
	InvoiceLine InvLine
}


type InvLine struct {
	// 3.1.6.1.1
	IssuerContractReference string `xml:"IssuerContractReference,omitempty"`
	IssuerContractDate *date
	IssuerTransactionReference string `xml:"IssuerTransactionReference,omitempty"`
	IssuerTransactionDate *date
	ReceiverContractReference string `xml:"ReceiverContractReference,omitempty"`
	ReceiverContractDate *date
	ReceiverTransactionReference string `xml:"ReceiverTransactionReference,omitempty"`
	ReceiverTransactionDate *date
	FileReference string `xml:"FileReference,omitempty"`
	// 3.1.6.1.10
	FileDate *date `xml:"FileDate,omitempty"`
	// 3.1.6.1.11
	SequenceNumber int64 `xml:"SequenceNumber,omitempty"`
	// 3.1.6.1.12
	DeliveryNotesReference *DelivRef
	// 3.1.6.1.13
	ItemDescription string `xml:"ItemDescription,omitempty"`
	// 3.1.6.1.14
	Quantity int64 `xml:"Quantity"`
	// 3.1.6.1.15
	UnitOfMeasure string `xml:"UnitOfMeasure,omitempty"`
	// 3.1.6.1.16
	UnitPriceWithoutTax float64 `xml:"UnitPriceWithoutTax"`
	// 3.1.6.1.17
	TotalCost double `xml:"TotalCost"`
	// 3.1.6.1.18
	DiscountsAndRebates *ItDisc
	// 3.1.6.1.19
	Charges *Charge
	// 3.1.6.1.20
	GrossAmount float64 `xml:"GrossAmount"`
	// 3.1.6.1.21
	TaxesWithheld *TaxWith
	// 3.1.6.1.22
	TaxesOutputs *ItTaxOut
	// 3.1.6.1.23
	LineItemPeriod *period
	// 3.1.6.1.24
	TransactionDate *date
	// 3.1.6.1.25
	AdditionalLineItemInformation string `xml:"AdditionalLineItemInformation,omitempty"`
	// 3.1.6.1.26
	SpecialTaxableEvent *specTaxEv
	// 3.1.6.1.27
	ArticleCode string `xml:"ArticleCode,omitempty"`
	// 3.1.6.1.28
//	Extensions
}

// 3.1.6.1.12
type DelivRef struct {

}

// 3.1.6.1.18
type Disc struct {
	Discount *DiscDet
}
// 3.1.6.1.18.1
type DiscDet struct {
	DiscountReason string `xml:"DiscountReason,omitempty"`
	DiscountRate float64 `xml:"DiscountRate,omitempty"`
	DiscountAmount float64 `xml:"DiscountAmount"`
}

// 3.1.6.1.19
type Charg struct {
	Charge *CharDet
}
// 3.1.6.1.19.1
type CharDet struct {
	// 3.1.6.1.19.1.1
	ChargeReason string `xml:"ChargeReason,omitempty"`
	ChargeRate float64 `xml:"ChargeRate,omitempty"`
	ChargeAmount float64 `xml:"ChargeAmount"`
}
// 3.1.6.1.22
type ItTaxOut struct {
	Tax string
}
// 3.1.6.1.26
type specTaxEv struct {
	SpecialTaxEventCode string `xml:"SpecialTaxEventCode"`
	SpecialTaxEventReason string `xml:"SpecialTaxEventReason,omitempty"`
}

// 3.1.7 payment details
type PayDet struct {

}

// 3.1.8
type LegalLit struct {

}

// 3.1.9
type AddData struct {

}

// 3.1.1.
// 3.1.1.5
type InvCorr struct {
	InvoiceNumber string `xml:"InvoiceNumber"`
	InvoiceSeriesCode string `xml:"InvoiceSeriesCode,omitempty"`
	ReasonCode string `xml:"ReasonCode"`
	ReasonDescription string `xml:"ReasonDescription"`
	TaxPeriod Period
	CorrectionMethod string `xml:"CorrectionMethod"`
	CorrectionMethodDescription string `xml:"CorrectionMethodDescription"`
	AdditionalReasonDescription string `xml:"AdditionalReasonDescription,omitempty"`
	InvoiceIssueDate date `xml:"InvoiceIssueDate"`
}



// 3.1.2
// 3.1.2.2
type PlaceIss struct {
	PostCode string
	PlaceOfIssueDescription string
}

func PrintInvoice(Inv Factura, out *os.File) {
	fmt.Fprintf(out,"******************* Factura *********************")
	fmt.Fprintf(out,"***************** End Factura *******************")

}


