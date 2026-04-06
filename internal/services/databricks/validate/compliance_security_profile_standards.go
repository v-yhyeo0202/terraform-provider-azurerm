// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package validate

type ComplianceStandard string

const (
	ComplianceStandardCANADAPROTECTEDB   ComplianceStandard = "CANADA_PROTECTED_B"
	ComplianceStandardCYBERESSENTIALPLUS ComplianceStandard = "CYBER_ESSENTIAL_PLUS"
	ComplianceStandardFEDRAMPHIGH        ComplianceStandard = "FEDRAMP_HIGH"
	ComplianceStandardGERMANYC5          ComplianceStandard = "GERMANY_C5"
	ComplianceStandardGERMANYTISAX       ComplianceStandard = "GERMANY_TISAX"
	ComplianceStandardHIPAA              ComplianceStandard = "HIPAA"
	ComplianceStandardHITRUST            ComplianceStandard = "HITRUST"
	ComplianceStandardIRAPPROTECTED      ComplianceStandard = "IRAP_PROTECTED"
	ComplianceStandardISMAP              ComplianceStandard = "ISMAP"
	ComplianceStandardKFSI               ComplianceStandard = "K_FSI"
	ComplianceStandardNONE               ComplianceStandard = "NONE"
	ComplianceStandardPCIDSS             ComplianceStandard = "PCI_DSS"
)

func PossibleValuesForHybridComputeModeComplianceStandard() []string {
	return []string{
		string(ComplianceStandardCANADAPROTECTEDB),
		string(ComplianceStandardCYBERESSENTIALPLUS),
		string(ComplianceStandardFEDRAMPHIGH),
		string(ComplianceStandardGERMANYC5),
		string(ComplianceStandardGERMANYTISAX),
		string(ComplianceStandardHITRUST),
		string(ComplianceStandardHIPAA),
		string(ComplianceStandardIRAPPROTECTED),
		string(ComplianceStandardISMAP),
		string(ComplianceStandardKFSI),
		string(ComplianceStandardPCIDSS),
	}
}

func PossibleValuesForServerlessComputeModeComplianceStandard() []string {
	return []string{
		string(ComplianceStandardHIPAA),
	}
}
