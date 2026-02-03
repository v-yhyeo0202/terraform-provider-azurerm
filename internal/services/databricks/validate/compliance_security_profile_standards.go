package validate

type ComplianceStandard string

const (
	ComplianceStandardCANADAPROTECTEDB   ComplianceStandard = "CANADA_PROTECTED_B"
	ComplianceStandardCYBERESSENTIALPLUS ComplianceStandard = "CYBER_ESSENTIAL_PLUS"
	ComplianceStandardHIPAA              ComplianceStandard = "HIPAA"
	ComplianceStandardHITRUST            ComplianceStandard = "HITRUST"
	ComplianceStandardIRAP               ComplianceStandard = "IRAP"
	ComplianceStandardKFSI               ComplianceStandard = "K_FSI"
	ComplianceStandardNONE               ComplianceStandard = "NONE"
	ComplianceStandardPCIDSS             ComplianceStandard = "PCI_DSS"
)

func PossibleValuesForComplianceStandard() []string {
	return []string{
		string(ComplianceStandardCANADAPROTECTEDB),
		string(ComplianceStandardCYBERESSENTIALPLUS),
		string(ComplianceStandardHIPAA),
		string(ComplianceStandardHITRUST),
		string(ComplianceStandardIRAP),
		string(ComplianceStandardNONE),
		string(ComplianceStandardPCIDSS),
	}
}

var serverlessComputeModeComplianceStandards = []string{
	string(ComplianceStandardHIPAA),
	string(ComplianceStandardNONE),
}
