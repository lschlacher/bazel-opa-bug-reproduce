package main

import (
	"errors"

	"github.com/open-policy-agent/opa/v1/rego"
)

var (
	OPAAccessDeniedErr   error = errors.New("opa access denied")
	OPABindingNotDefined error = errors.New("opa binding not defined")
)

const (
	OPAIsAllowed    = "allow"
	OPASessionAttrs = "session_attrs"
	OPAHeaders      = "headers"
	OPASubject      = "subject"
)

// OpaPDP acts as the Policy Decision Point (PDP) in terms of Open Policy Agent, which loads a policy file and performs access decision
type OpaPDP struct {
	policy          string
	allowQry        rego.PreparedEvalQuery
	sessionAttrsQry rego.PreparedEvalQuery
	headersQry      rego.PreparedEvalQuery
	subjectQry      rego.PreparedEvalQuery
}
