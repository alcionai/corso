package ediscovery

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// Ediscoveryroot 
type Ediscoveryroot struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The cases property
    cases []Case_escapedable
}
// NewEdiscoveryroot instantiates a new Ediscoveryroot and sets the default values.
func NewEdiscoveryroot()(*Ediscoveryroot) {
    m := &Ediscoveryroot{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateEdiscoveryrootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryrootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryroot(), nil
}
// GetCases gets the cases property value. The cases property
func (m *Ediscoveryroot) GetCases()([]Case_escapedable) {
    return m.cases
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Ediscoveryroot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cases"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCase_escapedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Case_escapedable, len(val))
            for i, v := range val {
                res[i] = v.(Case_escapedable)
            }
            m.SetCases(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *Ediscoveryroot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCases() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCases()))
        for i, v := range m.GetCases() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("cases", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCases sets the cases property value. The cases property
func (m *Ediscoveryroot) SetCases(value []Case_escapedable)() {
    m.cases = value
}
