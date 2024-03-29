package queries

type QueriesJSON struct {
	Username                string        `json:"username"`
	PendingTime             int64         `json:"pendingTime"`
	EngineName              string        `json:"engineName"`
	EngineStart             int64         `json:"engineStart"`
	SetupTimeNs             int64         `json:"setupTimeNs"`
	InputBytes              int64         `json:"inputBytes"`
	PlanningTime            int64         `json:"planningTime"`
	EngineStartTime         int64         `json:"engineStartTime"`
	ExecutionPlanningStart  int64         `json:"executionPlanningStart"`
	WaitTimeNs              int64         `json:"waitTimeNs"`
	QueryId                 string        `json:"queryId"`
	QueryType               string        `json:"queryType"`
	ParentsList             []ParentsList `json:"parentsList"`
	QueueName               string        `json:"queueName"`
	QueuedTime              int64         `json:"queuedTime"`
	MetadataRetrieval       int64         `json:"metadataRetrieval"`
	QueryEnqueued           int64         `json:"queryEnqueued"`
	ScannedDatasets         []interface{} `json:"scannedDatasets"`
	Start                   int64         `json:"start"`
	InputRecords            int64         `json:"inputRecords"`
	QueryCost               float64       `json:"queryCost"`
	MetadataRetrievalTime   int64         `json:"metadataRetrievalTime"`
	StartingTime            int64         `json:"startingTime"`
	Submitted               int64         `json:"submitted"`
	OutcomeReason           string        `json:"outcomeReason"`
	Accelerated             bool          `json:"accelerated"`
	ReflectionRelationships []interface{} `json:"reflectionRelationships"`
	AttemptCount            int64         `json:"attemptCount"`
	ExecutionCpuTimeNs      int64         `json:"executionCpuTimeNs"`
	Context                 string        `json:"context"`
	OutputBytes             int64         `json:"outputBytes"`
	PoolWaitTime            int64         `json:"poolWaitTime"`
	Finish                  int64         `json:"finish"`
	Outcome                 string        `json:"outcome"`
	ExecutionPlanningTime   int64         `json:"executionPlanningTime"`
	RunningTime             int64         `json:"runningTime"`
	PlanningStart           int64         `json:"planningStart"`
	ExecutionStart          int64         `json:"executionStart"`
	IsTruncatedQueryText    bool          `json:"isTruncatedQueryText"`
	MemoryAllocated         int64         `json:"memoryAllocated"`
	QueryText               string        `json:"queryText"`
	OutputRecords           int64         `json:"outputRecords"`
	RequestType             string        `json:"requestType"`
	ExecutionNodes          []interface{} `json:"executionNodes"`
}

// GetMetadataRetrieval allows us to get metadata retrieval on older profiles
func (q *QueriesJSON) GetMetadataRetrieval() int64 {
	epochSince2017 := int64(1500000000000)
	if q.MetadataRetrieval > epochSince2017 {
		return int64(q.MetadataRetrievalTime)
	}
	return q.MetadataRetrieval
}

type ParentsList struct {
	Datasetpath []string `json:"datasetpath"`
	Type        string   `json:"type"`
	Name        string   `json:"name"`
}
