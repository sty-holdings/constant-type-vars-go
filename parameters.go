// Package constant_type_vars
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
package constant_type_vars

import (
	"fmt"
	"strings"
)

//goland:noinspection ALL
const (
	PARAMETER_NATS_TOKEN     = "nats-token"
	PARAMETER_WEBSOCKET_HOST = "websocket-host"
)

var (
	ParameterNameFormat = "%v-%v-%v"
)

func GetParameterName(parameterType, programName, environment string) string {

	switch strings.ToLower(strings.Trim(parameterType, SPACES_ONE)) {
	case PARAMETER_NATS_TOKEN:
		return fmt.Sprintf(ParameterNameFormat, programName, environment, PARAMETER_NATS_TOKEN)
	case PARAMETER_WEBSOCKET_HOST:
		return fmt.Sprintf(ParameterNameFormat, programName, environment, PARAMETER_WEBSOCKET_HOST)
	default:
		return fmt.Sprintf("%v%v", TXT_MISSING_PARAMETER, FN_PARAMETER_TYPE)
	}
}
