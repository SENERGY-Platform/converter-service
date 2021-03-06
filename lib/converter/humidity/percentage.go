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

package humidity

const PercentageId = "urn:infai:ses:characteristic:aaf2385d-92a6-4fd0-ba4d-3e8567e733d6"
const PercentageName = "humidity percentage"

func init() {
	conceptToCharacteristic.Set(PercentageId, func(concept interface{}) (out interface{}, err error) {
		return concept, nil
	})

	characteristicToConcept.Set(PercentageId, func(in interface{}) (concept interface{}, err error) {
		return in, nil
	})
}
