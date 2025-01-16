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
)

// checks if the DatiPagamento type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatiPagamento{}

// DatiPagamento struct for DatiPagamento
type DatiPagamento struct {
	CondizioniPagamento NullableString `json:"condizioni_pagamento,omitempty"`
	DettaglioPagamento []DettaglioPagamento `json:"dettaglio_pagamento,omitempty"`
}

// NewDatiPagamento instantiates a new DatiPagamento object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatiPagamento() *DatiPagamento {
	this := DatiPagamento{}
	return &this
}

// NewDatiPagamentoWithDefaults instantiates a new DatiPagamento object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatiPagamentoWithDefaults() *DatiPagamento {
	this := DatiPagamento{}
	return &this
}

// GetCondizioniPagamento returns the CondizioniPagamento field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiPagamento) GetCondizioniPagamento() string {
	if o == nil || IsNil(o.CondizioniPagamento.Get()) {
		var ret string
		return ret
	}
	return *o.CondizioniPagamento.Get()
}

// GetCondizioniPagamentoOk returns a tuple with the CondizioniPagamento field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiPagamento) GetCondizioniPagamentoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CondizioniPagamento.Get(), o.CondizioniPagamento.IsSet()
}

// HasCondizioniPagamento returns a boolean if a field has been set.
func (o *DatiPagamento) HasCondizioniPagamento() bool {
	if o != nil && o.CondizioniPagamento.IsSet() {
		return true
	}

	return false
}

// SetCondizioniPagamento gets a reference to the given NullableString and assigns it to the CondizioniPagamento field.
func (o *DatiPagamento) SetCondizioniPagamento(v string) {
	o.CondizioniPagamento.Set(&v)
}
// SetCondizioniPagamentoNil sets the value for CondizioniPagamento to be an explicit nil
func (o *DatiPagamento) SetCondizioniPagamentoNil() {
	o.CondizioniPagamento.Set(nil)
}

// UnsetCondizioniPagamento ensures that no value is present for CondizioniPagamento, not even an explicit nil
func (o *DatiPagamento) UnsetCondizioniPagamento() {
	o.CondizioniPagamento.Unset()
}

// GetDettaglioPagamento returns the DettaglioPagamento field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiPagamento) GetDettaglioPagamento() []DettaglioPagamento {
	if o == nil {
		var ret []DettaglioPagamento
		return ret
	}
	return o.DettaglioPagamento
}

// GetDettaglioPagamentoOk returns a tuple with the DettaglioPagamento field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiPagamento) GetDettaglioPagamentoOk() ([]DettaglioPagamento, bool) {
	if o == nil || IsNil(o.DettaglioPagamento) {
		return nil, false
	}
	return o.DettaglioPagamento, true
}

// HasDettaglioPagamento returns a boolean if a field has been set.
func (o *DatiPagamento) HasDettaglioPagamento() bool {
	if o != nil && !IsNil(o.DettaglioPagamento) {
		return true
	}

	return false
}

// SetDettaglioPagamento gets a reference to the given []DettaglioPagamento and assigns it to the DettaglioPagamento field.
func (o *DatiPagamento) SetDettaglioPagamento(v []DettaglioPagamento) {
	o.DettaglioPagamento = v
}

func (o DatiPagamento) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatiPagamento) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CondizioniPagamento.IsSet() {
		toSerialize["condizioni_pagamento"] = o.CondizioniPagamento.Get()
	}
	if o.DettaglioPagamento != nil {
		toSerialize["dettaglio_pagamento"] = o.DettaglioPagamento
	}
	return toSerialize, nil
}

type NullableDatiPagamento struct {
	value *DatiPagamento
	isSet bool
}

func (v NullableDatiPagamento) Get() *DatiPagamento {
	return v.value
}

func (v *NullableDatiPagamento) Set(val *DatiPagamento) {
	v.value = val
	v.isSet = true
}

func (v NullableDatiPagamento) IsSet() bool {
	return v.isSet
}

func (v *NullableDatiPagamento) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatiPagamento(val *DatiPagamento) *NullableDatiPagamento {
	return &NullableDatiPagamento{value: val, isSet: true}
}

func (v NullableDatiPagamento) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatiPagamento) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


