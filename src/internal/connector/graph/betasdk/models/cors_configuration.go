package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CorsConfiguration 
type CorsConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The request headers that the origin domain may specify on the CORS request. The wildcard character * indicates that any header beginning with the specified prefix is allowed.
    allowedHeaders []string
    // The HTTP request methods that the origin domain may use for a CORS request.
    allowedMethods []string
    // The origin domains that are permitted to make a request against the service via CORS. The origin domain is the domain from which the request originates. The origin must be an exact case-sensitive match with the origin that the user age sends to the service.
    allowedOrigins []string
    // The maximum amount of time that a browser should cache the response to the preflight OPTIONS request.
    maxAgeInSeconds *int32
    // The OdataType property
    odataType *string
    // Resource within the application segment for which CORS permissions are granted. / grants permission for whole app segment.
    resource *string
}
// NewCorsConfiguration instantiates a new corsConfiguration and sets the default values.
func NewCorsConfiguration()(*CorsConfiguration) {
    m := &CorsConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCorsConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCorsConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCorsConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CorsConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAllowedHeaders gets the allowedHeaders property value. The request headers that the origin domain may specify on the CORS request. The wildcard character * indicates that any header beginning with the specified prefix is allowed.
func (m *CorsConfiguration) GetAllowedHeaders()([]string) {
    return m.allowedHeaders
}
// GetAllowedMethods gets the allowedMethods property value. The HTTP request methods that the origin domain may use for a CORS request.
func (m *CorsConfiguration) GetAllowedMethods()([]string) {
    return m.allowedMethods
}
// GetAllowedOrigins gets the allowedOrigins property value. The origin domains that are permitted to make a request against the service via CORS. The origin domain is the domain from which the request originates. The origin must be an exact case-sensitive match with the origin that the user age sends to the service.
func (m *CorsConfiguration) GetAllowedOrigins()([]string) {
    return m.allowedOrigins
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CorsConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowedHeaders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedHeaders(res)
        }
        return nil
    }
    res["allowedMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedMethods(res)
        }
        return nil
    }
    res["allowedOrigins"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedOrigins(res)
        }
        return nil
    }
    res["maxAgeInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxAgeInSeconds(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val)
        }
        return nil
    }
    return res
}
// GetMaxAgeInSeconds gets the maxAgeInSeconds property value. The maximum amount of time that a browser should cache the response to the preflight OPTIONS request.
func (m *CorsConfiguration) GetMaxAgeInSeconds()(*int32) {
    return m.maxAgeInSeconds
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CorsConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetResource gets the resource property value. Resource within the application segment for which CORS permissions are granted. / grants permission for whole app segment.
func (m *CorsConfiguration) GetResource()(*string) {
    return m.resource
}
// Serialize serializes information the current object
func (m *CorsConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedHeaders() != nil {
        err := writer.WriteCollectionOfStringValues("allowedHeaders", m.GetAllowedHeaders())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedMethods() != nil {
        err := writer.WriteCollectionOfStringValues("allowedMethods", m.GetAllowedMethods())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedOrigins() != nil {
        err := writer.WriteCollectionOfStringValues("allowedOrigins", m.GetAllowedOrigins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxAgeInSeconds", m.GetMaxAgeInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CorsConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowedHeaders sets the allowedHeaders property value. The request headers that the origin domain may specify on the CORS request. The wildcard character * indicates that any header beginning with the specified prefix is allowed.
func (m *CorsConfiguration) SetAllowedHeaders(value []string)() {
    m.allowedHeaders = value
}
// SetAllowedMethods sets the allowedMethods property value. The HTTP request methods that the origin domain may use for a CORS request.
func (m *CorsConfiguration) SetAllowedMethods(value []string)() {
    m.allowedMethods = value
}
// SetAllowedOrigins sets the allowedOrigins property value. The origin domains that are permitted to make a request against the service via CORS. The origin domain is the domain from which the request originates. The origin must be an exact case-sensitive match with the origin that the user age sends to the service.
func (m *CorsConfiguration) SetAllowedOrigins(value []string)() {
    m.allowedOrigins = value
}
// SetMaxAgeInSeconds sets the maxAgeInSeconds property value. The maximum amount of time that a browser should cache the response to the preflight OPTIONS request.
func (m *CorsConfiguration) SetMaxAgeInSeconds(value *int32)() {
    m.maxAgeInSeconds = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CorsConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetResource sets the resource property value. Resource within the application segment for which CORS permissions are granted. / grants permission for whole app segment.
func (m *CorsConfiguration) SetResource(value *string)() {
    m.resource = value
}
