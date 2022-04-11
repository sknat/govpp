//  Copyright (c) 2020 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package binapigen

import (
	"embed"
	"fmt"
	"io"
	"text/template"
)

type templateDef struct {
	filename     string
	templateName string
	deps         []dependency
}

type dependency struct {
	apiName            string
	versionConstraints string
}

type WrapperRenderArgs struct {
	PackageName string
	ImportPath  string
}

// A wrapper template will be generated if all of its api dependencies
// are also generated and their version satisfy the constraint
// Templates that expose the same API (for different binapi versions)
// must have mutually exclusive dependency constraints
// Typically, v1 will depend on api1 >= 1.0.1,<2.0 and v2 will depend
// on api1 >= 2.0
var (
	wrapperTemplateDB = []templateDef{
		{"vpp.go", "vpp_v1.go.tmpl", []dependency{}},
		{"ip_helpers.go", "ip_helpers_v1.go.tmpl", []dependency{{"ip_types", ">= 3.0.0"}}},
		{"ipip.go", "ipip_v1.go.tmpl", []dependency{{"ip_types", ">= 3.0.0"}, {"interface_types", ">= 1.0.0"}, {"ipip", ">= 2.0.2"}}},
		{"ipsec.go", "ipsec_v1.go.tmpl", []dependency{{"ipsec_types", ">= 3.0.1"}, {"interface_types", ">= 1.0.0"}, {"ipsec", ">= 5.0.2"}}},
		{"ipsec_helpers.go", "ipsec_helpers_v1.go.tmpl", []dependency{{"ipsec_types", ">= 3.0.1"}, {"tunnel_types", ">= 1.0.1"}}},
	}
)

//go:embed wrapper_templates/*
var content embed.FS
var templates *template.Template = nil

func renderTemplate(name string, args WrapperRenderArgs, out io.Writer) (err error) {
	if templates == nil {
		templates, err = template.ParseFS(content, "wrapper_templates/*")
		if err != nil {
			return fmt.Errorf("error loading templates: %v", err)
		}
	}
	return templates.ExecuteTemplate(out, name, args)
}
