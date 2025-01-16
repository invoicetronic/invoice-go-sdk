/*
Italian eInvoice API

The Italian eInvoice API is a RESTful API that allows you to send and receive invoices through the Italian [Servizio di Interscambio (SDI)][1], or Interchange Service. The API is designed by Invoicetronic to be simple and easy to use, abstracting away SDI complexity while providing complete control over the invoice send/receive process. The API also provides advanced features as encryption at rest, invoice validation, multiple upload formats, webhooks, event logging, client SDKs for commonly used languages, and CLI tools.  For more information, see  [Invoicetronic website][2]  [1]: https://www.fatturapa.gov.it/it/sistemainterscambio/cose-il-sdi/ [2]: https://invoicetronic.com/

API version: 1.0.0
Contact: support@invoicetronic.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package invoicesdk

import (
	"encoding/json"
	"time"
)

// checks if the DatiVeicoli type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatiVeicoli{}

// DatiVeicoli struct for DatiVeicoli
type DatiVeicoli struct {
	Data NullableTime `json:"data,omitempty"`
	TotalePercorso NullableString `json:"totale_percorso,omitempty"`
}

// NewDatiVeicoli instantiates a new DatiVeicoli object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatiVeicoli() *DatiVeicoli {
	this := DatiVeicoli{}
	return &this
}

// NewDatiVeicoliWithDefaults instantiates a new DatiVeicoli object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatiVeicoliWithDefaults() *DatiVeicoli {
	this := DatiVeicoli{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiVeicoli) GetData() time.Time {
	if o == nil || IsNil(o.Data.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiVeicoli) GetDataOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *DatiVeicoli) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableTime and assigns it to the Data field.
func (o *DatiVeicoli) SetData(v time.Time) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *DatiVeicoli) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *DatiVeicoli) UnsetData() {
	o.Data.Unset()
}

// GetTotalePercorso returns the TotalePercorso field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiVeicoli) GetTotalePercorso() string {
	if o == nil || IsNil(o.TotalePercorso.Get()) {
		var ret string
		return ret
	}
	return *o.TotalePercorso.Get()
}

// GetTotalePercorsoOk returns a tuple with the TotalePercorso field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiVeicoli) GetTotalePercorsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalePercorso.Get(), o.TotalePercorso.IsSet()
}

// HasTotalePercorso returns a boolean if a field has been set.
func (o *DatiVeicoli) HasTotalePercorso() bool {
	if o != nil && o.TotalePercorso.IsSet() {
		return true
	}

	return false
}

// SetTotalePercorso gets a reference to the given NullableString and assigns it to the TotalePercorso field.
func (o *DatiVeicoli) SetTotalePercorso(v string) {
	o.TotalePercorso.Set(&v)
}
// SetTotalePercorsoNil sets the value for TotalePercorso to be an explicit nil
func (o *DatiVeicoli) SetTotalePercorsoNil() {
	o.TotalePercorso.Set(nil)
}

// UnsetTotalePercorso ensures that no value is present for TotalePercorso, not even an explicit nil
func (o *DatiVeicoli) UnsetTotalePercorso() {
	o.TotalePercorso.Unset()
}

func (o DatiVeicoli) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatiVeicoli) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	if o.TotalePercorso.IsSet() {
		toSerialize["totale_percorso"] = o.TotalePercorso.Get()
	}
	return toSerialize, nil
}

type NullableDatiVeicoli struct {
	value *DatiVeicoli
	isSet bool
}

func (v NullableDatiVeicoli) Get() *DatiVeicoli {
	return v.value
}

func (v *NullableDatiVeicoli) Set(val *DatiVeicoli) {
	v.value = val
	v.isSet = true
}

func (v NullableDatiVeicoli) IsSet() bool {
	return v.isSet
}

func (v *NullableDatiVeicoli) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatiVeicoli(val *DatiVeicoli) *NullableDatiVeicoli {
	return &NullableDatiVeicoli{value: val, isSet: true}
}

func (v NullableDatiVeicoli) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatiVeicoli) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


