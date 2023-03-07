package cmsdto

type PackageData struct {
	Id          int64                `json:"id" bson:"id"`
	Type        string               `json:"type" bson:"type"`
	Name        string               `json:"name" bson:"name"`
	Dimension   Dimensions           `json:"dimension" bson:"dimension"`
	Weight      float64              `json:"weight" bson:"weight"`
	Cbm         float64              `json:"cbm" bson:"cbm"`
	WeightIndex float64              `json:"weightIndex" bson:"weightIndex"`
	PackageDef  PackageDefinitionDto `json:"packageDef" bson:"packageDef"`
}

type Dimensions struct {
	Length float64 `json:"length" bson:"length"`
	Width  float64 `json:"width" bson:"width"`
	Height float64 `json:"height" bson:"height"`
	Unit   int     `json:"unit" bson:"unit"`
}

type PackageDefinitionDto struct {
	Metric    []Matric   `json:"metric"`
	Variables []Variable `json:"variables"`
	Formula   string     `json:"formula"`
}

type Matric struct {
	Ref string `json:"ref"`
}

type Variable struct {
	Ref        string       `json:"ref"`
	Name       string       `json:"name"`
	Label      PackageLabel `json:"label"`
	Type       string       `json:"type"`
	IsCurrency bool         `json:"isCurrency"`
	Value      string       `json:"value"`
}

type PackageLabel struct {
	En string `json:"en"`
	ID string `json:"id"`
}

type DeliveryFeeDto struct {
	Metric []struct {
		Ref      string  `json:"ref"`
		Entity   string  `json:"entity"`
		MinValue float64 `json:"min-value"`
		MaxValue float64 `json:"max-value"`
	} `json:"metric"`
	Variables []struct {
		Ref   string `json:"ref"`
		Name  string `json:"name"`
		Label struct {
			En string `json:"en"`
			ID string `json:"id"`
		} `json:"label"`
		Type       string `json:"type"`
		IsCurrency bool   `json:"isCurrency"`
		Value      string `json:"value"`
	} `json:"variables"`
	Formula string `json:"formula"`
}

type DeliveryStatusDto struct {
	Index       float32   `json:"index" bson:"index"`
	Status      string    `json:"status" bson:"status"`
	Code        string    `json:"code" bson:"code"`
	Description string    `json:"description" bson:"description"`
	Attempt     int       `json:"attempt" bson:"attempt"`
	PrevIndex   []float32 `json:"prevIndex" bson:"prevIndex"`
	NextIndex   []float32 `json:"nextIndex" bson:"nextIndex"`
	BatchStatus string    `json:"batchStatus" bson:"batchStatus"`
}
