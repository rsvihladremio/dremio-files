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

// Package profiles provides shared parsing of a query profile profile.json file
package profiles

type Profile struct {
	State                     int64                     `json:"state"`
	DatasetProfile            []DatasetProfile          `json:"datasetProfile"`
	Id                        Id                        `json:"id"`
	TotalFragments            int64                     `json:"totalFragments"`
	PlanningStart             float64                   `json:"planningStart"`
	AccelerationProfile       AccelerationProfile       `json:"accelerationProfile"`
	Start                     float64                   `json:"start"`
	NodeProfile               []NodeProfile             `json:"nodeProfile"`
	SerializedPlan            string                    `json:"serializedPlan"`
	NumPlanCacheUsed          int64                     `json:"numPlanCacheUsed"`
	NumJoinsInFinalPrel       int64                     `json:"numJoinsInFinalPrel"`
	End                       float64                   `json:"end"`
	FinishedFragments         int64                     `json:"finishedFragments"`
	ClientInfo                ClientInfo                `json:"clientInfo"`
	PlanPhases                []PlanPhases              `json:"planPhases"`
	CommandPoolWaitMillis     int64                     `json:"commandPoolWaitMillis"`
	Query                     string                    `json:"query"`
	Foreman                   Foreman                   `json:"foreman"`
	Plan                      string                    `json:"plan"`
	ResourceSchedulingProfile ResourceSchedulingProfile `json:"resourceSchedulingProfile"`
	RelInfoMap                RelInfoMap                `json:"relInfoMap"`
	NonDefaultOptionsJSON     string                    `json:"nonDefaultOptionsJSON"`
	DremioVersion             string                    `json:"dremioVersion"`
	OperatorTypeMetricsMap    OperatorTypeMetricsMap    `json:"operatorTypeMetricsMap"`
	StateList                 []StateList               `json:"stateList"`
	FragmentProfile           []FragmentProfile         `json:"fragmentProfile"`
	User                      string                    `json:"user"`
	PlanningEnd               float64                   `json:"planningEnd"`
	JsonPlan                  string                    `json:"jsonPlan"`
	NumJoinsInUserQuery       int64                     `json:"numJoinsInUserQuery"`
}

type DatasetProfile struct {
	Type        int64  `json:"type"`
	Sql         string `json:"sql"`
	DatasetPath string `json:"datasetPath"`
}

type Id struct {
	Part2 float64 `json:"part2"`
	Part1 float64 `json:"part1"`
}

type AccelerationProfile struct {
	MillisTakenGettingMaterializations int64         `json:"millisTakenGettingMaterializations"`
	MillisTakenNormalizing             int64         `json:"millisTakenNormalizing"`
	MillisTakenSubstituting            int64         `json:"millisTakenSubstituting"`
	LayoutProfiles                     []interface{} `json:"layoutProfiles"`
	NormalizedQueryPlans               []interface{} `json:"normalizedQueryPlans"`
	AccelerationDetails                string        `json:"accelerationDetails"`
	Accelerated                        bool          `json:"accelerated"`
	NumSubstitutions                   int64         `json:"numSubstitutions"`
}

type NodeProfile struct {
	Endpoint                   Endpoint `json:"endpoint"`
	MaxMemoryUsed              float64  `json:"maxMemoryUsed"`
	TimeEnqueuedBeforeSubmitMs int64    `json:"timeEnqueuedBeforeSubmitMs"`
	NumberOfCores              int64    `json:"numberOfCores"`
}

type ClientInfo struct {
	MinorVersion     int64  `json:"minorVersion"`
	PatchVersion     int64  `json:"patchVersion"`
	Application      string `json:"application"`
	BuildNumber      int64  `json:"buildNumber"`
	VersionQualifier string `json:"versionQualifier"`
	Name             string `json:"name"`
	Version          string `json:"version"`
	MajorVersion     int64  `json:"majorVersion"`
}

type PlanPhases struct {
	PhaseName            string               `json:"phaseName"`
	DurationMillis       int64                `json:"durationMillis"`
	Plan                 string               `json:"plan"`
	TimeBreakdownPerRule TimeBreakdownPerRule `json:"timeBreakdownPerRule"`
}

type TimeBreakdownPerRule struct {
}

type Foreman struct {
	Roles           Roles   `json:"roles"`
	StartTime       float64 `json:"startTime"`
	MaxDirectMemory float64 `json:"maxDirectMemory"`
	ConduitPort     int64   `json:"conduitPort"`
	DremioVersion   string  `json:"dremioVersion"`
	FabricPort      int64   `json:"fabricPort"`
	UserPort        int64   `json:"userPort"`
	AvailableCores  int64   `json:"availableCores"`
	NodeTag         string  `json:"nodeTag"`
	Address         string  `json:"address"`
}

type Roles struct {
	JavaExecutor bool `json:"javaExecutor"`
	Master       bool `json:"master"`
	SqlQuery     bool `json:"sqlQuery"`
}

type ResourceSchedulingProfile struct {
	ResourceSchedulingEnd   float64              `json:"resourceSchedulingEnd"`
	QueueName               string               `json:"queueName"`
	QueueId                 string               `json:"queueId"`
	RuleContent             string               `json:"ruleContent"`
	RuleName                string               `json:"ruleName"`
	RuleAction              string               `json:"ruleAction"`
	SchedulingProperties    SchedulingProperties `json:"schedulingProperties"`
	ResourceSchedulingStart float64              `json:"resourceSchedulingStart"`
}

type SchedulingProperties struct {
	QueryType  string `json:"queryType"`
	QueryLabel string `json:"queryLabel"`
	QueryCost  int64  `json:"queryCost"`
}

type RelInfoMap struct {
}

type OperatorTypeMetricsMap struct {
	MetricsDef []MetricsDef `json:"metricsDef"`
}

type MetricsDef struct {
	MetricDef []MetricDef `json:"metricDef"`
}

type MetricDef struct {
	DisplayCode     string `json:"displayCode"`
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	DisplayType     int64  `json:"displayType"`
	AggregationType int64  `json:"aggregationType"`
}

type StateList struct {
	State     int64   `json:"state"`
	StartTime float64 `json:"startTime"`
}

type FragmentProfile struct {
	MajorFragmentId      int64                  `json:"majorFragmentId"`
	MinorFragmentProfile []MinorFragmentProfile `json:"minorFragmentProfile"`
	NodePhaseProfile     []NodePhaseProfile     `json:"nodePhaseProfile"`
	PhaseWeight          int64                  `json:"phaseWeight"`
}

type MinorFragmentProfile struct {
	MaxMemoryUsed                   float64                      `json:"maxMemoryUsed"`
	BlockedOnDownstreamDuration     int64                        `json:"blockedOnDownstreamDuration"`
	NumShortSlices                  int64                        `json:"numShortSlices"`
	LastProgress                    float64                      `json:"lastProgress"`
	RunQLoad                        int64                        `json:"runQLoad"`
	StartTime                       float64                      `json:"startTime"`
	MaxIncomingMemoryUsed           int64                        `json:"maxIncomingMemoryUsed"`
	SleepingDuration                int64                        `json:"sleepingDuration"`
	NumSlices                       int64                        `json:"numSlices"`
	FirstRun                        float64                      `json:"firstRun"`
	NumLongSlices                   int64                        `json:"numLongSlices"`
	State                           int64                        `json:"state"`
	MinorFragmentId                 int64                        `json:"minorFragmentId"`
	Endpoint                        Endpoint                     `json:"endpoint"`
	NumRuns                         int64                        `json:"numRuns"`
	BlockedOnMemoryDuration         int64                        `json:"blockedOnMemoryDuration"`
	OperatorProfile                 []OperatorProfile            `json:"operatorProfile"`
	EndTime                         float64                      `json:"endTime"`
	FinishDuration                  int64                        `json:"finishDuration"`
	BlockedOnUpstreamDuration       int64                        `json:"blockedOnUpstreamDuration"`
	PerResourceBlockedDuration      []PerResourceBlockedDuration `json:"perResourceBlockedDuration"`
	LastUpdate                      float64                      `json:"lastUpdate"`
	BlockedDuration                 int64                        `json:"blockedDuration"`
	SetupDuration                   int64                        `json:"setupDuration"`
	BlockedOnSharedResourceDuration int64                        `json:"blockedOnSharedResourceDuration"`
	MemoryUsed                      int64                        `json:"memoryUsed"`
	RunDuration                     int64                        `json:"runDuration"`
}

type OperatorProfile struct {
	ProcessNanos             int64          `json:"processNanos"`
	PeakLocalMemoryAllocated int64          `json:"peakLocalMemoryAllocated"`
	Metric                   []Metric       `json:"metric"`
	OperatorSubtype          int64          `json:"operatorSubtype"`
	OutputBytes              int64          `json:"outputBytes"`
	InputProfile             []InputProfile `json:"inputProfile"`
	OperatorType             int64          `json:"operatorType"`
	SetupNanos               int64          `json:"setupNanos"`
	OutputRecords            int64          `json:"outputRecords"`
	RemovedFiles             int64          `json:"removedFiles"`
	OperatorId               int64          `json:"operatorId"`
	WaitNanos                int64          `json:"waitNanos"`
	AddedFiles               int64          `json:"addedFiles"`
}

type Metric struct {
	LongValue int64 `json:"longValue"`
	MetricId  int64 `json:"metricId"`
}

type InputProfile struct {
	Records int64 `json:"records"`
	Batches int64 `json:"batches"`
	Size    int64 `json:"size"`
}

type PerResourceBlockedDuration struct {
	Resource string `json:"resource"`
	Category int64  `json:"category"`
	Duration int64  `json:"duration"`
}

type NodePhaseProfile struct {
	Endpoint      Endpoint `json:"endpoint"`
	MaxMemoryUsed float64  `json:"maxMemoryUsed"`
}

type Endpoint struct {
	Address    string `json:"address"`
	FabricPort int64  `json:"fabricPort"`
}
