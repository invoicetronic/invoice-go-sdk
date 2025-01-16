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

// checks if the RappresentanteFiscaleCessionarioCommittente type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RappresentanteFiscaleCessionarioCommittente{}

// RappresentanteFiscaleCessionarioCommittente struct for RappresentanteFiscaleCessionarioCommittente
type RappresentanteFiscaleCessionarioCommittente struct {
	Denominazione NullableString `json:"denominazione,omitempty"`
	Nome NullableString `json:"nome,omitempty"`
	Cognome NullableString `json:"cognome,omitempty"`
	IdFiscaleIva *IdFiscaleIVA `json:"id_fiscale_iva,omitempty"`
}

// NewRappresentanteFiscaleCessionarioCommittente instantiates a new RappresentanteFiscaleCessionarioCommittente object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRappresentanteFiscaleCessionarioCommittente() *RappresentanteFiscaleCessionarioCommittente {
	this := RappresentanteFiscaleCessionarioCommittente{}
	return &this
}

// NewRappresentanteFiscaleCessionarioCommittenteWithDefaults instantiates a new RappresentanteFiscaleCessionarioCommittente object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRappresentanteFiscaleCessionarioCommittenteWithDefaults() *RappresentanteFiscaleCessionarioCommittente {
	this := RappresentanteFiscaleCessionarioCommittente{}
	return &this
}

// GetDenominazione returns the Denominazione field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RappresentanteFiscaleCessionarioCommittente) GetDenominazione() string {
	if o == nil || IsNil(o.Denominazione.Get()) {
		var ret string
		return ret
	}
	return *o.Denominazione.Get()
}

// GetDenominazioneOk returns a tuple with the Denominazione field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RappresentanteFiscaleCessionarioCommittente) GetDenominazioneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Denominazione.Get(), o.Denominazione.IsSet()
}

// HasDenominazione returns a boolean if a field has been set.
func (o *RappresentanteFiscaleCessionarioCommittente) HasDenominazione() bool {
	if o != nil && o.Denominazione.IsSet() {
		return true
	}

	return false
}

// SetDenominazione gets a reference to the given NullableString and assigns it to the Denominazione field.
func (o *RappresentanteFiscaleCessionarioCommittente) SetDenominazione(v string) {
	o.Denominazione.Set(&v)
}
// SetDenominazioneNil sets the value for Denominazione to be an explicit nil
func (o *RappresentanteFiscaleCessionarioCommittente) SetDenominazioneNil() {
	o.Denominazione.Set(nil)
}

// UnsetDenominazione ensures that no value is present for Denominazione, not even an explicit nil
func (o *RappresentanteFiscaleCessionarioCommittente) UnsetDenominazione() {
	o.Denominazione.Unset()
}

// GetNome returns the Nome field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RappresentanteFiscaleCessionarioCommittente) GetNome() string {
	if o == nil || IsNil(o.Nome.Get()) {
		var ret string
		return ret
	}
	return *o.Nome.Get()
}

// GetNomeOk returns a tuple with the Nome field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RappresentanteFiscaleCessionarioCommittente) GetNomeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nome.Get(), o.Nome.IsSet()
}

// HasNome returns a boolean if a field has been set.
func (o *RappresentanteFiscaleCessionarioCommittente) HasNome() bool {
	if o != nil && o.Nome.IsSet() {
		return true
	}

	return false
}

// SetNome gets a reference to the given NullableString and assigns it to the Nome field.
func (o *RappresentanteFiscaleCessionarioCommittente) SetNome(v string) {
	o.Nome.Set(&v)
}
// SetNomeNil sets the value for Nome to be an explicit nil
func (o *RappresentanteFiscaleCessionarioCommittente) SetNomeNil() {
	o.Nome.Set(nil)
}

// UnsetNome ensures that no value is present for Nome, not even an explicit nil
func (o *RappresentanteFiscaleCessionarioCommittente) UnsetNome() {
	o.Nome.Unset()
}

// GetCognome returns the Cognome field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RappresentanteFiscaleCessionarioCommittente) GetCognome() string {
	if o == nil || IsNil(o.Cognome.Get()) {
		var ret string
		return ret
	}
	return *o.Cognome.Get()
}

// GetCognomeOk returns a tuple with the Cognome field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RappresentanteFiscaleCessionarioCommittente) GetCognomeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cognome.Get(), o.Cognome.IsSet()
}

// HasCognome returns a boolean if a field has been set.
func (o *RappresentanteFiscaleCessionarioCommittente) HasCognome() bool {
	if o != nil && o.Cognome.IsSet() {
		return true
	}

	return false
}

// SetCognome gets a reference to the given NullableString and assigns it to the Cognome field.
func (o *RappresentanteFiscaleCessionarioCommittente) SetCognome(v string) {
	o.Cognome.Set(&v)
}
// SetCognomeNil sets the value for Cognome to be an explicit nil
func (o *RappresentanteFiscaleCessionarioCommittente) SetCognomeNil() {
	o.Cognome.Set(nil)
}

// UnsetCognome ensures that no value is present for Cognome, not even an explicit nil
func (o *RappresentanteFiscaleCessionarioCommittente) UnsetCognome() {
	o.Cognome.Unset()
}

// GetIdFiscaleIva returns the IdFiscaleIva field value if set, zero value otherwise.
func (o *RappresentanteFiscaleCessionarioCommittente) GetIdFiscaleIva() IdFiscaleIVA {
	if o == nil || IsNil(o.IdFiscaleIva) {
		var ret IdFiscaleIVA
		return ret
	}
	return *o.IdFiscaleIva
}

// GetIdFiscaleIvaOk returns a tuple with the IdFiscaleIva field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RappresentanteFiscaleCessionarioCommittente) GetIdFiscaleIvaOk() (*IdFiscaleIVA, bool) {
	if o == nil || IsNil(o.IdFiscaleIva) {
		return nil, false
	}
	return o.IdFiscaleIva, true
}

// HasIdFiscaleIva returns a boolean if a field has been set.
func (o *RappresentanteFiscaleCessionarioCommittente) HasIdFiscaleIva() bool {
	if o != nil && !IsNil(o.IdFiscaleIva) {
		return true
	}

	return false
}

// SetIdFiscaleIva gets a reference to the given IdFiscaleIVA and assigns it to the IdFiscaleIva field.
func (o *RappresentanteFiscaleCessionarioCommittente) SetIdFiscaleIva(v IdFiscaleIVA) {
	o.IdFiscaleIva = &v
}

func (o RappresentanteFiscaleCessionarioCommittente) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RappresentanteFiscaleCessionarioCommittente) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Denominazione.IsSet() {
		toSerialize["denominazione"] = o.Denominazione.Get()
	}
	if o.Nome.IsSet() {
		toSerialize["nome"] = o.Nome.Get()
	}
	if o.Cognome.IsSet() {
		toSerialize["cognome"] = o.Cognome.Get()
	}
	if !IsNil(o.IdFiscaleIva) {
		toSerialize["id_fiscale_iva"] = o.IdFiscaleIva
	}
	return toSerialize, nil
}

type NullableRappresentanteFiscaleCessionarioCommittente struct {
	value *RappresentanteFiscaleCessionarioCommittente
	isSet bool
}

func (v NullableRappresentanteFiscaleCessionarioCommittente) Get() *RappresentanteFiscaleCessionarioCommittente {
	return v.value
}

func (v *NullableRappresentanteFiscaleCessionarioCommittente) Set(val *RappresentanteFiscaleCessionarioCommittente) {
	v.value = val
	v.isSet = true
}

func (v NullableRappresentanteFiscaleCessionarioCommittente) IsSet() bool {
	return v.isSet
}

func (v *NullableRappresentanteFiscaleCessionarioCommittente) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRappresentanteFiscaleCessionarioCommittente(val *RappresentanteFiscaleCessionarioCommittente) *NullableRappresentanteFiscaleCessionarioCommittente {
	return &NullableRappresentanteFiscaleCessionarioCommittente{value: val, isSet: true}
}

func (v NullableRappresentanteFiscaleCessionarioCommittente) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRappresentanteFiscaleCessionarioCommittente) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


