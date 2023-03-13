package cmsdto

type PackageData struct {
	Id          int64                `json:"id,omitempty" bson:"id"`
	Type        string               `json:"type,omitempty" bson:"type"`
	Name        string               `json:"name,omitempty" bson:"name"`
	Dimension   Dimensions           `json:"dimension,omitempty" bson:"dimension"`
	Weight      float64              `json:"weight,omitempty" bson:"weight"`
	Cbm         float64              `json:"cbm,omitempty" bson:"cbm"`
	WeightIndex float64              `json:"weightIndex,omitempty" bson:"weightIndex"`
	PackageDef  PackageDefinitionDto `json:"packageDef,omitempty" bson:"packageDef"`
}

type Dimensions struct {
	Length float64 `json:"length,omitempty" bson:"length"`
	Width  float64 `json:"width,omitempty" bson:"width"`
	Height float64 `json:"height,omitempty" bson:"height"`
	Unit   int     `json:"unit,omitempty" bson:"unit"`
}

type PackageDefinitionDto struct {
	Metric    []Matric   `json:"metric,omitempty"`
	Variables []Variable `json:"variables,omitempty"`
	Formula   string     `json:"formula,omitempty"`
}

type Matric struct {
	Ref string `json:"ref,omitempty"`
}

type Variable struct {
	Ref        string       `json:"ref,omitempty"`
	Name       string       `json:"name,omitempty"`
	Label      PackageLabel `json:"label,omitempty"`
	Type       string       `json:"type,omitempty"`
	IsCurrency bool         `json:"isCurrency,omitempty"`
	Value      string       `json:"value,omitempty"`
}

type PackageLabel struct {
	En string `json:"en,omitempty"`
	ID string `json:"id,omitempty"`
}

type DeliveryFeeDto struct {
	Metric []struct {
		Ref      string  `json:"ref,omitempty"`
		Entity   string  `json:"entity,omitempty"`
		MinValue float64 `json:"min-value,omitempty"`
		MaxValue float64 `json:"max-value,omitempty"`
	} `json:"metric,omitempty"`
	Variables []struct {
		Ref   string `json:"ref,omitempty"`
		Name  string `json:"name,omitempty"`
		Label struct {
			En string `json:"en,omitempty"`
			ID string `json:"id,omitempty"`
		} `json:"label,omitempty"`
		Type       string `json:"type,omitempty"`
		IsCurrency bool   `json:"isCurrency,omitempty"`
		Value      string `json:"value,omitempty"`
	} `json:"variables,omitempty"`
	Formula string `json:"formula,omitempty"`
}

type DeliveryStatusDto struct {
	Index       float32   `json:"index,omitempty" bson:"index"`
	Status      string    `json:"status,omitempty" bson:"status"`
	Code        string    `json:"code,omitempty" bson:"code"`
	Description string    `json:"description,omitempty" bson:"description"`
	Attempt     int       `json:"attempt,omitempty" bson:"attempt"`
	PrevIndex   []float32 `json:"prevIndex,omitempty" bson:"prevIndex"`
	NextIndex   []float32 `json:"nextIndex,omitempty" bson:"nextIndex"`
	BatchStatus string    `json:"batchStatus,omitempty" bson:"batchStatus"`
}
