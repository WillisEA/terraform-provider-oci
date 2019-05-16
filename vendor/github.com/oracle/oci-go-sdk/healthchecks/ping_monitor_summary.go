// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Health Checks API
//
// API for the Health Checks service. Use this API to manage endpoint probes and monitors.
// For more information, see
// Overview of the Health Checks Service (https://docs.cloud.oracle.com/iaas/Content/HealthChecks/Concepts/healthchecks.htm).
//

package healthchecks

import (
	"github.com/oracle/oci-go-sdk/common"
)

// PingMonitorSummary This model contains all of the mutable and immutable summary properties for an HTTP monitor.
type PingMonitorSummary struct {

	// The OCID of the resource.
	Id *string `mandatory:"false" json:"id"`

	// A URL for fetching the probe results.
	ResultsUrl *string `mandatory:"false" json:"resultsUrl"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// A user-friendly and mutable name suitable for display in a user interface.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The monitor interval in seconds. Valid values: 10, 30, and 60.
	IntervalInSeconds *int `mandatory:"false" json:"intervalInSeconds"`

	// Enables or disables the monitor. Set to 'true' to launch monitoring.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace.  For more information,
	// see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	Protocol PingMonitorSummaryProtocolEnum `mandatory:"false" json:"protocol,omitempty"`
}

func (m PingMonitorSummary) String() string {
	return common.PointerString(m)
}

// PingMonitorSummaryProtocolEnum Enum with underlying type: string
type PingMonitorSummaryProtocolEnum string

// Set of constants representing the allowable values for PingMonitorSummaryProtocolEnum
const (
	PingMonitorSummaryProtocolIcmp PingMonitorSummaryProtocolEnum = "ICMP"
	PingMonitorSummaryProtocolTcp  PingMonitorSummaryProtocolEnum = "TCP"
)

var mappingPingMonitorSummaryProtocol = map[string]PingMonitorSummaryProtocolEnum{
	"ICMP": PingMonitorSummaryProtocolIcmp,
	"TCP":  PingMonitorSummaryProtocolTcp,
}

// GetPingMonitorSummaryProtocolEnumValues Enumerates the set of values for PingMonitorSummaryProtocolEnum
func GetPingMonitorSummaryProtocolEnumValues() []PingMonitorSummaryProtocolEnum {
	values := make([]PingMonitorSummaryProtocolEnum, 0)
	for _, v := range mappingPingMonitorSummaryProtocol {
		values = append(values, v)
	}
	return values
}
