package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsFirewallNetworkProfile windows Firewall Profile Policies.
type WindowsFirewallNetworkProfile struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Configures the firewall to merge authorized application rules from group policy with those from local store instead of ignoring the local store rules. When AuthorizedApplicationRulesFromGroupPolicyNotMerged and AuthorizedApplicationRulesFromGroupPolicyMerged are both true, AuthorizedApplicationRulesFromGroupPolicyMerged takes priority.
    authorizedApplicationRulesFromGroupPolicyMerged *bool
    // Configures the firewall to prevent merging authorized application rules from group policy with those from local store instead of ignoring the local store rules. When AuthorizedApplicationRulesFromGroupPolicyNotMerged and AuthorizedApplicationRulesFromGroupPolicyMerged are both true, AuthorizedApplicationRulesFromGroupPolicyMerged takes priority.
    authorizedApplicationRulesFromGroupPolicyNotMerged *bool
    // Configures the firewall to merge connection security rules from group policy with those from local store instead of ignoring the local store rules. When ConnectionSecurityRulesFromGroupPolicyNotMerged and ConnectionSecurityRulesFromGroupPolicyMerged are both true, ConnectionSecurityRulesFromGroupPolicyMerged takes priority.
    connectionSecurityRulesFromGroupPolicyMerged *bool
    // Configures the firewall to prevent merging connection security rules from group policy with those from local store instead of ignoring the local store rules. When ConnectionSecurityRulesFromGroupPolicyNotMerged and ConnectionSecurityRulesFromGroupPolicyMerged are both true, ConnectionSecurityRulesFromGroupPolicyMerged takes priority.
    connectionSecurityRulesFromGroupPolicyNotMerged *bool
    // State Management Setting.
    firewallEnabled *StateManagementSetting
    // Configures the firewall to merge global port rules from group policy with those from local store instead of ignoring the local store rules. When GlobalPortRulesFromGroupPolicyNotMerged and GlobalPortRulesFromGroupPolicyMerged are both true, GlobalPortRulesFromGroupPolicyMerged takes priority.
    globalPortRulesFromGroupPolicyMerged *bool
    // Configures the firewall to prevent merging global port rules from group policy with those from local store instead of ignoring the local store rules. When GlobalPortRulesFromGroupPolicyNotMerged and GlobalPortRulesFromGroupPolicyMerged are both true, GlobalPortRulesFromGroupPolicyMerged takes priority.
    globalPortRulesFromGroupPolicyNotMerged *bool
    // Configures the firewall to block all incoming connections by default. When InboundConnectionsRequired and InboundConnectionsBlocked are both true, InboundConnectionsBlocked takes priority.
    inboundConnectionsBlocked *bool
    // Configures the firewall to allow all incoming connections by default. When InboundConnectionsRequired and InboundConnectionsBlocked are both true, InboundConnectionsBlocked takes priority.
    inboundConnectionsRequired *bool
    // Prevents the firewall from displaying notifications when an application is blocked from listening on a port. When InboundNotificationsRequired and InboundNotificationsBlocked are both true, InboundNotificationsBlocked takes priority.
    inboundNotificationsBlocked *bool
    // Allows the firewall to display notifications when an application is blocked from listening on a port. When InboundNotificationsRequired and InboundNotificationsBlocked are both true, InboundNotificationsBlocked takes priority.
    inboundNotificationsRequired *bool
    // Configures the firewall to block all incoming traffic regardless of other policy settings. When IncomingTrafficRequired and IncomingTrafficBlocked are both true, IncomingTrafficBlocked takes priority.
    incomingTrafficBlocked *bool
    // Configures the firewall to allow incoming traffic pursuant to other policy settings. When IncomingTrafficRequired and IncomingTrafficBlocked are both true, IncomingTrafficBlocked takes priority.
    incomingTrafficRequired *bool
    // The OdataType property
    odataType *string
    // Configures the firewall to block all outgoing connections by default. When OutboundConnectionsRequired and OutboundConnectionsBlocked are both true, OutboundConnectionsBlocked takes priority. This setting will get applied to Windows releases version 1809 and above.
    outboundConnectionsBlocked *bool
    // Configures the firewall to allow all outgoing connections by default. When OutboundConnectionsRequired and OutboundConnectionsBlocked are both true, OutboundConnectionsBlocked takes priority. This setting will get applied to Windows releases version 1809 and above.
    outboundConnectionsRequired *bool
    // Configures the firewall to merge Firewall Rule policies from group policy with those from local store instead of ignoring the local store rules. When PolicyRulesFromGroupPolicyNotMerged and PolicyRulesFromGroupPolicyMerged are both true, PolicyRulesFromGroupPolicyMerged takes priority.
    policyRulesFromGroupPolicyMerged *bool
    // Configures the firewall to prevent merging Firewall Rule policies from group policy with those from local store instead of ignoring the local store rules. When PolicyRulesFromGroupPolicyNotMerged and PolicyRulesFromGroupPolicyMerged are both true, PolicyRulesFromGroupPolicyMerged takes priority.
    policyRulesFromGroupPolicyNotMerged *bool
    // Configures the firewall to allow the host computer to respond to unsolicited network traffic of that traffic is secured by IPSec even when stealthModeBlocked is set to true. When SecuredPacketExemptionBlocked and SecuredPacketExemptionAllowed are both true, SecuredPacketExemptionAllowed takes priority.
    securedPacketExemptionAllowed *bool
    // Configures the firewall to block the host computer to respond to unsolicited network traffic of that traffic is secured by IPSec even when stealthModeBlocked is set to true. When SecuredPacketExemptionBlocked and SecuredPacketExemptionAllowed are both true, SecuredPacketExemptionAllowed takes priority.
    securedPacketExemptionBlocked *bool
    // Prevent the server from operating in stealth mode. When StealthModeRequired and StealthModeBlocked are both true, StealthModeBlocked takes priority.
    stealthModeBlocked *bool
    // Allow the server to operate in stealth mode. When StealthModeRequired and StealthModeBlocked are both true, StealthModeBlocked takes priority.
    stealthModeRequired *bool
    // Configures the firewall to block unicast responses to multicast broadcast traffic. When UnicastResponsesToMulticastBroadcastsRequired and UnicastResponsesToMulticastBroadcastsBlocked are both true, UnicastResponsesToMulticastBroadcastsBlocked takes priority.
    unicastResponsesToMulticastBroadcastsBlocked *bool
    // Configures the firewall to allow unicast responses to multicast broadcast traffic. When UnicastResponsesToMulticastBroadcastsRequired and UnicastResponsesToMulticastBroadcastsBlocked are both true, UnicastResponsesToMulticastBroadcastsBlocked takes priority.
    unicastResponsesToMulticastBroadcastsRequired *bool
}
// NewWindowsFirewallNetworkProfile instantiates a new windowsFirewallNetworkProfile and sets the default values.
func NewWindowsFirewallNetworkProfile()(*WindowsFirewallNetworkProfile) {
    m := &WindowsFirewallNetworkProfile{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsFirewallNetworkProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsFirewallNetworkProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsFirewallNetworkProfile(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsFirewallNetworkProfile) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAuthorizedApplicationRulesFromGroupPolicyMerged gets the authorizedApplicationRulesFromGroupPolicyMerged property value. Configures the firewall to merge authorized application rules from group policy with those from local store instead of ignoring the local store rules. When AuthorizedApplicationRulesFromGroupPolicyNotMerged and AuthorizedApplicationRulesFromGroupPolicyMerged are both true, AuthorizedApplicationRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) GetAuthorizedApplicationRulesFromGroupPolicyMerged()(*bool) {
    return m.authorizedApplicationRulesFromGroupPolicyMerged
}
// GetAuthorizedApplicationRulesFromGroupPolicyNotMerged gets the authorizedApplicationRulesFromGroupPolicyNotMerged property value. Configures the firewall to prevent merging authorized application rules from group policy with those from local store instead of ignoring the local store rules. When AuthorizedApplicationRulesFromGroupPolicyNotMerged and AuthorizedApplicationRulesFromGroupPolicyMerged are both true, AuthorizedApplicationRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) GetAuthorizedApplicationRulesFromGroupPolicyNotMerged()(*bool) {
    return m.authorizedApplicationRulesFromGroupPolicyNotMerged
}
// GetConnectionSecurityRulesFromGroupPolicyMerged gets the connectionSecurityRulesFromGroupPolicyMerged property value. Configures the firewall to merge connection security rules from group policy with those from local store instead of ignoring the local store rules. When ConnectionSecurityRulesFromGroupPolicyNotMerged and ConnectionSecurityRulesFromGroupPolicyMerged are both true, ConnectionSecurityRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) GetConnectionSecurityRulesFromGroupPolicyMerged()(*bool) {
    return m.connectionSecurityRulesFromGroupPolicyMerged
}
// GetConnectionSecurityRulesFromGroupPolicyNotMerged gets the connectionSecurityRulesFromGroupPolicyNotMerged property value. Configures the firewall to prevent merging connection security rules from group policy with those from local store instead of ignoring the local store rules. When ConnectionSecurityRulesFromGroupPolicyNotMerged and ConnectionSecurityRulesFromGroupPolicyMerged are both true, ConnectionSecurityRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) GetConnectionSecurityRulesFromGroupPolicyNotMerged()(*bool) {
    return m.connectionSecurityRulesFromGroupPolicyNotMerged
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsFirewallNetworkProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authorizedApplicationRulesFromGroupPolicyMerged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizedApplicationRulesFromGroupPolicyMerged(val)
        }
        return nil
    }
    res["authorizedApplicationRulesFromGroupPolicyNotMerged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizedApplicationRulesFromGroupPolicyNotMerged(val)
        }
        return nil
    }
    res["connectionSecurityRulesFromGroupPolicyMerged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionSecurityRulesFromGroupPolicyMerged(val)
        }
        return nil
    }
    res["connectionSecurityRulesFromGroupPolicyNotMerged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionSecurityRulesFromGroupPolicyNotMerged(val)
        }
        return nil
    }
    res["firewallEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseStateManagementSetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallEnabled(val.(*StateManagementSetting))
        }
        return nil
    }
    res["globalPortRulesFromGroupPolicyMerged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGlobalPortRulesFromGroupPolicyMerged(val)
        }
        return nil
    }
    res["globalPortRulesFromGroupPolicyNotMerged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGlobalPortRulesFromGroupPolicyNotMerged(val)
        }
        return nil
    }
    res["inboundConnectionsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInboundConnectionsBlocked(val)
        }
        return nil
    }
    res["inboundConnectionsRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInboundConnectionsRequired(val)
        }
        return nil
    }
    res["inboundNotificationsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInboundNotificationsBlocked(val)
        }
        return nil
    }
    res["inboundNotificationsRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInboundNotificationsRequired(val)
        }
        return nil
    }
    res["incomingTrafficBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncomingTrafficBlocked(val)
        }
        return nil
    }
    res["incomingTrafficRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncomingTrafficRequired(val)
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
    res["outboundConnectionsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutboundConnectionsBlocked(val)
        }
        return nil
    }
    res["outboundConnectionsRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutboundConnectionsRequired(val)
        }
        return nil
    }
    res["policyRulesFromGroupPolicyMerged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyRulesFromGroupPolicyMerged(val)
        }
        return nil
    }
    res["policyRulesFromGroupPolicyNotMerged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyRulesFromGroupPolicyNotMerged(val)
        }
        return nil
    }
    res["securedPacketExemptionAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecuredPacketExemptionAllowed(val)
        }
        return nil
    }
    res["securedPacketExemptionBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecuredPacketExemptionBlocked(val)
        }
        return nil
    }
    res["stealthModeBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStealthModeBlocked(val)
        }
        return nil
    }
    res["stealthModeRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStealthModeRequired(val)
        }
        return nil
    }
    res["unicastResponsesToMulticastBroadcastsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnicastResponsesToMulticastBroadcastsBlocked(val)
        }
        return nil
    }
    res["unicastResponsesToMulticastBroadcastsRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnicastResponsesToMulticastBroadcastsRequired(val)
        }
        return nil
    }
    return res
}
// GetFirewallEnabled gets the firewallEnabled property value. State Management Setting.
func (m *WindowsFirewallNetworkProfile) GetFirewallEnabled()(*StateManagementSetting) {
    return m.firewallEnabled
}
// GetGlobalPortRulesFromGroupPolicyMerged gets the globalPortRulesFromGroupPolicyMerged property value. Configures the firewall to merge global port rules from group policy with those from local store instead of ignoring the local store rules. When GlobalPortRulesFromGroupPolicyNotMerged and GlobalPortRulesFromGroupPolicyMerged are both true, GlobalPortRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) GetGlobalPortRulesFromGroupPolicyMerged()(*bool) {
    return m.globalPortRulesFromGroupPolicyMerged
}
// GetGlobalPortRulesFromGroupPolicyNotMerged gets the globalPortRulesFromGroupPolicyNotMerged property value. Configures the firewall to prevent merging global port rules from group policy with those from local store instead of ignoring the local store rules. When GlobalPortRulesFromGroupPolicyNotMerged and GlobalPortRulesFromGroupPolicyMerged are both true, GlobalPortRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) GetGlobalPortRulesFromGroupPolicyNotMerged()(*bool) {
    return m.globalPortRulesFromGroupPolicyNotMerged
}
// GetInboundConnectionsBlocked gets the inboundConnectionsBlocked property value. Configures the firewall to block all incoming connections by default. When InboundConnectionsRequired and InboundConnectionsBlocked are both true, InboundConnectionsBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) GetInboundConnectionsBlocked()(*bool) {
    return m.inboundConnectionsBlocked
}
// GetInboundConnectionsRequired gets the inboundConnectionsRequired property value. Configures the firewall to allow all incoming connections by default. When InboundConnectionsRequired and InboundConnectionsBlocked are both true, InboundConnectionsBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) GetInboundConnectionsRequired()(*bool) {
    return m.inboundConnectionsRequired
}
// GetInboundNotificationsBlocked gets the inboundNotificationsBlocked property value. Prevents the firewall from displaying notifications when an application is blocked from listening on a port. When InboundNotificationsRequired and InboundNotificationsBlocked are both true, InboundNotificationsBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) GetInboundNotificationsBlocked()(*bool) {
    return m.inboundNotificationsBlocked
}
// GetInboundNotificationsRequired gets the inboundNotificationsRequired property value. Allows the firewall to display notifications when an application is blocked from listening on a port. When InboundNotificationsRequired and InboundNotificationsBlocked are both true, InboundNotificationsBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) GetInboundNotificationsRequired()(*bool) {
    return m.inboundNotificationsRequired
}
// GetIncomingTrafficBlocked gets the incomingTrafficBlocked property value. Configures the firewall to block all incoming traffic regardless of other policy settings. When IncomingTrafficRequired and IncomingTrafficBlocked are both true, IncomingTrafficBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) GetIncomingTrafficBlocked()(*bool) {
    return m.incomingTrafficBlocked
}
// GetIncomingTrafficRequired gets the incomingTrafficRequired property value. Configures the firewall to allow incoming traffic pursuant to other policy settings. When IncomingTrafficRequired and IncomingTrafficBlocked are both true, IncomingTrafficBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) GetIncomingTrafficRequired()(*bool) {
    return m.incomingTrafficRequired
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WindowsFirewallNetworkProfile) GetOdataType()(*string) {
    return m.odataType
}
// GetOutboundConnectionsBlocked gets the outboundConnectionsBlocked property value. Configures the firewall to block all outgoing connections by default. When OutboundConnectionsRequired and OutboundConnectionsBlocked are both true, OutboundConnectionsBlocked takes priority. This setting will get applied to Windows releases version 1809 and above.
func (m *WindowsFirewallNetworkProfile) GetOutboundConnectionsBlocked()(*bool) {
    return m.outboundConnectionsBlocked
}
// GetOutboundConnectionsRequired gets the outboundConnectionsRequired property value. Configures the firewall to allow all outgoing connections by default. When OutboundConnectionsRequired and OutboundConnectionsBlocked are both true, OutboundConnectionsBlocked takes priority. This setting will get applied to Windows releases version 1809 and above.
func (m *WindowsFirewallNetworkProfile) GetOutboundConnectionsRequired()(*bool) {
    return m.outboundConnectionsRequired
}
// GetPolicyRulesFromGroupPolicyMerged gets the policyRulesFromGroupPolicyMerged property value. Configures the firewall to merge Firewall Rule policies from group policy with those from local store instead of ignoring the local store rules. When PolicyRulesFromGroupPolicyNotMerged and PolicyRulesFromGroupPolicyMerged are both true, PolicyRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) GetPolicyRulesFromGroupPolicyMerged()(*bool) {
    return m.policyRulesFromGroupPolicyMerged
}
// GetPolicyRulesFromGroupPolicyNotMerged gets the policyRulesFromGroupPolicyNotMerged property value. Configures the firewall to prevent merging Firewall Rule policies from group policy with those from local store instead of ignoring the local store rules. When PolicyRulesFromGroupPolicyNotMerged and PolicyRulesFromGroupPolicyMerged are both true, PolicyRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) GetPolicyRulesFromGroupPolicyNotMerged()(*bool) {
    return m.policyRulesFromGroupPolicyNotMerged
}
// GetSecuredPacketExemptionAllowed gets the securedPacketExemptionAllowed property value. Configures the firewall to allow the host computer to respond to unsolicited network traffic of that traffic is secured by IPSec even when stealthModeBlocked is set to true. When SecuredPacketExemptionBlocked and SecuredPacketExemptionAllowed are both true, SecuredPacketExemptionAllowed takes priority.
func (m *WindowsFirewallNetworkProfile) GetSecuredPacketExemptionAllowed()(*bool) {
    return m.securedPacketExemptionAllowed
}
// GetSecuredPacketExemptionBlocked gets the securedPacketExemptionBlocked property value. Configures the firewall to block the host computer to respond to unsolicited network traffic of that traffic is secured by IPSec even when stealthModeBlocked is set to true. When SecuredPacketExemptionBlocked and SecuredPacketExemptionAllowed are both true, SecuredPacketExemptionAllowed takes priority.
func (m *WindowsFirewallNetworkProfile) GetSecuredPacketExemptionBlocked()(*bool) {
    return m.securedPacketExemptionBlocked
}
// GetStealthModeBlocked gets the stealthModeBlocked property value. Prevent the server from operating in stealth mode. When StealthModeRequired and StealthModeBlocked are both true, StealthModeBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) GetStealthModeBlocked()(*bool) {
    return m.stealthModeBlocked
}
// GetStealthModeRequired gets the stealthModeRequired property value. Allow the server to operate in stealth mode. When StealthModeRequired and StealthModeBlocked are both true, StealthModeBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) GetStealthModeRequired()(*bool) {
    return m.stealthModeRequired
}
// GetUnicastResponsesToMulticastBroadcastsBlocked gets the unicastResponsesToMulticastBroadcastsBlocked property value. Configures the firewall to block unicast responses to multicast broadcast traffic. When UnicastResponsesToMulticastBroadcastsRequired and UnicastResponsesToMulticastBroadcastsBlocked are both true, UnicastResponsesToMulticastBroadcastsBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) GetUnicastResponsesToMulticastBroadcastsBlocked()(*bool) {
    return m.unicastResponsesToMulticastBroadcastsBlocked
}
// GetUnicastResponsesToMulticastBroadcastsRequired gets the unicastResponsesToMulticastBroadcastsRequired property value. Configures the firewall to allow unicast responses to multicast broadcast traffic. When UnicastResponsesToMulticastBroadcastsRequired and UnicastResponsesToMulticastBroadcastsBlocked are both true, UnicastResponsesToMulticastBroadcastsBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) GetUnicastResponsesToMulticastBroadcastsRequired()(*bool) {
    return m.unicastResponsesToMulticastBroadcastsRequired
}
// Serialize serializes information the current object
func (m *WindowsFirewallNetworkProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("authorizedApplicationRulesFromGroupPolicyMerged", m.GetAuthorizedApplicationRulesFromGroupPolicyMerged())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("authorizedApplicationRulesFromGroupPolicyNotMerged", m.GetAuthorizedApplicationRulesFromGroupPolicyNotMerged())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("connectionSecurityRulesFromGroupPolicyMerged", m.GetConnectionSecurityRulesFromGroupPolicyMerged())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("connectionSecurityRulesFromGroupPolicyNotMerged", m.GetConnectionSecurityRulesFromGroupPolicyNotMerged())
        if err != nil {
            return err
        }
    }
    if m.GetFirewallEnabled() != nil {
        cast := (*m.GetFirewallEnabled()).String()
        err := writer.WriteStringValue("firewallEnabled", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("globalPortRulesFromGroupPolicyMerged", m.GetGlobalPortRulesFromGroupPolicyMerged())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("globalPortRulesFromGroupPolicyNotMerged", m.GetGlobalPortRulesFromGroupPolicyNotMerged())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("inboundConnectionsBlocked", m.GetInboundConnectionsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("inboundConnectionsRequired", m.GetInboundConnectionsRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("inboundNotificationsBlocked", m.GetInboundNotificationsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("inboundNotificationsRequired", m.GetInboundNotificationsRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("incomingTrafficBlocked", m.GetIncomingTrafficBlocked())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("incomingTrafficRequired", m.GetIncomingTrafficRequired())
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
        err := writer.WriteBoolValue("outboundConnectionsBlocked", m.GetOutboundConnectionsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("outboundConnectionsRequired", m.GetOutboundConnectionsRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("policyRulesFromGroupPolicyMerged", m.GetPolicyRulesFromGroupPolicyMerged())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("policyRulesFromGroupPolicyNotMerged", m.GetPolicyRulesFromGroupPolicyNotMerged())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("securedPacketExemptionAllowed", m.GetSecuredPacketExemptionAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("securedPacketExemptionBlocked", m.GetSecuredPacketExemptionBlocked())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("stealthModeBlocked", m.GetStealthModeBlocked())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("stealthModeRequired", m.GetStealthModeRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("unicastResponsesToMulticastBroadcastsBlocked", m.GetUnicastResponsesToMulticastBroadcastsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("unicastResponsesToMulticastBroadcastsRequired", m.GetUnicastResponsesToMulticastBroadcastsRequired())
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
func (m *WindowsFirewallNetworkProfile) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAuthorizedApplicationRulesFromGroupPolicyMerged sets the authorizedApplicationRulesFromGroupPolicyMerged property value. Configures the firewall to merge authorized application rules from group policy with those from local store instead of ignoring the local store rules. When AuthorizedApplicationRulesFromGroupPolicyNotMerged and AuthorizedApplicationRulesFromGroupPolicyMerged are both true, AuthorizedApplicationRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) SetAuthorizedApplicationRulesFromGroupPolicyMerged(value *bool)() {
    m.authorizedApplicationRulesFromGroupPolicyMerged = value
}
// SetAuthorizedApplicationRulesFromGroupPolicyNotMerged sets the authorizedApplicationRulesFromGroupPolicyNotMerged property value. Configures the firewall to prevent merging authorized application rules from group policy with those from local store instead of ignoring the local store rules. When AuthorizedApplicationRulesFromGroupPolicyNotMerged and AuthorizedApplicationRulesFromGroupPolicyMerged are both true, AuthorizedApplicationRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) SetAuthorizedApplicationRulesFromGroupPolicyNotMerged(value *bool)() {
    m.authorizedApplicationRulesFromGroupPolicyNotMerged = value
}
// SetConnectionSecurityRulesFromGroupPolicyMerged sets the connectionSecurityRulesFromGroupPolicyMerged property value. Configures the firewall to merge connection security rules from group policy with those from local store instead of ignoring the local store rules. When ConnectionSecurityRulesFromGroupPolicyNotMerged and ConnectionSecurityRulesFromGroupPolicyMerged are both true, ConnectionSecurityRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) SetConnectionSecurityRulesFromGroupPolicyMerged(value *bool)() {
    m.connectionSecurityRulesFromGroupPolicyMerged = value
}
// SetConnectionSecurityRulesFromGroupPolicyNotMerged sets the connectionSecurityRulesFromGroupPolicyNotMerged property value. Configures the firewall to prevent merging connection security rules from group policy with those from local store instead of ignoring the local store rules. When ConnectionSecurityRulesFromGroupPolicyNotMerged and ConnectionSecurityRulesFromGroupPolicyMerged are both true, ConnectionSecurityRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) SetConnectionSecurityRulesFromGroupPolicyNotMerged(value *bool)() {
    m.connectionSecurityRulesFromGroupPolicyNotMerged = value
}
// SetFirewallEnabled sets the firewallEnabled property value. State Management Setting.
func (m *WindowsFirewallNetworkProfile) SetFirewallEnabled(value *StateManagementSetting)() {
    m.firewallEnabled = value
}
// SetGlobalPortRulesFromGroupPolicyMerged sets the globalPortRulesFromGroupPolicyMerged property value. Configures the firewall to merge global port rules from group policy with those from local store instead of ignoring the local store rules. When GlobalPortRulesFromGroupPolicyNotMerged and GlobalPortRulesFromGroupPolicyMerged are both true, GlobalPortRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) SetGlobalPortRulesFromGroupPolicyMerged(value *bool)() {
    m.globalPortRulesFromGroupPolicyMerged = value
}
// SetGlobalPortRulesFromGroupPolicyNotMerged sets the globalPortRulesFromGroupPolicyNotMerged property value. Configures the firewall to prevent merging global port rules from group policy with those from local store instead of ignoring the local store rules. When GlobalPortRulesFromGroupPolicyNotMerged and GlobalPortRulesFromGroupPolicyMerged are both true, GlobalPortRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) SetGlobalPortRulesFromGroupPolicyNotMerged(value *bool)() {
    m.globalPortRulesFromGroupPolicyNotMerged = value
}
// SetInboundConnectionsBlocked sets the inboundConnectionsBlocked property value. Configures the firewall to block all incoming connections by default. When InboundConnectionsRequired and InboundConnectionsBlocked are both true, InboundConnectionsBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) SetInboundConnectionsBlocked(value *bool)() {
    m.inboundConnectionsBlocked = value
}
// SetInboundConnectionsRequired sets the inboundConnectionsRequired property value. Configures the firewall to allow all incoming connections by default. When InboundConnectionsRequired and InboundConnectionsBlocked are both true, InboundConnectionsBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) SetInboundConnectionsRequired(value *bool)() {
    m.inboundConnectionsRequired = value
}
// SetInboundNotificationsBlocked sets the inboundNotificationsBlocked property value. Prevents the firewall from displaying notifications when an application is blocked from listening on a port. When InboundNotificationsRequired and InboundNotificationsBlocked are both true, InboundNotificationsBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) SetInboundNotificationsBlocked(value *bool)() {
    m.inboundNotificationsBlocked = value
}
// SetInboundNotificationsRequired sets the inboundNotificationsRequired property value. Allows the firewall to display notifications when an application is blocked from listening on a port. When InboundNotificationsRequired and InboundNotificationsBlocked are both true, InboundNotificationsBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) SetInboundNotificationsRequired(value *bool)() {
    m.inboundNotificationsRequired = value
}
// SetIncomingTrafficBlocked sets the incomingTrafficBlocked property value. Configures the firewall to block all incoming traffic regardless of other policy settings. When IncomingTrafficRequired and IncomingTrafficBlocked are both true, IncomingTrafficBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) SetIncomingTrafficBlocked(value *bool)() {
    m.incomingTrafficBlocked = value
}
// SetIncomingTrafficRequired sets the incomingTrafficRequired property value. Configures the firewall to allow incoming traffic pursuant to other policy settings. When IncomingTrafficRequired and IncomingTrafficBlocked are both true, IncomingTrafficBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) SetIncomingTrafficRequired(value *bool)() {
    m.incomingTrafficRequired = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsFirewallNetworkProfile) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOutboundConnectionsBlocked sets the outboundConnectionsBlocked property value. Configures the firewall to block all outgoing connections by default. When OutboundConnectionsRequired and OutboundConnectionsBlocked are both true, OutboundConnectionsBlocked takes priority. This setting will get applied to Windows releases version 1809 and above.
func (m *WindowsFirewallNetworkProfile) SetOutboundConnectionsBlocked(value *bool)() {
    m.outboundConnectionsBlocked = value
}
// SetOutboundConnectionsRequired sets the outboundConnectionsRequired property value. Configures the firewall to allow all outgoing connections by default. When OutboundConnectionsRequired and OutboundConnectionsBlocked are both true, OutboundConnectionsBlocked takes priority. This setting will get applied to Windows releases version 1809 and above.
func (m *WindowsFirewallNetworkProfile) SetOutboundConnectionsRequired(value *bool)() {
    m.outboundConnectionsRequired = value
}
// SetPolicyRulesFromGroupPolicyMerged sets the policyRulesFromGroupPolicyMerged property value. Configures the firewall to merge Firewall Rule policies from group policy with those from local store instead of ignoring the local store rules. When PolicyRulesFromGroupPolicyNotMerged and PolicyRulesFromGroupPolicyMerged are both true, PolicyRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) SetPolicyRulesFromGroupPolicyMerged(value *bool)() {
    m.policyRulesFromGroupPolicyMerged = value
}
// SetPolicyRulesFromGroupPolicyNotMerged sets the policyRulesFromGroupPolicyNotMerged property value. Configures the firewall to prevent merging Firewall Rule policies from group policy with those from local store instead of ignoring the local store rules. When PolicyRulesFromGroupPolicyNotMerged and PolicyRulesFromGroupPolicyMerged are both true, PolicyRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) SetPolicyRulesFromGroupPolicyNotMerged(value *bool)() {
    m.policyRulesFromGroupPolicyNotMerged = value
}
// SetSecuredPacketExemptionAllowed sets the securedPacketExemptionAllowed property value. Configures the firewall to allow the host computer to respond to unsolicited network traffic of that traffic is secured by IPSec even when stealthModeBlocked is set to true. When SecuredPacketExemptionBlocked and SecuredPacketExemptionAllowed are both true, SecuredPacketExemptionAllowed takes priority.
func (m *WindowsFirewallNetworkProfile) SetSecuredPacketExemptionAllowed(value *bool)() {
    m.securedPacketExemptionAllowed = value
}
// SetSecuredPacketExemptionBlocked sets the securedPacketExemptionBlocked property value. Configures the firewall to block the host computer to respond to unsolicited network traffic of that traffic is secured by IPSec even when stealthModeBlocked is set to true. When SecuredPacketExemptionBlocked and SecuredPacketExemptionAllowed are both true, SecuredPacketExemptionAllowed takes priority.
func (m *WindowsFirewallNetworkProfile) SetSecuredPacketExemptionBlocked(value *bool)() {
    m.securedPacketExemptionBlocked = value
}
// SetStealthModeBlocked sets the stealthModeBlocked property value. Prevent the server from operating in stealth mode. When StealthModeRequired and StealthModeBlocked are both true, StealthModeBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) SetStealthModeBlocked(value *bool)() {
    m.stealthModeBlocked = value
}
// SetStealthModeRequired sets the stealthModeRequired property value. Allow the server to operate in stealth mode. When StealthModeRequired and StealthModeBlocked are both true, StealthModeBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) SetStealthModeRequired(value *bool)() {
    m.stealthModeRequired = value
}
// SetUnicastResponsesToMulticastBroadcastsBlocked sets the unicastResponsesToMulticastBroadcastsBlocked property value. Configures the firewall to block unicast responses to multicast broadcast traffic. When UnicastResponsesToMulticastBroadcastsRequired and UnicastResponsesToMulticastBroadcastsBlocked are both true, UnicastResponsesToMulticastBroadcastsBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) SetUnicastResponsesToMulticastBroadcastsBlocked(value *bool)() {
    m.unicastResponsesToMulticastBroadcastsBlocked = value
}
// SetUnicastResponsesToMulticastBroadcastsRequired sets the unicastResponsesToMulticastBroadcastsRequired property value. Configures the firewall to allow unicast responses to multicast broadcast traffic. When UnicastResponsesToMulticastBroadcastsRequired and UnicastResponsesToMulticastBroadcastsBlocked are both true, UnicastResponsesToMulticastBroadcastsBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) SetUnicastResponsesToMulticastBroadcastsRequired(value *bool)() {
    m.unicastResponsesToMulticastBroadcastsRequired = value
}
