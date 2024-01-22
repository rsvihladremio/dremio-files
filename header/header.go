//      Copyright 2023 Dremio Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package header provides shared parsing of a query profile header.json
package header

type Header struct {
	ClusterInfo   ClusterInfo `json:"clusterInfo"`
	Submission    Submission  `json:"submission"`
	Job           Job         `json:"job"`
	DremioVersion string      `json:"dremioVersion"`
}

type ClusterInfo struct {
	Source        []Source `json:"source"`
	Node          []Node   `json:"node"`
	JavaVmVersion string   `json:"javaVmVersion"`
	JreVersion    string   `json:"jreVersion"`
	Edition       string   `json:"edition"`
	Identity      Identity `json:"identity"`
	Version       Version  `json:"version"`
}

type Source struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Node struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type Identity struct {
	Identity string      `json:"identity"`
	Created  float64     `json:"created"`
	Version  FullVersion `json:"version"`
}

type FullVersion struct {
	Patch       int64  `json:"patch"`
	BuildNumber int64  `json:"buildNumber"`
	Qualifier   string `json:"qualifier"`
	Major       int64  `json:"major"`
	Minor       int64  `json:"minor"`
}

type Version struct {
	Version string `json:"version"`
}

type Submission struct {
	SubmissionId string  `json:"submissionId"`
	First        string  `json:"first"`
	Last         string  `json:"last"`
	Email        string  `json:"email"`
	Date         float64 `json:"date"`
}

type Job struct {
	AttemptId           string      `json:"attemptId"`
	Endpoint            Endpoint    `json:"endpoint"`
	AccelerationDetails string      `json:"accelerationDetails"`
	StateList           []StateList `json:"stateList"`
	State               int64       `json:"state"`
	Info                Info        `json:"info"`
	Stats               Stats       `json:"stats"`
	Details             Details     `json:"details"`
}

type Endpoint struct {
	ConduitPort     int64         `json:"conduitPort"`
	DremioVersion   string        `json:"dremioVersion"`
	UserPort        int64         `json:"userPort"`
	FabricPort      int64         `json:"fabricPort"`
	StartTime       float64       `json:"startTime"`
	NodeTag         string        `json:"nodeTag"`
	Address         string        `json:"address"`
	Roles           EndpointRoles `json:"roles"`
	MaxDirectMemory float64       `json:"maxDirectMemory"`
	AvailableCores  int64         `json:"availableCores"`
}

type EndpointRoles struct {
	SqlQuery     bool `json:"sqlQuery"`
	JavaExecutor bool `json:"javaExecutor"`
	Master       bool `json:"master"`
}

type StateList struct {
	State     int64   `json:"state"`
	StartTime float64 `json:"startTime"`
}

type Info struct {
	JobId                  JobId                  `json:"jobId"`
	Sql                    string                 `json:"sql"`
	User                   string                 `json:"user"`
	StartTime              float64                `json:"startTime"`
	DatasetVersion         string                 `json:"datasetVersion"`
	Description            string                 `json:"description"`
	ScanPaths              []ScanPaths            `json:"scanPaths"`
	ExecutionCpuTimeNs     int64                  `json:"executionCpuTimeNs"`
	Parents                []Parents              `json:"parents"`
	JoinAnalysis           JoinAnalysis           `json:"joinAnalysis"`
	CommandPoolWaitMillis  int64                  `json:"commandPoolWaitMillis"`
	IsTruncatedSql         bool                   `json:"isTruncatedSql"`
	Context                []string               `json:"context"`
	OutputTable            []string               `json:"outputTable"`
	FinishTime             float64                `json:"finishTime"`
	FieldOrigins           []FieldOrigins         `json:"fieldOrigins"`
	ResultMetadata         []ResultMetadata       `json:"resultMetadata"`
	ResourceSchedulingInfo ResourceSchedulingInfo `json:"resourceSchedulingInfo"`
	SetupTimeNs            float64                `json:"setupTimeNs"`
	WaitTimeNs             int64                  `json:"waitTimeNs"`
	RequestType            int64                  `json:"requestType"`
	GrandParents           []GrandParents         `json:"grandParents"`
	OriginalCost           float64                `json:"originalCost"`
	DatasetPath            []string               `json:"datasetPath"`
	BatchSchema            string                 `json:"batchSchema"`
	NodeDetails            []NodeDetails          `json:"nodeDetails"`
	MemoryAllocated        float64                `json:"memoryAllocated"`
	QueryType              int64                  `json:"queryType"`
}

type JobId struct {
	Id string `json:"id"`
}

type ScanPaths struct {
	Path []string `json:"path"`
}

type Parents struct {
	Type        int64    `json:"type"`
	DatasetPath []string `json:"datasetPath"`
}

type JoinAnalysis struct {
}

type FieldOrigins struct {
	Name    string    `json:"name"`
	Origins []Origins `json:"origins"`
}

type Origins struct {
	Table      []string `json:"table"`
	ColumnName string   `json:"columnName"`
	Derived    bool     `json:"derived"`
}

type ResultMetadata struct {
	Path                 string             `json:"path"`
	Footer               Footer             `json:"footer"`
	ScreenNodeEndpoint   ScreenNodeEndpoint `json:"screenNodeEndpoint"`
	ArrowMetadataVersion int64              `json:"arrowMetadataVersion"`
}

type Footer struct {
	Field []Field `json:"field"`
}
type Field struct {
	MajorType MajorType `json:"majorType"`
	NamePart  NamePart  `json:"namePart"`
	Child     []Child   `json:"child"`
}

type MajorType struct {
	MinorType int64 `json:"minorType"`
	Mode      int64 `json:"mode"`
	Precision int64 `json:"precision"`
}

//	type NamePart struct {
//		Type int64  `json:"type"`
//		Name string `json:"name"`
//	}

type Child struct {
	MajorType MajorType `json:"majorType"`
	NamePart  NamePart  `json:"namePart"`
}

//	type MajorType struct{
//		MinorType int64 `json:"minorType"`
//		Mode int64 `json:"mode"`
//	}

type NamePart struct {
	Type int64  `json:"type"`
	Name string `json:"name"`
}

type ScreenNodeEndpoint struct {
	AvailableCores  int64       `json:"availableCores"`
	NodeTag         string      `json:"nodeTag"`
	ProvisionId     string      `json:"provisionId"`
	UserPort        int64       `json:"userPort"`
	FabricPort      int64       `json:"fabricPort"`
	Roles           Roles       `json:"roles"`
	StartTime       float64     `json:"startTime"`
	MaxDirectMemory float64     `json:"maxDirectMemory"`
	EngineId        EngineId    `json:"engineId"`
	SubEngineId     SubEngineId `json:"subEngineId"`
	Address         string      `json:"address"`
}

type Roles struct {
	DistributedCache bool `json:"distributedCache"`
	Master           bool `json:"master"`
	SqlQuery         bool `json:"sqlQuery"`
	LogicalPlan      bool `json:"logicalPlan"`
	PhysicalPlan     bool `json:"physicalPlan"`
	JavaExecutor     bool `json:"javaExecutor"`
}

type EngineId struct {
	Id string `json:"id"`
}
type SubEngineId struct {
	Id string `json:"id"`
}
type ResourceSchedulingInfo struct {
	ResourceSchedulingEnd   float64 `json:"resourceSchedulingEnd"`
	QueryCost               int64   `json:"queryCost"`
	EngineName              string  `json:"engineName"`
	QueueName               string  `json:"queueName"`
	QueueId                 string  `json:"queueId"`
	ResourceSchedulingStart float64 `json:"resourceSchedulingStart"`
}
type GrandParents struct {
	DatasetPath []string `json:"datasetPath"`
	Level       int64    `json:"level"`
}
type NodeDetails struct {
	FabricPort      int64  `json:"fabricPort"`
	MaxMemoryUsedKb int64  `json:"maxMemoryUsedKb"`
	HostName        string `json:"hostName"`
	HostIp          string `json:"hostIp"`
}
type Stats struct {
	AddedFiles      int64 `json:"addedFiles"`
	RemovedFiles    int64 `json:"removedFiles"`
	InputBytes      int64 `json:"inputBytes"`
	OutputBytes     int64 `json:"outputBytes"`
	InputRecords    int64 `json:"inputRecords"`
	OutputRecords   int64 `json:"outputRecords"`
	IsOutputLimited bool  `json:"isOutputLimited"`
}
type Details struct {
	DataVolume          int64   `json:"dataVolume"`
	OutputRecords       int64   `json:"outputRecords"`
	PeakMemory          float64 `json:"peakMemory"`
	TotalMemory         float64 `json:"totalMemory"`
	CpuUsed             int64   `json:"cpuUsed"`
	TimeSpentInPlanning int64   `json:"timeSpentInPlanning"`
	WaitInClient        int64   `json:"waitInClient"`
}
