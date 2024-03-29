package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"sort"
	"strings"
)

type ElasticSearchRest struct {
}

func (es *ElasticSearchRest) GetBoundary(d float64, centerLat float64, centerLng float64, bearingInDegrees float64) []float64 {
	lat1 := centerLat * (math.Pi / 180) //Only do this if you need to convert from deg. to dec.
	lon1 := centerLng * (math.Pi / 180) //Only do this if you need to convert from deg. to dec.
	var R float64 = 6371000
	brng := bearingInDegrees * (math.Pi / 180)

	lat2 := math.Asin(math.Sin(lat1)*math.Cos(d/R) + math.Cos(lat1)*math.Sin(d/R)*math.Cos(brng))
	lon2 := lon1 + math.Atan2(math.Sin(brng)*math.Sin(d/R)*math.Cos(lat1), math.Cos(d/R)-math.Sin(lat1)*math.Sin(lat2))
	return []float64{lat2 * (180 / math.Pi), lon2 * (180 / math.Pi)}

}

func (es *ElasticSearchRest) GetLatLng(clat float64, clng float64, dist float64, bearingDeg float64) []float64 {

	var rad float64 = 6371000
	brng := bearingDeg * (math.Pi / 180)
	newLat := math.Asin(math.Sin(clat)*math.Cos(dist/rad) + math.Cos(clat)*math.Sin(dist/rad)*math.Cos(brng))
	newLng := clng + math.Atan2(math.Sin(brng)*math.Sin(dist/rad)*math.Cos(clat), math.Cos(dist/rad)-math.Sin(clat)*math.Sin(newLat))

	return []float64{newLat * (180 / math.Pi), newLng * (180 / math.Pi)}
}

func (es *ElasticSearchRest) SquareBoundary(centroid []float64, semiDiagonal float64) [][]float64 {
	var vertices [][]float64

	log.Println("centroid: ", centroid)
	log.Println("semiDiagonal in Meter: ", semiDiagonal)

	cLat := centroid[0] * (math.Pi / 180)
	cLng := centroid[1] * (math.Pi / 180)

	top_left := es.GetLatLng(cLat, cLng, semiDiagonal, 135)
	top_right := es.GetLatLng(cLat, cLng, semiDiagonal, 45)
	bottom_right := es.GetLatLng(cLat, cLng, semiDiagonal, 315)
	bottom_left := es.GetLatLng(cLat, cLng, semiDiagonal, 225)

	vertices = append(vertices, top_left, top_right, bottom_right, bottom_left)
	log.Println("vertices: ", vertices)
	return vertices
}

func calculateHexagonCoordinates(centroid []float64, gridSize float64) [][]float64 {
	coordinates := make([][]float64, 6)

	for i := 0; i < 6; i++ {
		angle := math.Pi / 3 * float64(i)
		x := centroid[0] + gridSize*math.Cos(angle)
		y := centroid[1] + gridSize*math.Sin(angle)
		coordinates[i] = []float64{x, y}
	}

	// find and print the coordinates of the vertices of the hexagonal area
	vertices := HexagonVertices(centroid[0], centroid[1], gridSize)
	fmt.Println("The coordinates of the vertices of the hexagonal area are:")
	for i, v := range vertices {
		fmt.Printf("Vertex %d: (%f, %f)\n", i+1, v[0], v[1])
		coordinates[i] = []float64{v[0], v[1]}
	}

	return coordinates
}

func (es *ElasticSearchRest) HexagonalBoundary(centroid []float64, gridSize float64) [][]float64 {
	var vertices [][]float64

	log.Println("centroid: ", centroid)
	log.Println("gridSize in Meter: ", gridSize)

	hexagonCoordinates := calculateHexagonCoordinates(centroid, gridSize)

	for i, point := range hexagonCoordinates {
		fmt.Printf("Vertex %d: X: %f, Y: %f\n", i+1, point[0], point[1])
		vertices = append(vertices, []float64{point[0], point[1]})
	}

	vertices = append(vertices, []float64{hexagonCoordinates[0][0], hexagonCoordinates[0][1]})

	log.Println("vertices: ", vertices)
	return vertices
}

func toRadians(deg float64) float64 {
	return deg * math.Pi / 180
}

func (es *ElasticSearchRest) Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371 // in kilometers

	// convert to radians
	lat1, lon1, lat2, lon2 = toRadians(lat1), toRadians(lon1), toRadians(lat2), toRadians(lon2)

	// calculate haversine
	dLat, dLon := lat2-lat1, lon2-lon1
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadius * c

	return distance
}

func (es *ElasticSearchRest) IndexDoc(data string, id string) (error, bool) {

	url := os.Getenv("ELASTIC_URL") + "/" + os.Getenv("LOCATION_INDEX") + "/_doc/" + id
	method := "POST"

	payload := strings.NewReader(data)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err, false
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err, false
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err, false
	}
	fmt.Println(string(body))

	return nil, true
}

type DriverSearchQuery struct {
	Query struct {
		Bool struct {
			Must   []map[string]interface{} `json:"must"`
			Filter struct {
				Shape struct {
					Location struct {
						Shape struct {
							Type        string      `json:"type"`
							Coordinates [][]float64 `json:"coordinates"`
						} `json:"shape"`
						Relation string `json:"relation"`
					} `json:"location"`
				} `json:"shape"`
			} `json:"filter"`
		} `json:"bool"`
	} `json:"query"`
}
type DriverSearchHexagonQuery struct {
	Query struct {
		Bool struct {
			Must   []map[string]interface{} `json:"must"`
			Filter struct {
				Shape struct {
					Location struct {
						Shape struct {
							Type        string        `json:"type"`
							Coordinates [][][]float64 `json:"coordinates"`
						} `json:"shape"`
						Relation string `json:"relation"`
					} `json:"location"`
				} `json:"shape"`
			} `json:"filter"`
		} `json:"bool"`
	} `json:"query"`
}
type DriverLocationDto struct {
	DriverID    string `json:"driverId"`
	City        string `json:"city"`
	BshtTag     string `json:"bshtTag"`
	LocationObj struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"location"`
	OnlineOfflineStatus string `json:"onlineOfflineStatus"`
	Status              string `json:"status"`
	Timestamp           int64  `json:"timestamp" omitempty:"true"`
}

func (es *ElasticSearchRest) SortByDistance(drivers []DriverLocationDto, targetLocation []float64) []DriverLocationDto {

	sort.Slice(drivers, func(i, j int) bool {

		distI := es.Haversine(drivers[i].LocationObj.Coordinates[0], drivers[i].LocationObj.Coordinates[1], targetLocation[0], targetLocation[1])
		distJ := es.Haversine(drivers[j].LocationObj.Coordinates[0], drivers[j].LocationObj.Coordinates[1], targetLocation[0], targetLocation[1])
		return distI < distJ
	})

	return drivers
}

func (es *ElasticSearchRest) SearchDriversWithInBoundary(topLeft []float64, bottomRight []float64, onlineOffline string, status string, city string, timeRange string, driverId string) ([]DriverLocationDto, error, bool) {
	searchQuery := DriverSearchQuery{}
	dlArr := []DriverLocationDto{}

	if onlineOffline != "" {

		match1 := make(map[string]interface{})
		m1 := make(map[string]interface{})
		m1["onlineOfflineStatus"] = onlineOffline
		match1["match"] = m1
		searchQuery.Query.Bool.Must = append(searchQuery.Query.Bool.Must, match1)

	}

	if status != "" {

		match2 := make(map[string]interface{})
		m2 := make(map[string]interface{})
		m2["status"] = "available"
		match2["match"] = m2

		searchQuery.Query.Bool.Must = append(searchQuery.Query.Bool.Must, match2)

	}

	if city != "" {
		match3 := make(map[string]interface{})
		m3 := make(map[string]interface{})
		m3["city"] = city
		match3["match"] = m3

		searchQuery.Query.Bool.Must = append(searchQuery.Query.Bool.Must, match3)

	}

	if driverId != "" {
		match5 := make(map[string]interface{})
		m5 := make(map[string]interface{})
		m5["driverId"] = driverId
		match5["match"] = m5

		searchQuery.Query.Bool.Must = append(searchQuery.Query.Bool.Must, match5)

	}

	match4 := make(map[string]interface{})
	m4 := make(map[string]interface{})
	m44 := make(map[string]interface{})
	m44["gte"] = timeRange
	m44["lte"] = "now"
	m4["timestamp"] = m44
	match4["range"] = m4

	searchQuery.Query.Bool.Must = append(searchQuery.Query.Bool.Must, match4)

	log.Println(searchQuery)

	searchQuery.Query.Bool.Filter.Shape.Location.Relation = "within"
	searchQuery.Query.Bool.Filter.Shape.Location.Shape.Type = "envelope"
	searchQuery.Query.Bool.Filter.Shape.Location.Shape.Coordinates = [][]float64{topLeft, bottomRight}

	queryByte, queryErr := json.Marshal(&searchQuery)
	log.Println(string(queryByte))
	log.Println("queryErr: ", queryErr)

	url := os.Getenv("ELASTIC_URL") + "/" + os.Getenv("LOCATION_INDEX") + "/_search"
	method := "GET"

	payload := strings.NewReader(string(queryByte))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return dlArr, err, false
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return dlArr, err, false
	}
	defer res.Body.Close()

	r := make(map[string]interface{})

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return dlArr, err, false
	}
	fmt.Println(string(body))
	ee := json.Unmarshal(body, &r)

	if ee != nil {
		return dlArr, ee, false
	}

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		dl := DriverLocationDto{}
		a := hit.(map[string]interface{})["_source"]
		bb, _ := json.Marshal(&a)
		log.Println(string(bb))
		json.Unmarshal(bb, &dl)
		dlArr = append(dlArr, dl)
	}

	return dlArr, nil, true
}

func (es *ElasticSearchRest) SearchDriversWithInHexagonalBoundary(boundary [][]float64, onlineOffline string, status string, city string, timeRange string, driverId string, targetPoint []float64) ([]DriverLocationDto, error, bool) {
	searchQuery := DriverSearchHexagonQuery{}
	dlArr := []DriverLocationDto{}

	if onlineOffline != "" {

		match1 := make(map[string]interface{})
		m1 := make(map[string]interface{})
		m1["onlineOfflineStatus"] = onlineOffline
		match1["match"] = m1
		searchQuery.Query.Bool.Must = append(searchQuery.Query.Bool.Must, match1)

	}

	if status != "" {

		match2 := make(map[string]interface{})
		m2 := make(map[string]interface{})
		m2["status"] = "available"
		match2["match"] = m2

		searchQuery.Query.Bool.Must = append(searchQuery.Query.Bool.Must, match2)

	}

	if city != "" {
		match3 := make(map[string]interface{})
		m3 := make(map[string]interface{})
		m3["city"] = city
		match3["match"] = m3

		searchQuery.Query.Bool.Must = append(searchQuery.Query.Bool.Must, match3)

	}

	if driverId != "" {
		match5 := make(map[string]interface{})
		m5 := make(map[string]interface{})
		m5["driverId"] = driverId
		match5["match"] = m5

		searchQuery.Query.Bool.Must = append(searchQuery.Query.Bool.Must, match5)

	}

	match4 := make(map[string]interface{})
	m4 := make(map[string]interface{})
	m44 := make(map[string]interface{})
	m44["gte"] = timeRange
	m44["lte"] = "now"
	m4["timestamp"] = m44
	match4["range"] = m4

	searchQuery.Query.Bool.Must = append(searchQuery.Query.Bool.Must, match4)

	log.Println(searchQuery)

	searchQuery.Query.Bool.Filter.Shape.Location.Relation = "within"
	searchQuery.Query.Bool.Filter.Shape.Location.Shape.Type = "polygon"
	coord := [][][]float64{}
	coord = append(coord, boundary)
	searchQuery.Query.Bool.Filter.Shape.Location.Shape.Coordinates = coord

	queryByte, queryErr := json.Marshal(&searchQuery)
	log.Println(string(queryByte))
	log.Println("queryErr: ", queryErr)

	url := os.Getenv("ELASTIC_URL") + "/" + os.Getenv("LOCATION_INDEX") + "/_search"
	method := "GET"

	payload := strings.NewReader(string(queryByte))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return dlArr, err, false
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return dlArr, err, false
	}
	defer res.Body.Close()

	r := make(map[string]interface{})

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return dlArr, err, false
	}
	fmt.Println(string(body))
	ee := json.Unmarshal(body, &r)

	if ee != nil {
		return dlArr, ee, false
	}

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		dl := DriverLocationDto{}
		a := hit.(map[string]interface{})["_source"]
		bb, _ := json.Marshal(&a)
		log.Println(string(bb))
		json.Unmarshal(bb, &dl)
		dlArr = append(dlArr, dl)
	}

	if len(dlArr) > 1 {
		dlArr = es.SortByDistance(dlArr, targetPoint)
	}

	return dlArr, nil, true
}

type DriverShortSearchQuery struct {
	Query struct {
		Bool struct {
			Must []map[string]interface{} `json:"must"`
		} `json:"bool"`
	} `json:"query"`
}

type LocationData struct {
	DriverRequest struct {
		Id           string `json:"id"`
		MobileNumber string `json:"mobileNumber"`
		Name         string `json:"name"`
	} `json:"drivers"`
	LocationObj struct {
		Coordinates []float64 `json:"coordinates"`
	} `json:"location"`
}

type DriverLocationModel struct {
	DriverID    string `json:"driverId"`
	City        string `json:"city"`
	BshtTag     string `json:"bshtTag"`
	LocationObj struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"location"`
	OnlineOfflineStatus string `json:"onlineOfflineStatus"`
	Status              string `json:"status"`
	Timestamp           int64  `json:"timestamp" omitempty:"true"`
}

func (es *ElasticSearchRest) SearchDriversWithoutBoundary(onlineOffline string, status string, city string, timeRange string) ([]LocationData, error, bool) {
	searchQuery := DriverShortSearchQuery{}
	dlArr := []LocationData{}

	if onlineOffline != "" {

		match1 := make(map[string]interface{})
		m1 := make(map[string]interface{})
		m1["onlineOfflineStatus"] = onlineOffline
		match1["match"] = m1
		searchQuery.Query.Bool.Must = append(searchQuery.Query.Bool.Must, match1)

	}

	if status != "" {

		match2 := make(map[string]interface{})
		m2 := make(map[string]interface{})
		m2["status"] = status
		match2["match"] = m2

		searchQuery.Query.Bool.Must = append(searchQuery.Query.Bool.Must, match2)

	}

	if city != "" {
		match3 := make(map[string]interface{})
		m3 := make(map[string]interface{})
		m3["city"] = city
		match3["match"] = m3

		searchQuery.Query.Bool.Must = append(searchQuery.Query.Bool.Must, match3)

	}

	match4 := make(map[string]interface{})
	m4 := make(map[string]interface{})
	m44 := make(map[string]interface{})
	m44["gte"] = "now-" + timeRange

	m44["lte"] = "now"
	m4["timestamp"] = m44
	match4["range"] = m4

	searchQuery.Query.Bool.Must = append(searchQuery.Query.Bool.Must, match4)

	log.Println(searchQuery)

	queryByte, queryErr := json.Marshal(&searchQuery)
	log.Println(string(queryByte))
	log.Println("queryErr: ", queryErr)

	url := os.Getenv("ELASTIC_URL") + "/" + os.Getenv("LOCATION_INDEX") + "/_search"
	method := "GET"

	payload := strings.NewReader(string(queryByte))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return dlArr, err, false
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return dlArr, err, false
	}
	defer res.Body.Close()

	r := make(map[string]interface{})

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return dlArr, err, false
	}
	fmt.Println(string(body))
	eee := json.Unmarshal(body, &r)

	if eee != nil {
		return dlArr, eee, false
	}
	// // Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		dl := LocationData{}
		dlm := DriverLocationModel{}
		a := hit.(map[string]interface{})["_source"]
		bb, _ := json.Marshal(&a)
		//log.Println(string(bb))
		json.Unmarshal(bb, &dlm)

		dl.DriverRequest.Id = dlm.DriverID
		//driverDto, _ := FindDriverById(dlm.DriverID)
		//dl.DriverRequest.MobileNumber = driverDto.DriverDetails.Mobileno
		//dl.DriverRequest.Name = fmt.Sprint(driverDto.DriverDetails.Firstname, " ", driverDto.DriverDetails.Lastname)
		dl.LocationObj.Coordinates = dlm.LocationObj.Coordinates

		dlArr = append(dlArr, dl)
	}

	return dlArr, nil, true
}

type DriverLocationRequest struct {
	DriverID string `json:"driverId"`
	City     int64  `json:"city"`
	BshtTag  string `json:"bshtTag"`
	Location struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"location"`
	OnlineOfflineStatus string `json:"onlineOfflineStatus"`
	Status              string `json:"status"`
	Timestamp           int64  `json:"timestamp"`
}

func (es *ElasticSearchRest) GetDriversFromElastic(driverReqObj DriverLocationRequest) ([]DriverLocationDto, error, bool) {
	searchQuery := DriverShortSearchQuery{}
	dlArr := []DriverLocationDto{}
	r := make(map[string]interface{})

	if driverReqObj.OnlineOfflineStatus != "" {

		match1 := make(map[string]interface{})
		m1 := make(map[string]interface{})
		m1["onlineOfflineStatus"] = driverReqObj.OnlineOfflineStatus
		match1["match"] = m1
		searchQuery.Query.Bool.Must = append(searchQuery.Query.Bool.Must, match1)

	}

	if driverReqObj.Status != "" {

		match2 := make(map[string]interface{})
		m2 := make(map[string]interface{})
		m2["status"] = driverReqObj.Status
		match2["match"] = m2

		searchQuery.Query.Bool.Must = append(searchQuery.Query.Bool.Must, match2)

	}

	if driverReqObj.City > 0 {
		match3 := make(map[string]interface{})
		m3 := make(map[string]interface{})
		m3["city"] = fmt.Sprint(driverReqObj.City)
		match3["match"] = m3

		searchQuery.Query.Bool.Must = append(searchQuery.Query.Bool.Must, match3)

	}

	if driverReqObj.DriverID != "" {
		match5 := make(map[string]interface{})
		m5 := make(map[string]interface{})
		m5["driverId"] = driverReqObj.DriverID
		match5["match"] = m5

		searchQuery.Query.Bool.Must = append(searchQuery.Query.Bool.Must, match5)

	}

	match4 := make(map[string]interface{})
	m4 := make(map[string]interface{})
	m44 := make(map[string]interface{})
	m44["gte"] = "now-5m"
	m44["lte"] = "now"
	m4["timestamp"] = m44
	match4["range"] = m4

	searchQuery.Query.Bool.Must = append(searchQuery.Query.Bool.Must, match4)

	log.Println(searchQuery)

	queryByte, queryErr := json.Marshal(&searchQuery)
	log.Println(string(queryByte))
	log.Println("queryErr: ", queryErr)

	url := os.Getenv("ELASTIC_URL") + "/" + os.Getenv("LOCATION_INDEX") + "/_search"
	method := "GET"

	payload := strings.NewReader(string(queryByte))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return dlArr, err, false
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return dlArr, err, false
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return dlArr, err, false
	}
	fmt.Println(string(body))
	ee := json.Unmarshal(body, &r)

	if ee != nil {
		return dlArr, ee, false
	}

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		dl := DriverLocationDto{}
		a := hit.(map[string]interface{})["_source"]
		bb, _ := json.Marshal(&a)
		log.Println(string(bb))
		json.Unmarshal(bb, &dl)
		dlArr = append(dlArr, dl)
	}

	return dlArr, nil, true
}

func (es *ElasticSearchRest) DeleteDoc(driverId string) error {

	url := os.Getenv("ELASTIC_URL") + "/" + os.Getenv("LOCATION_INDEX") + "/_delete_by_query"
	method := "POST"

	payload := strings.NewReader(`{
    "query": {
        "match": {
            "driverId": "` + driverId + `"
        }
    }
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))
	return nil

}

// ToRadians converts degrees to radians
func ToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// ToDegrees converts radians to degrees
func ToDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}

// HaversineDistance returns the great circle distance between two points
// on a sphere using the Haversine formula
func HaversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert degrees to radians
	lat1 = ToRadians(lat1)
	lon1 = ToRadians(lon1)
	lat2 = ToRadians(lat2)
	lon2 = ToRadians(lon2)

	// calculate the angular distance between the points
	dlat := lat2 - lat1
	dlon := lon2 - lon1
	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dlon/2)*math.Sin(dlon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// assume the earth's radius is 6371 km
	r := 6371000.0

	// return the distance in meters
	return c * r
}

// VincentyDistance returns the latitude and longitude of a point given the latitude and longitude of another point, the distance between them, and the initial bearing angle
func VincentyDistance(lat1, lon1, distance, bearing float64) (float64, float64) {
	// convert degrees to radians
	lat1 = ToRadians(lat1)
	lon1 = ToRadians(lon1)
	bearing = ToRadians(bearing)

	// assume the earth's radius is 6371 km
	r := 6371000.0

	// calculate the angular distance
	angularDistance := distance / r

	// calculate the latitude of the second point
	lat2 := math.Asin(math.Sin(lat1)*math.Cos(angularDistance) + math.Cos(lat1)*math.Sin(angularDistance)*math.Cos(bearing))

	// calculate the longitude of the second point
	lon2 := lon1 + math.Atan2(math.Sin(bearing)*math.Sin(angularDistance)*math.Cos(lat1), math.Cos(angularDistance)-math.Sin(lat1)*math.Sin(lat2))

	// convert radians to degrees
	lat2 = ToDegrees(lat2)
	lon2 = ToDegrees(lon2)

	// return the latitude and longitude of the second point
	return lat2, lon2
}

// HexagonVertices returns the latitude and longitude of the vertices of a regular hexagonal area on the earth
func HexagonVertices(lat, lon, distance float64) [][2]float64 {
	// initialize an empty slice of coordinates
	vertices := [][2]float64{}

	// assume that the first vertex is at the same longitude as the centroid
	// and calculate its latitude using the Haversine formula
	lat1 := lat + ToDegrees(distance/6371000)

	// append the first vertex to the slice
	vertices = append(vertices, [2]float64{lat1, lon})

	// calculate the angle between each vertex and the centroid, which is 60 degrees for a regular hexagon
	angle := 60.0

	// iterate over the remaining vertices and calculate their latitude and longitude using the Vincenty formula
	for i := 1; i < 6; i++ {
		lat2, lon2 := VincentyDistance(lat, lon, distance, angle*float64(i))
		vertices = append(vertices, [2]float64{lat2, lon2})
	}

	// return the slice of coordinates
	return vertices
}
