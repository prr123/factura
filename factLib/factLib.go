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
	ThirdParty TParty
	// 1.5
	Batch BatchList
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
	AddressInSpain AddrSpain
	OverseasAddress AddrOver
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
	TaxesOutputs TaxOut
	// 3.1.4
	TaxesWithheld TaxWith
	// 3.1.5
	InvoiceTotals InvTot
	// 3.1.6
	Items InvItems
	// 3.1.7
	PaymentDetails PayDet
	// 3.1.8
	LegalLiterals LegalLit
	// 3.1.9
	AdditionalData AddData
	
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
	Corrective InvCorr
}


// 3.1.2
type InvData struct {
	IssueDate date
	OperationDate date
	PlaceOfIssue string
	InvoicePeriod Period
}

// 3.1.3
type TaxOut struct {

}

// 3.1.4
type TaxWith struct {

}

// 3.1.5
type InvTot struct {

}

// 3.1.6
type InvItems struct {

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


type Period struct {
	StartDate date `xml:"StartDate"`
	EndDate date `xml:"EndDate"`
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


