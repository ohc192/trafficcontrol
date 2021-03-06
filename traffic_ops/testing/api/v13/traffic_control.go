/*

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

package v13

import (
	v12 "github.com/apache/trafficcontrol/lib/go-tc"
	"github.com/apache/trafficcontrol/lib/go-tc/v13"
)

// TrafficControl - maps to the tc-fixtures.json file
type TrafficControl struct {
	ASNs                           []v12.ASN                           `json:"asns"`
	CDNs                           []v13.CDN                           `json:"cdns"`
	CacheGroups                    []v13.CacheGroupNullable            `json:"cachegroups"`
	DeliveryServiceRequests        []v12.DeliveryServiceRequest        `json:"deliveryServiceRequests"`
	DeliveryServiceRequestComments []v12.DeliveryServiceRequestComment `json:"deliveryServiceRequestComments"`
	DeliveryServices               []v12.DeliveryServiceV13            `json:"deliveryservices"`
	Divisions                      []v12.Division                      `json:"divisions"`
	Federations                    []v13.CDNFederation                 `json:"federations"`
	Coordinates                    []v13.Coordinate                    `json:"coordinates"`
	Origins                        []v13.Origin                        `json:"origins"`
	Profiles                       []v13.Profile                       `json:"profiles"`
	Parameters                     []v12.Parameter                     `json:"parameters"`
	ProfileParameters              []v13.ProfileParameter              `json:"profileParameters"`
	PhysLocations                  []v12.PhysLocation                  `json:"physLocations"`
	Regions                        []v12.Region                        `json:"regions"`
	Roles                          []v13.Role                          `json:"roles"`
	Servers                        []v13.Server                        `json:"servers"`
	Statuses                       []v12.Status                        `json:"statuses"`
	StaticDNSEntries               []v13.StaticDNSEntry                `json:"staticdnsentries"`
	Tenants                        []v12.Tenant                        `json:"tenants"`
	Types                          []v12.Type                          `json:"types"`
}
