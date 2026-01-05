// Copyright 2025 EURECOM
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
//
// Contributors:
//   Giulio CAROTA
//   Thomas DU
//   Adlen KSENTINI

package contexts

type app interface {
}

type TraffInflAppCtx struct {
	app
	AppFunc map[string]*AppFunctionCtx
}

func NewTraffInflCtx(app app) *TraffInflAppCtx {
	ctx := &TraffInflAppCtx{app: app,
		AppFunc: make(map[string]*AppFunctionCtx),
	}
	return ctx
}

func (c *TraffInflAppCtx) GetAf(id string) *AppFunctionCtx {
	af, ok := c.AppFunc[id]
	if ok {
		return af
	} else {
		return nil
	}
}

func (c *TraffInflAppCtx) AddAf(id string) *AppFunctionCtx {
	_, ok := c.AppFunc[id]
	if ok {
		return nil
	} else {
		newCtx := NewAf(id)
		c.AppFunc[id] = newCtx
		return newCtx
	}
}
