package report

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

var instanceReportMtx = sync.Mutex{}

type InstanceReport struct {
	Events []InstanceReportEvent `json:"events"`
}

type InstanceReportEvent struct {
	ReportedAt                int64  `json:"reported_at"`
	LicenseID                 string `json:"license_id"`
	InstanceID                string `json:"instance_id"`
	ClusterID                 string `json:"cluster_id"`
	UserAgent                 string `json:"user_agent"`
	AppStatus                 string `json:"app_status,omitempty"`
	ResourceStates            string `json:"resource_states,omitempty"`
	K8sVersion                string `json:"k8s_version"`
	K8sDistribution           string `json:"k8s_distribution,omitempty"`
	DownstreamChannelID       string `json:"downstream_channel_id,omitempty"`
	DownstreamChannelSequence int64  `json:"downstream_channel_sequence"`
	DownstreamChannelName     string `json:"downstream_channel_name,omitempty"`
	Tags                      string `json:"tags"`
}

func (r *InstanceReport) GetType() ReportType {
	return ReportTypeInstance
}

func (r *InstanceReport) GetSecretName() string {
	return fmt.Sprintf(ReportSecretNameFormat, r.GetType())
}

func (r *InstanceReport) GetSecretKey() string {
	return ReportSecretKey
}

func (r *InstanceReport) AppendEvents(report Report) error {
	reportToAppend, ok := report.(*InstanceReport)
	if !ok {
		return errors.Errorf("report is not an instance report")
	}

	r.Events = append(r.Events, reportToAppend.Events...)
	if len(r.Events) > r.GetEventLimit() {
		r.Events = r.Events[len(r.Events)-r.GetEventLimit():]
	}

	// remove one event at a time until the report is under the size limit
	encoded, err := EncodeReport(r)
	if err != nil {
		return errors.Wrap(err, "failed to encode report")
	}
	for len(encoded) > r.GetSizeLimit() {
		r.Events = r.Events[1:]
		if len(r.Events) == 0 {
			return errors.Errorf("size of latest event exceeds report size limit")
		}
		encoded, err = EncodeReport(r)
		if err != nil {
			return errors.Wrap(err, "failed to encode report")
		}
	}

	return nil
}

func (r *InstanceReport) GetEventLimit() int {
	return ReportEventLimit
}

func (r *InstanceReport) GetSizeLimit() int {
	return ReportSizeLimit
}

func (r *InstanceReport) GetMtx() *sync.Mutex {
	return &instanceReportMtx
}
