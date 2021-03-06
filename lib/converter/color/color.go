/*
 * Copyright 2019 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package color

import (
	"github.com/SENERGY-Platform/converter/lib/converter/base"
	"github.com/SENERGY-Platform/converter/lib/model"
)

var characteristicToConcept = &base.CastCharacteristicToConcept{}
var conceptToCharacteristic = &base.CastConceptToCharacteristic{}

const Color = "urn:infai:ses:concept:8b1161d5-7878-4dd2-a36c-6f98f6b94bf8"

func init() {
	base.Concepts[Color] = base.GetConceptCastFunction(characteristicToConcept, conceptToCharacteristic)
	base.ConceptRepo.Register(model.Concept{Id: Color, Name: "color", BaseCharacteristicId: Rgb}, []model.Characteristic{
		{
			Id:   Rgb,
			Name: "RGB",
			Type: model.Structure,
			SubCharacteristics: []model.Characteristic{
				{Id: RgbR, Name: "r", Type: model.Integer},
				{Id: RgbG, Name: "g", Type: model.Integer},
				{Id: RgbB, Name: "b", Type: model.Integer},
			},
		},
		{
			Id:   Hex,
			Name: "hex",
			Type: model.String,
		},
		{
			Id:   Hsb,
			Name: "HSB",
			Type: model.Structure,
			SubCharacteristics: []model.Characteristic{
				{Id: HsbH, Name: "h", Type: model.Integer},
				{Id: HsbS, Name: "s", Type: model.Integer},
				{Id: HsbB, Name: "b", Type: model.Integer},
			},
		},
	})
}
