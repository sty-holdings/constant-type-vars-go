// Package reusable_const_vars
/*
This file contains USA states and postal codes

RESTRICTIONS:
	- Do not edit this comment section.

NOTES:
    To improve code readability, the constant names do not follow camelCase.
	Do not remove IDE inspection directives

COPYRIGHT and WARRANTY:
	Copyright 2022
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.

*/
package reusable_const_vars

//goland:noinspection ALL
const (
	// Extensions
	NC_INTERNAL      = "nc_internal"
	STRIPE_EXTENSION = "stripe"
	//
	// Fully qualified filenames
	FQN_CERTIFICATE          = "CertificateFQN"
	FQN_FIREBASE_CREDENTIALS = "FirebaseCredentialsFQN"
	FQN_GCP_CREDENTIALS      = "GCPCredentialsFQN"
	FQN_NATS_CREDENTIALS     = "NATSCredsFQN"
	FQN_TLS_CABUNDLE         = "TLS:TLSCABundle"
	FQN_TLS_CERTIFICATE      = "TLS:TLSCert"
	FQN_TLS_PRIVATE_KEY      = "TLS:TLSKey"
	// Operators
	OPER_DOUBE_EQUAL_SIGN = "=="
	OPER_EQUAL_SIGN       = "="
	// Text Strings
	// Software
	LOCAL_HOST              = "localhost"
	PID_FILENAME            = "server.pid"
	DEFAULT_VERSION         = "9999.9999.9999"
	ENVIRONMENT_LOCAL       = "local"
	ENVIRONMENT_DEVELOPMENT = "development"
	ENVIRONMENT_PRODUCTION  = "production"
	// Testing
	TEST_POSITIVE_SUCCESS = "Positive Case: Successful: "
	TEST_NEGATIVE_SUCCESS = "Negative Case: Successful: "
	// Unassigned
	ATTEMPT_LIMIT_EXCEEDED  = "Attempt limit exceeded"
	AUTHENTICATOR_SERVICE   = "AuthenticatorService"
	NATS_NON_TLS_CONNECTION = "NON-TLS"
	NATS_TLS_CONNECTION     = "TLS"
	VALID                   = "VALID"
	// Values
	VAL_EMPTY = ""
	VAL_ZERO  = 0
)
