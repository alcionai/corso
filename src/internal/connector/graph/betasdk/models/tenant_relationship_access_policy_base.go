package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantRelationshipAccessPolicyBase 
type TenantRelationshipAccessPolicyBase struct {
    PolicyBase
    // The definition property
    definition []string
}
// NewTenantRelationshipAccessPolicyBase instantiates a new TenantRelationshipAccessPolicyBase and sets the default values.
func NewTenantRelationshipAccessPolicyBase()(*TenantRelationshipAccessPolicyBase) {
    m := &TenantRelationshipAccessPolicyBase{
        PolicyBase: *NewPolicyBase(),
    }
    odataTypeValue := "#microsoft.graph.tenantRelationshipAccessPolicyBase";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTenantRelationshipAccessPolicyBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantRelationshipAccessPolicyBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.crossTenantAccessPolicy":
                        return NewCrossTenantAccessPolicy(), nil
                }
            }
        }
    }
    return NewTenantRelationshipAccessPolicyBase(), nil
}
// GetDefinition gets the definition property value. The definition property
func (m *TenantRelationshipAccessPolicyBase) GetDefinition()([]string) {
    return m.definition
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantRelationshipAccessPolicyBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["definition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefinition(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TenantRelationshipAccessPolicyBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDefinition() != nil {
        err = writer.WriteCollectionOfStringValues("definition", m.GetDefinition())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefinition sets the definition property value. The definition property
func (m *TenantRelationshipAccessPolicyBase) SetDefinition(value []string)() {
    m.definition = value
}
