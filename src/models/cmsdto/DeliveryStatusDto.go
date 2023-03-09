package cmsdto

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/lo"
)

func GetDeliveryStatusList() ([]DeliveryStatusDto, error) {
	mydir, _ := os.Getwd()
	utilsDir := filepath.Join(mydir, "/src/utilities")
	log.Println("Exec base dir", utilsDir)

	envFileBytes, err := ioutil.ReadFile(filepath.Join(utilsDir, "delivery_status.json"))
	log.Println("string(envFileBytes): ", string(envFileBytes))
	var deliveryStatusArr []DeliveryStatusDto
	delErr := json.Unmarshal(envFileBytes, &deliveryStatusArr)

	if delErr != nil {
		log.Println("delErr: ", delErr)
		return []DeliveryStatusDto{}, delErr
	}
	if err != nil {
		log.Println("err: ", err)
		return []DeliveryStatusDto{}, err
	}

	return deliveryStatusArr, nil
}

func GetStatus(code string, attempt int) (DeliveryStatusDto, error) {
	statusList, e := GetDeliveryStatusList()
	if e != nil {
		return DeliveryStatusDto{}, e
	}

	log.Println("Current CODE ", code, " Current Attempt ", attempt)
	selectedStatus := lo.Filter(statusList, func(st DeliveryStatusDto, x int) bool {
		return strings.EqualFold(st.Code, code) && st.Attempt == attempt
	})

	log.Println("selectedStatus: ", selectedStatus)

	if len(selectedStatus) == 0 {
		return DeliveryStatusDto{}, errors.New("No status found")
	}

	return selectedStatus[0], nil
}

func GetStatusList(code string) ([]DeliveryStatusDto, error) {
	statusList, e := GetDeliveryStatusList()
	if e != nil {
		return []DeliveryStatusDto{}, e
	}
	selectedStatus := lo.Filter(statusList, func(st DeliveryStatusDto, x int) bool {
		return strings.EqualFold(st.Code, code)
	})

	if len(selectedStatus) == 0 {
		return []DeliveryStatusDto{}, errors.New("No status found")
	}

	return selectedStatus, nil
}
