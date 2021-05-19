// Copyright © 2021 Kaleido, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apiserver

import (
	"net/http"

	"github.com/kaleido-io/firefly/internal/config"
	"github.com/kaleido-io/firefly/pkg/fftypes"
	"github.com/kaleido-io/firefly/internal/i18n"
	"github.com/kaleido-io/firefly/internal/oapispec"
)

var postDataDefs = &oapispec.Route{
	Name:   "postDataDefs",
	Path:   "namespaces/{ns}/definitions/data/broadcast",
	Method: http.MethodPost,
	PathParams: []oapispec.PathParam{
		{Name: "ns", ExampleFromConf: config.NamespacesDefault, Description: i18n.MsgTBD},
	},
	QueryParams:     nil,
	FilterFactory:   nil,
	Description:     i18n.MsgTBD,
	JSONInputValue:  func() interface{} { return &fftypes.DataDefinition{} },
	JSONInputMask:   []string{"ID", "Namespace", "Hash", "Created", "Validator"},
	JSONOutputValue: func() interface{} { return &fftypes.Message{} },
	JSONOutputCode:  http.StatusAccepted,
	JSONHandler: func(r oapispec.APIRequest) (output interface{}, err error) {
		output, err = r.Or.BroadcastDataDefinition(r.Ctx, r.PP["ns"], r.Input.(*fftypes.DataDefinition))
		return output, err
	},
}