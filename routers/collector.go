// Copyright 2016 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package routers

import (
	"github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"

	"github.com/Unknwon/orbiter/models"
	"github.com/Unknwon/orbiter/modules/context"
	"github.com/Unknwon/orbiter/modules/form"
)

func Collectors(ctx *context.Context) {
	ctx.Data["Title"] = "Collector"
	ctx.Data["PageIsCollector"] = true

	collectors, err := models.ListCollectors()
	if err != nil {
		ctx.Error(500, err.Error())
		return
	}
	ctx.Data["Collectors"] = collectors

	ctx.HTML(200, "collector/list")
}

func NewCollector(ctx *context.Context) {
	ctx.Data["Title"] = "New Collector"
	ctx.Data["PageIsCollector"] = true
	ctx.HTML(200, "collector/new")
}

type NewCollectorForm struct {
	Name string `binding:"Required;MaxSize(30)" name:"Collector name"`
}

func (f *NewCollectorForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	return form.Validate(errs, ctx.Data, f)
}

func NewCollectorPost(ctx *context.Context, form NewCollectorForm) {
	ctx.Data["Title"] = "New Collector"
	ctx.Data["PageIsCollector"] = true

	if ctx.HasError() {
		ctx.HTML(200, "collector/new")
		return
	}

	_, err := models.NewCollector(form.Name, models.COLLECT_TYPE_GITHUB)
	if err != nil {
		if models.IsErrCollectorExists(err) {
			ctx.Data["Err_Name"] = true
			ctx.RenderWithErr("Collector name has been used.", "collector/new", form)
		} else {
			ctx.Error(500, err.Error())
		}
		return
	}

	ctx.Redirect("/collectors")
}