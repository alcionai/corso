package termstore

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// Store 
type Store struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // Default language of the term store.
    defaultLanguageTag *string
    // Collection of all groups available in the term store.
    groups []Groupable
    // List of languages for the term store.
    languageTags []string
    // Collection of all sets available in the term store.
    sets []Setable
}
// NewStore instantiates a new store and sets the default values.
func NewStore()(*Store) {
    m := &Store{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateStoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStore(), nil
}
// GetDefaultLanguageTag gets the defaultLanguageTag property value. Default language of the term store.
func (m *Store) GetDefaultLanguageTag()(*string) {
    return m.defaultLanguageTag
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Store) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["defaultLanguageTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLanguageTag(val)
        }
        return nil
    }
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Groupable, len(val))
            for i, v := range val {
                res[i] = v.(Groupable)
            }
            m.SetGroups(res)
        }
        return nil
    }
    res["languageTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetLanguageTags(res)
        }
        return nil
    }
    res["sets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Setable, len(val))
            for i, v := range val {
                res[i] = v.(Setable)
            }
            m.SetSets(res)
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. Collection of all groups available in the term store.
func (m *Store) GetGroups()([]Groupable) {
    return m.groups
}
// GetLanguageTags gets the languageTags property value. List of languages for the term store.
func (m *Store) GetLanguageTags()([]string) {
    return m.languageTags
}
// GetSets gets the sets property value. Collection of all sets available in the term store.
func (m *Store) GetSets()([]Setable) {
    return m.sets
}
// Serialize serializes information the current object
func (m *Store) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("defaultLanguageTag", m.GetDefaultLanguageTag())
        if err != nil {
            return err
        }
    }
    if m.GetGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroups()))
        for i, v := range m.GetGroups() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("groups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLanguageTags() != nil {
        err = writer.WriteCollectionOfStringValues("languageTags", m.GetLanguageTags())
        if err != nil {
            return err
        }
    }
    if m.GetSets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSets()))
        for i, v := range m.GetSets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sets", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultLanguageTag sets the defaultLanguageTag property value. Default language of the term store.
func (m *Store) SetDefaultLanguageTag(value *string)() {
    m.defaultLanguageTag = value
}
// SetGroups sets the groups property value. Collection of all groups available in the term store.
func (m *Store) SetGroups(value []Groupable)() {
    m.groups = value
}
// SetLanguageTags sets the languageTags property value. List of languages for the term store.
func (m *Store) SetLanguageTags(value []string)() {
    m.languageTags = value
}
// SetSets sets the sets property value. Collection of all sets available in the term store.
func (m *Store) SetSets(value []Setable)() {
    m.sets = value
}
