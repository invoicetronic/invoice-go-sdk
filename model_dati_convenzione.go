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

// checks if the DatiConvenzione type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatiConvenzione{}

// DatiConvenzione struct for DatiConvenzione
type DatiConvenzione struct {
	RiferimentoNumeroLinea []int32 `json:"riferimento_numero_linea,omitempty"`
	IdDocumento NullableString `json:"id_documento,omitempty"`
	Data NullableTime `json:"data,omitempty"`
	NumItem NullableString `json:"num_item,omitempty"`
	CodiceCommessaConvenzione NullableString `json:"codice_commessa_convenzione,omitempty"`
	CodiceCup NullableString `json:"codice_cup,omitempty"`
	CodiceCig NullableString `json:"codice_cig,omitempty"`
}

// NewDatiConvenzione instantiates a new DatiConvenzione object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatiConvenzione() *DatiConvenzione {
	this := DatiConvenzione{}
	return &this
}

// NewDatiConvenzioneWithDefaults instantiates a new DatiConvenzione object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatiConvenzioneWithDefaults() *DatiConvenzione {
	this := DatiConvenzione{}
	return &this
}

// GetRiferimentoNumeroLinea returns the RiferimentoNumeroLinea field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiConvenzione) GetRiferimentoNumeroLinea() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.RiferimentoNumeroLinea
}

// GetRiferimentoNumeroLineaOk returns a tuple with the RiferimentoNumeroLinea field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiConvenzione) GetRiferimentoNumeroLineaOk() ([]int32, bool) {
	if o == nil || IsNil(o.RiferimentoNumeroLinea) {
		return nil, false
	}
	return o.RiferimentoNumeroLinea, true
}

// HasRiferimentoNumeroLinea returns a boolean if a field has been set.
func (o *DatiConvenzione) HasRiferimentoNumeroLinea() bool {
	if o != nil && !IsNil(o.RiferimentoNumeroLinea) {
		return true
	}

	return false
}

// SetRiferimentoNumeroLinea gets a reference to the given []int32 and assigns it to the RiferimentoNumeroLinea field.
func (o *DatiConvenzione) SetRiferimentoNumeroLinea(v []int32) {
	o.RiferimentoNumeroLinea = v
}

// GetIdDocumento returns the IdDocumento field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiConvenzione) GetIdDocumento() string {
	if o == nil || IsNil(o.IdDocumento.Get()) {
		var ret string
		return ret
	}
	return *o.IdDocumento.Get()
}

// GetIdDocumentoOk returns a tuple with the IdDocumento field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiConvenzione) GetIdDocumentoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdDocumento.Get(), o.IdDocumento.IsSet()
}

// HasIdDocumento returns a boolean if a field has been set.
func (o *DatiConvenzione) HasIdDocumento() bool {
	if o != nil && o.IdDocumento.IsSet() {
		return true
	}

	return false
}

// SetIdDocumento gets a reference to the given NullableString and assigns it to the IdDocumento field.
func (o *DatiConvenzione) SetIdDocumento(v string) {
	o.IdDocumento.Set(&v)
}
// SetIdDocumentoNil sets the value for IdDocumento to be an explicit nil
func (o *DatiConvenzione) SetIdDocumentoNil() {
	o.IdDocumento.Set(nil)
}

// UnsetIdDocumento ensures that no value is present for IdDocumento, not even an explicit nil
func (o *DatiConvenzione) UnsetIdDocumento() {
	o.IdDocumento.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiConvenzione) GetData() time.Time {
	if o == nil || IsNil(o.Data.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiConvenzione) GetDataOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *DatiConvenzione) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableTime and assigns it to the Data field.
func (o *DatiConvenzione) SetData(v time.Time) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *DatiConvenzione) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *DatiConvenzione) UnsetData() {
	o.Data.Unset()
}

// GetNumItem returns the NumItem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiConvenzione) GetNumItem() string {
	if o == nil || IsNil(o.NumItem.Get()) {
		var ret string
		return ret
	}
	return *o.NumItem.Get()
}

// GetNumItemOk returns a tuple with the NumItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiConvenzione) GetNumItemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumItem.Get(), o.NumItem.IsSet()
}

// HasNumItem returns a boolean if a field has been set.
func (o *DatiConvenzione) HasNumItem() bool {
	if o != nil && o.NumItem.IsSet() {
		return true
	}

	return false
}

// SetNumItem gets a reference to the given NullableString and assigns it to the NumItem field.
func (o *DatiConvenzione) SetNumItem(v string) {
	o.NumItem.Set(&v)
}
// SetNumItemNil sets the value for NumItem to be an explicit nil
func (o *DatiConvenzione) SetNumItemNil() {
	o.NumItem.Set(nil)
}

// UnsetNumItem ensures that no value is present for NumItem, not even an explicit nil
func (o *DatiConvenzione) UnsetNumItem() {
	o.NumItem.Unset()
}

// GetCodiceCommessaConvenzione returns the CodiceCommessaConvenzione field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiConvenzione) GetCodiceCommessaConvenzione() string {
	if o == nil || IsNil(o.CodiceCommessaConvenzione.Get()) {
		var ret string
		return ret
	}
	return *o.CodiceCommessaConvenzione.Get()
}

// GetCodiceCommessaConvenzioneOk returns a tuple with the CodiceCommessaConvenzione field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiConvenzione) GetCodiceCommessaConvenzioneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodiceCommessaConvenzione.Get(), o.CodiceCommessaConvenzione.IsSet()
}

// HasCodiceCommessaConvenzione returns a boolean if a field has been set.
func (o *DatiConvenzione) HasCodiceCommessaConvenzione() bool {
	if o != nil && o.CodiceCommessaConvenzione.IsSet() {
		return true
	}

	return false
}

// SetCodiceCommessaConvenzione gets a reference to the given NullableString and assigns it to the CodiceCommessaConvenzione field.
func (o *DatiConvenzione) SetCodiceCommessaConvenzione(v string) {
	o.CodiceCommessaConvenzione.Set(&v)
}
// SetCodiceCommessaConvenzioneNil sets the value for CodiceCommessaConvenzione to be an explicit nil
func (o *DatiConvenzione) SetCodiceCommessaConvenzioneNil() {
	o.CodiceCommessaConvenzione.Set(nil)
}

// UnsetCodiceCommessaConvenzione ensures that no value is present for CodiceCommessaConvenzione, not even an explicit nil
func (o *DatiConvenzione) UnsetCodiceCommessaConvenzione() {
	o.CodiceCommessaConvenzione.Unset()
}

// GetCodiceCup returns the CodiceCup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiConvenzione) GetCodiceCup() string {
	if o == nil || IsNil(o.CodiceCup.Get()) {
		var ret string
		return ret
	}
	return *o.CodiceCup.Get()
}

// GetCodiceCupOk returns a tuple with the CodiceCup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiConvenzione) GetCodiceCupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodiceCup.Get(), o.CodiceCup.IsSet()
}

// HasCodiceCup returns a boolean if a field has been set.
func (o *DatiConvenzione) HasCodiceCup() bool {
	if o != nil && o.CodiceCup.IsSet() {
		return true
	}

	return false
}

// SetCodiceCup gets a reference to the given NullableString and assigns it to the CodiceCup field.
func (o *DatiConvenzione) SetCodiceCup(v string) {
	o.CodiceCup.Set(&v)
}
// SetCodiceCupNil sets the value for CodiceCup to be an explicit nil
func (o *DatiConvenzione) SetCodiceCupNil() {
	o.CodiceCup.Set(nil)
}

// UnsetCodiceCup ensures that no value is present for CodiceCup, not even an explicit nil
func (o *DatiConvenzione) UnsetCodiceCup() {
	o.CodiceCup.Unset()
}

// GetCodiceCig returns the CodiceCig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiConvenzione) GetCodiceCig() string {
	if o == nil || IsNil(o.CodiceCig.Get()) {
		var ret string
		return ret
	}
	return *o.CodiceCig.Get()
}

// GetCodiceCigOk returns a tuple with the CodiceCig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiConvenzione) GetCodiceCigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodiceCig.Get(), o.CodiceCig.IsSet()
}

// HasCodiceCig returns a boolean if a field has been set.
func (o *DatiConvenzione) HasCodiceCig() bool {
	if o != nil && o.CodiceCig.IsSet() {
		return true
	}

	return false
}

// SetCodiceCig gets a reference to the given NullableString and assigns it to the CodiceCig field.
func (o *DatiConvenzione) SetCodiceCig(v string) {
	o.CodiceCig.Set(&v)
}
// SetCodiceCigNil sets the value for CodiceCig to be an explicit nil
func (o *DatiConvenzione) SetCodiceCigNil() {
	o.CodiceCig.Set(nil)
}

// UnsetCodiceCig ensures that no value is present for CodiceCig, not even an explicit nil
func (o *DatiConvenzione) UnsetCodiceCig() {
	o.CodiceCig.Unset()
}

func (o DatiConvenzione) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatiConvenzione) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RiferimentoNumeroLinea != nil {
		toSerialize["riferimento_numero_linea"] = o.RiferimentoNumeroLinea
	}
	if o.IdDocumento.IsSet() {
		toSerialize["id_documento"] = o.IdDocumento.Get()
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	if o.NumItem.IsSet() {
		toSerialize["num_item"] = o.NumItem.Get()
	}
	if o.CodiceCommessaConvenzione.IsSet() {
		toSerialize["codice_commessa_convenzione"] = o.CodiceCommessaConvenzione.Get()
	}
	if o.CodiceCup.IsSet() {
		toSerialize["codice_cup"] = o.CodiceCup.Get()
	}
	if o.CodiceCig.IsSet() {
		toSerialize["codice_cig"] = o.CodiceCig.Get()
	}
	return toSerialize, nil
}

type NullableDatiConvenzione struct {
	value *DatiConvenzione
	isSet bool
}

func (v NullableDatiConvenzione) Get() *DatiConvenzione {
	return v.value
}

func (v *NullableDatiConvenzione) Set(val *DatiConvenzione) {
	v.value = val
	v.isSet = true
}

func (v NullableDatiConvenzione) IsSet() bool {
	return v.isSet
}

func (v *NullableDatiConvenzione) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatiConvenzione(val *DatiConvenzione) *NullableDatiConvenzione {
	return &NullableDatiConvenzione{value: val, isSet: true}
}

func (v NullableDatiConvenzione) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatiConvenzione) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


