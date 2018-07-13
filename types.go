package sparkpost

// SendingDomains because the API was being dumb
type SendingDomains struct {
	SendingDomainResult []SendingDomainResult `json:"results"`
}

// SendingDomain struct because that's how the sparkpost api returns it
type SendingDomain struct {
	SendingDomainResult SendingDomainResult `json:"results"`
}

// SendingDomainResult type for the sending domains config in sparkpost
type SendingDomainResult struct {
	// omitempty so we can share the same structs for both retrieving all sending domains and just a single one
	Dkim                  SenderDkim   `json:"dkim,omitempty"`
	Domain                string       `json:"domain"`
	Status                SenderStatus `json:"status"`
	IsDefaultBounceDomain bool         `json:"is_default_bounce_domain"`
	CreationTime          string       `json:"creation_time"`
	SharedWithSubaccounts bool         `json:"shared_with_subaccounts"`
}

// SenderStatus type for capturing the status of a sender domain in sparkpost
type SenderStatus struct {
	MxStatus                  string `json:"mx_status"`
	SpfStatus                 string `json:"spf_status"`
	CnameStatus               string `json:"cname_status"`
	OwnershipVerfied          bool   `json:"ownership_verified"`
	AbuseAtStatus             string `json:"abuse_at_status"`
	ComplianceStatus          string `json:"compliance_status"`
	VerificationMailboxStatus string `json:"verification_mailbox_status"`
	DkimStatus                string `json:"dkim_status"`
	PostmansterAtStatus       string `json:"postmaster_at_status"`
}

// SenderDkim type for returning dkim information about a sparkpost sender
type SenderDkim struct {
	SigningDomain string `json:"signing_domain"`
	Headers       string `json:"headers"`
	Public        string `json:"public"`
	Selector      string `json:"selector"`
}

// TrackingDomainResults results wrapper
type TrackingDomainResults struct {
	Results []TrackingDomain `json:"results"`
}

// TrackingDomain representing a tracking domain
type TrackingDomain struct {
	Default      bool   `json:"default"`
	Domain       string `json:"domain"`
	SubaccountID int    `json:"subaccount_id,omitempty"`
}
