package queries

type JsonObject struct {
	Username                string        `json:"username"`
	PendingTime             int64         `json:"pendingTime"`
	EngineName              string        `json:"engineName"`
	EngineStart             int64         `json:"engineStart"`
	SetupTimeNs             int64         `json:"setupTimeNs"`
	InputBytes              int64         `json:"inputBytes"`
	PlanningTime            int64         `json:"planningTime"`
	EngineStartTime         float64       `json:"engineStartTime"`
	ExecutionPlanningStart  float64       `json:"executionPlanningStart"`
	WaitTimeNs              int64         `json:"waitTimeNs"`
	QueryId                 string        `json:"queryId"`
	QueryType               string        `json:"queryType"`
	ParentsList             []ParentsList `json:"parentsList"`
	QueueName               string        `json:"queueName"`
	QueuedTime              int64         `json:"queuedTime"`
	MetadataRetrieval       int64         `json:"metadataRetrieval"`
	QueryEnqueued           float64       `json:"queryEnqueued"`
	ScannedDatasets         []interface{} `json:"scannedDatasets"`
	Start                   float64       `json:"start"`
	InputRecords            int64         `json:"inputRecords"`
	QueryCost               float64       `json:"queryCost"`
	MetadataRetrievalTime   float64       `json:"metadataRetrievalTime"`
	StartingTime            int64         `json:"startingTime"`
	Submitted               float64       `json:"submitted"`
	OutcomeReason           string        `json:"outcomeReason"`
	Accelerated             bool          `json:"accelerated"`
	ReflectionRelationships []interface{} `json:"reflectionRelationships"`
	AttemptCount            int64         `json:"attemptCount"`
	ExecutionCpuTimeNs      int64         `json:"executionCpuTimeNs"`
	Context                 string        `json:"context"`
	OutputBytes             int64         `json:"outputBytes"`
	PoolWaitTime            int64         `json:"poolWaitTime"`
	Finish                  float64       `json:"finish"`
	Outcome                 string        `json:"outcome"`
	ExecutionPlanningTime   int64         `json:"executionPlanningTime"`
	RunningTime             int64         `json:"runningTime"`
	PlanningStart           float64       `json:"planningStart"`
	ExecutionStart          float64       `json:"executionStart"`
	MemoryAllocated         int64         `json:"memoryAllocated"`
	QueryText               string        `json:"queryText"`
	OutputRecords           int64         `json:"outputRecords"`
	RequestType             string        `json:"requestType"`
	ExecutionNodes          []interface{} `json:"executionNodes"`
}

type ParentsList struct {
	Datasetpath []string `json:"datasetpath"`
	Type        string   `json:"type"`
	Name        string   `json:"name"`
}
