package models

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/RaRa-Delivery/rara-ms-models/src/models/order"
	"github.com/RaRa-Delivery/rara-ms-models/src/utility/lg"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OnlineScreenResponse struct {
	TodaysEarning      string                    `json:"earningType" bson:"earningType"`
	Earning            Earning                   `json:"earning" bson:"earning"`
	Orders             []order.BatchForDriverApp `json:"nodes" bson:"nodes"`
	Kecamatan          []string                  `json:"kecamatan" bson:"kecamatan"`
	IsBatchAssigned    bool                      `json:"isBatchAssigned" bson:"isBatchAssigned"`
	IsBatchAccepted    bool                      `json:"isBatchAccepted" bson:"isBatchAccepted"`
	IsFirstStartPickup bool                      `json:"isFirstStartPickup" bson:"isFirstStartPickup"`
	StartTime          int64                     `json:"startTime" bson:"startTime"`
	Duration           int64                     `json:"duration" bson:"duration"`
	TimeRemaining      int64                     `json:"timeRemaining" bson:"timeRemaining"`
	Redirect           string                    `json:"redirect" bson:"redirect"`
	Polling            int                       `json:"polling" bson:"polling"`
}

type BatchNavigation struct {
	Id primitive.ObjectID `json:"_id" bson:"_id"`

	BatchId     string   `json:"batchId" bson:"batchId"`
	Status      string   `json:"status" bson:"status"`
	PickId      string   `json:"pickId" bson:"pickId"`
	DropId      string   `json:"dropId" bson:"dropId"`
	NextStatus  string   `json:"nextStatus" bson:"nextStatus"`
	TrackingIds []string `json:"trackingIds" bson:"trackingByIds"`
	Redirect    string   `json:"redirect" bson:"redirect"`
	CreatedAt   int64    `json:"createdAt" bson:"createdAt"`
	Polling     int      `json:"polling" bson:"polling"`
}

type BatchResponse struct {
	BatchId string   `json:"batchId,omitempty"`
	Nodes   []*BNode `json:"nodes,omitempty"`
}

type BNode struct {
	Id         string `json:"id,omitempty"`
	Type       string `json:"type,omitempty"`
	Status     string `json:"status,omitempty"`
	TrackingId string `json:"trackingId,omitempty"`
	PickupId   string `json:"pickupId,omitempty"`
}

func SetBatchNavigationData(batchId, statusCode, pickId, dropId, reqId string, trackingIds []string, pollingInterVal, pollingStopInterVal int32, batchRes *BatchResponse) BatchNavigation {
	log.Println(lg.Debug(reqId, ":"), lg.Info(" set batch navigation data: ", batchId))

	batchNavigation := BatchNavigation{}
	batchNavigation.BatchId = batchId
	batchNavigation.Status = statusCode

	if len(batchRes.Nodes) > 0 {
		if strings.EqualFold(statusCode, "BA") {
			log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
			batchNavigation.Polling = int(pollingStopInterVal)
			batchNavigation.PickId = pickId
			batchNavigation.NextStatus = "PS"
			batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/pickup/", batchNavigation.PickId, "/START_PICK_UP")
		}
		if strings.EqualFold(statusCode, "PS") {
			log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
			batchNavigation.Polling = int(pollingStopInterVal)
			batchNavigation.PickId = pickId
			batchNavigation.NextStatus = "PS"
			batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/pickup/", batchNavigation.PickId, "/ARRIVED_AT_PICKUP")
		}
		if strings.EqualFold(statusCode, "PA") {
			log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
			batchNavigation.Polling = int(pollingInterVal)
			batchNavigation.PickId = pickId
			batchNavigation.NextStatus = "PICK_LIST"
			batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/pickup/", batchNavigation.PickId, "/PICK_UP_CHECKLIST")
		}
		if strings.EqualFold(statusCode, "PP") {
			log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
			pplist := lo.Filter(batchRes.Nodes, func(n *BNode, i int) bool {
				return strings.EqualFold(pickId, n.PickupId) && strings.EqualFold(n.Type, "DROP")
			})
			checked := false
			//If all parcels are not picked yet
			for _, pp := range pplist {

				log.Println(lg.Debug(reqId, ":"), lg.Info(" pp.Status: ", pp.Status))
				log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
				if !strings.EqualFold(pp.Status, "PP") && !strings.EqualFold(pp.Status, "PF") {
					log.Println(lg.Debug(reqId, ":"), " All parcels are not picked yet. Stay at pickup list screen")
					batchNavigation.Polling = int(pollingInterVal)
					batchNavigation.PickId = pickId
					batchNavigation.NextStatus = "PICK_LIST"
					batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/pickup/", batchNavigation.PickId, "/PICK_UP_CHECKLIST")
					checked = true
					break
				}
			}

			//If all parcels are already picked
			if !checked {

				for _, pp := range pplist {

					log.Println(lg.Debug(reqId, ":"), lg.Info(" pp.Status: ", pp.Status))
					log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))

					log.Println(lg.Debug(reqId, ":"), " All parcels are picked up. Check the next node pickup or dropoff")
					pickNodeIndex := -1
					log.Println(lg.Debug(reqId, ":"), lg.Info(" pickId: ", pickId))
					for i, po := range batchRes.Nodes {
						if strings.EqualFold(pickId, po.PickupId) && strings.EqualFold(po.Type, "PICK") {
							pickNodeIndex = i
							break
						}
					}
					log.Println(lg.Debug(reqId, ":"), lg.Info(" pickNodeIndex: ", pickNodeIndex))
					log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
					nextNodeIndex := pickNodeIndex + 1

					//Check the next node exists or not
					if nextNodeIndex < 0 || nextNodeIndex >= len(batchRes.Nodes) {
						// Handle the index out of bounds error
						log.Println(lg.Debug(reqId, ":"), " Next node does not exist")

						batchNavigation.Polling = int(pollingStopInterVal)
						batchNavigation.PickId = pickId
						batchNavigation.NextStatus = "NONE"
						batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/pickup/", batchNavigation.PickId, "/NONE")
					} else {
						log.Println(lg.Debug(reqId, ":"), lg.Info(" Next node index: ", nextNodeIndex))
						nodeType := batchRes.Nodes[nextNodeIndex].Type
						log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$nodeType$$$$$$$$$$: ", nodeType, " == ", batchId))
						if strings.EqualFold(nodeType, "PICK") {
							log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
							batchNavigation.Polling = int(pollingStopInterVal)
							batchNavigation.PickId = batchRes.Nodes[nextNodeIndex].PickupId
							batchNavigation.NextStatus = "PS"
							batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/pickup/", batchNavigation.PickId, "/START_PICK_UP")

						}
						if strings.EqualFold(nodeType, "DROP") {
							log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
							batchNavigation.Polling = int(pollingInterVal)
							batchNavigation.PickId = pickId
							for _, n := range batchRes.Nodes {

								if strings.EqualFold(n.PickupId, pickId) && strings.EqualFold(n.Type, "DROP") && strings.EqualFold(n.Status, "PP") {
									batchNavigation.NextStatus = "SD"
									batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/dropoff/", n.Id, "/START_DROP_OFF?trackingIds=", n.TrackingId)
									break
								}
							}

						}
					}

				}
			}

		}
		if strings.EqualFold(statusCode, "PF") {
			log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
			pplist := lo.Filter(batchRes.Nodes, func(n *BNode, i int) bool {
				return strings.EqualFold(pickId, n.PickupId) && strings.EqualFold(n.Type, "DROP")
			})
			for _, pp := range pplist {
				//If all parcels are not picked yet
				log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
				if !strings.EqualFold(pp.Status, "PP") && !strings.EqualFold(pp.Status, "PF") {
					log.Println(lg.Debug(reqId, ":"), " All parcels are not picked yet. Stay at pickup list screen")
					batchNavigation.Polling = int(pollingInterVal)
					batchNavigation.PickId = pickId
					batchNavigation.NextStatus = "PP"
					batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/pickup/", batchNavigation.PickId, "/PICK_UP_CHECKLIST")
				} else {
					log.Println(lg.Debug(reqId, ":"), " All parcels are picked up. Check the next node pickup or dropoff")
					pickNodeIndex := -1
					for i, po := range batchRes.Nodes {
						if strings.EqualFold(pickId, po.PickupId) && strings.EqualFold(po.Type, "PICK") {
							pickNodeIndex = i
							break
						}
					}
					log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
					nextNodeIndex := pickNodeIndex + 1

					//Check the next node exists or not
					if nextNodeIndex < 0 || nextNodeIndex >= len(batchRes.Nodes) {
						log.Println(lg.Debug(reqId, ":"), " Next node does not exist")
						// Handle the index out of bounds error
						batchNavigation.PickId = pickId
						batchNavigation.NextStatus = "NONE"
						batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/pickup/", batchNavigation.PickId, "/NONE")
					} else {
						log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
						nodeType := batchRes.Nodes[nextNodeIndex].Type

						if strings.EqualFold(nodeType, "PICK") {
							log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
							batchNavigation.PickId = batchRes.Nodes[nextNodeIndex].PickupId
							batchNavigation.NextStatus = "PS"
							batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/pickup/", batchNavigation.PickId, "/START_PICK_UP")

						}
						if strings.EqualFold(nodeType, "DROP") {
							log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
							batchNavigation.Polling = int(pollingInterVal)
							batchNavigation.PickId = pickId
							for _, n := range batchRes.Nodes {

								if strings.EqualFold(n.PickupId, pickId) && strings.EqualFold(n.Type, "DROP") && strings.EqualFold(n.Status, "PP") {
									batchNavigation.NextStatus = "SD"
									batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/dropoff/", n.Id, "/START_DROP_OFF?trackingIds=", n.TrackingId)
									break
								}
							}

						}
						log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
					}
					log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
				}
				log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
			}

		}
		if strings.EqualFold(statusCode, "SD") {
			log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
			batchNavigation.Polling = int(pollingStopInterVal)
			batchNavigation.PickId = pickId
			batchNavigation.NextStatus = "AD"
			batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/dropoff/", dropId, "/ARRIVED_AT_DROP_OFF?trackingIds=", strings.Join(trackingIds, ","))
		}
		if strings.EqualFold(statusCode, "AD") {
			log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
			batchNavigation.Polling = int(pollingStopInterVal)
			batchNavigation.PickId = pickId
			batchNavigation.NextStatus = "DROP_LIST"
			batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/dropoff/", dropId, "/DROP_OFF_CHECKLIST?trackingIds=", strings.Join(trackingIds, ","))
		}
		if strings.EqualFold(statusCode, "DL") {
			log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
			nodeIndex := -1
			for i, po := range batchRes.Nodes {
				if strings.EqualFold(dropId, po.Id) && strings.EqualFold(po.Type, "DROP") {
					nodeIndex = i
					break
				}
			}

			nextIndex := nodeIndex + 1
			//Check the next node exists or not
			if nextIndex < 0 || nextIndex >= len(batchRes.Nodes) {
				log.Println(lg.Debug(reqId, ":"), " Next node does not exist")
				// Handle the index out of bounds error
				batchNavigation.Polling = int(pollingStopInterVal)
				batchNavigation.PickId = pickId
				batchNavigation.NextStatus = "NONE"
				batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/pickup/", batchNavigation.PickId, "/NONE")
			} else {
				isAd := false
				log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
				for _, po := range batchRes.Nodes {
					log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
					if strings.EqualFold(dropId, po.Id) && strings.EqualFold(po.Type, "DROP") && strings.EqualFold(po.Status, "AD") {
						log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
						batchNavigation.Polling = int(pollingStopInterVal)
						batchNavigation.PickId = pickId
						batchNavigation.NextStatus = "DROP_LIST"
						batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/dropoff/", dropId, "/DROP_OFF_CHECKLIST?trackingIds=", strings.Join(trackingIds, ","))
						isAd = true
						break
					}
				}
				if !isAd {
					log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
					nextNode := batchRes.Nodes[nextIndex]
					if strings.EqualFold(nextNode.Type, "PICK") {
						log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
						batchNavigation.Polling = int(pollingStopInterVal)
						batchNavigation.PickId = nextNode.PickupId
						batchNavigation.NextStatus = "PS"
						batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/pickup/", batchNavigation.PickId, "/START_PICK_UP")

					}
					if strings.EqualFold(nextNode.Type, "DROP") {
						log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
						batchNavigation.Polling = int(pollingStopInterVal)
						batchNavigation.PickId = nextNode.PickupId
						batchNavigation.NextStatus = "SD"
						batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/dropoff/", nextNode.Id, "/START_DROP_OFF?trackingIds=", nextNode.TrackingId)
					}
					log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
				}
				log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))

			}
			log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))

		}
		if strings.EqualFold(statusCode, "DF") {
			log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
			batchNavigation.Polling = int(pollingStopInterVal)
			batchNavigation.PickId = pickId
			batchNavigation.NextStatus = "RS"
			batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/dropoff/", dropId, "/START_RETURN?trackingIds=", strings.Join(trackingIds, ","))
		}
		if strings.EqualFold(statusCode, "RS") {
			log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
			batchNavigation.Polling = int(pollingStopInterVal)
			batchNavigation.PickId = pickId
			batchNavigation.NextStatus = "SA"
			batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/dropoff/", dropId, "/ARRIVED_AT_SENDER?trackingIds=", strings.Join(trackingIds, ","))
		}
		if strings.EqualFold(statusCode, "SA") {
			log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
			batchNavigation.Polling = int(pollingStopInterVal)
			batchNavigation.PickId = pickId
			batchNavigation.NextStatus = "RTS"
			batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/dropoff/", dropId, "/RETURN_CHECKLIST?trackingIds=", strings.Join(trackingIds, ","))
		}
		if strings.EqualFold(statusCode, "RTS") {
			log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
			nodeIndex := -1
			for i, po := range batchRes.Nodes {
				if strings.EqualFold(dropId, po.Id) && strings.EqualFold(po.Type, "DROP") {
					nodeIndex = i
					break
				}
			}

			nextIndex := nodeIndex + 1

			//Check the next node exists or not
			if nextIndex < 0 || nextIndex >= len(batchRes.Nodes) {
				log.Println(lg.Debug(reqId, ":"), " Next node does not exist")
				// Handle the index out of bounds error
				batchNavigation.Polling = int(pollingStopInterVal)
				batchNavigation.PickId = pickId
				batchNavigation.NextStatus = "NONE"
				batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/pickup/", batchNavigation.PickId, "/NONE")
			} else {

				isAd := false
				log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
				for _, po := range batchRes.Nodes {
					log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
					if strings.EqualFold(dropId, po.Id) && strings.EqualFold(po.Type, "DROP") && strings.EqualFold(po.Status, "AD") {
						log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
						batchNavigation.Polling = int(pollingStopInterVal)
						batchNavigation.PickId = pickId
						batchNavigation.NextStatus = "DROP_LIST"
						batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/dropoff/", dropId, "/DROP_OFF_CHECKLIST?trackingIds=", strings.Join(trackingIds, ","))
						isAd = true
						break
					}
				}
				if !isAd {
					log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
					nextNode := batchRes.Nodes[nextIndex]
					if strings.EqualFold(nextNode.Type, "PICK") {
						log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
						batchNavigation.Polling = int(pollingStopInterVal)
						batchNavigation.PickId = nextNode.PickupId
						batchNavigation.NextStatus = "PS"
						batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/pickup/", batchNavigation.PickId, "/START_PICK_UP")

					}
					if strings.EqualFold(nextNode.Type, "DROP") {
						log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
						batchNavigation.Polling = int(pollingStopInterVal)
						batchNavigation.PickId = nextNode.PickupId
						batchNavigation.NextStatus = "SD"
						batchNavigation.Redirect = fmt.Sprint("/batch/", batchId, "/dropoff/", nextNode.Id, "/START_DROP_OFF?trackingIds=", nextNode.TrackingId)
					}
				}

				log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))
			}
			log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))

		}
		log.Println(lg.Debug(reqId, ":"), lg.Mg(" $$$$$$$$$$$$$$$$$$$$: ", batchId))

	}

	return batchNavigation
}

func StatusMapping(code string) int {
	val := 0
	switch code {
	case "BA":
		val = 1
	case "PS":
		val = 2
	case "PA":
		val = 3
	case "PP":
		val = 4
	case "PF":
		val = 5
	case "SD":
		val = 6
	case "AD":
		val = 7
	case "DL":
		val = 8
	case "DF":
		val = 9
	case "RS":
		val = 10
	case "SA":
		val = 11
	case "RTS":
		val = 12
	}

	return val
}

func GenerateBatchNavigationData(d []order.BatchForDriverApp, batchId, reqId string, pollingInterVal, pollingStopInterVal int32) (BatchNavigation, error) {

	log.Println(lg.Debug(reqId, ": "), lg.Mg("GenerateBatchNavigation: ", batchId))
	pickId := ""
	dropId := ""
	st := ""
	stCode := 0
	track := ""

	ns := lo.Filter(d, func(n order.BatchForDriverApp, i int) bool {
		return strings.EqualFold(n.Type, "PICKUP")
	})

	avoidPickups := d[0].AvoidPickups
	avoidPickupId := ""
	if len(avoidPickups) > 0 {
		avoidPickup := avoidPickups[len(avoidPickups)-1]
		avoidPickupStr := fmt.Sprint("PICK@@@", strings.ToLower(avoidPickup.Name), "@@@", strings.ToLower(avoidPickup.Address))
		avoidPickupId = base64.StdEncoding.EncodeToString([]byte(avoidPickupStr))
		s := 0
		log.Println(lg.Debug(reqId, ": "), lg.Green("len(ns): ", len(ns)))
		log.Println(lg.Debug(reqId, ": "), lg.Green("Avoid pickup: ", avoidPickupStr))
		if len(ns) > 0 {
			log.Println(lg.Debug(reqId, ": "), lg.Info("================================"))
			log.Println(lg.Info("Batch navigation: ", pickId))

			for i, nso := range ns {
				if strings.EqualFold(nso.Id, avoidPickupId) && i < len(ns) {
					log.Println(lg.Debug(reqId, ": "), lg.Info("============avoid pickup index: ", i, "===================="))
					s++
				}
				if s > 0 {

					ds := lo.Filter(d, func(n order.BatchForDriverApp, i int) bool {
						return strings.EqualFold(n.Type, "DROPOFF") && strings.EqualFold(n.PickupId, nso.Id)
					})

					log.Println(lg.Debug(reqId, ": "), lg.Info("============drop count: ", len(ds), "===================="))

					for _, ord := range ds {

						for j, od := range ord.Orders {

							log.Println(lg.Debug(reqId, ": "), lg.Info("==============", od.Status, "=================="))
							ind := StatusMapping(od.Status)
							log.Println(lg.Debug(reqId, ": "), lg.Info("===========current code before check: ", ind, "=========current status: ", ord.Orders[0].Status, "============"))
							if ind > stCode && ind != 8 && ind != 12 {
								log.Println(lg.Debug(reqId, ": "), lg.Info("===========current code after check: ", ind, "=========current status: ", ord.Orders[0].Status, "============"))
								stCode = ind
								st = od.Status
								track = od.TrackingId
								pickId = ord.PickupId
								dropId = ord.Id
							}
							log.Println(lg.Debug(reqId, ": "), lg.Info("==============", j, "=================="))
						}

					}

					log.Println(lg.Debug(reqId, ": "), lg.Yellow("=====Pickup ID: =====", pickId, "========Status code: ====", st, "===dropID===", dropId))

				}
			}

			log.Println(lg.Debug(reqId, ": "), lg.Mg("=====Pickup ID: =====", pickId, "========Status code: ====", st, "===dropID===", dropId))

		}

		log.Println(lg.Debug(reqId, ": "), lg.Dl("=====Pickup ID: =====", pickId, "========Status code: ====", st, "===dropID===", dropId))

	} else {

		log.Println(lg.Debug(reqId, ": "), lg.Green("len(ns): ", len(ns)))
		if len(ns) > 0 {
			log.Println(lg.Debug(reqId, ": "), lg.Info("==========No avoid pickup======================"))
			log.Println(lg.Debug(reqId, ": "), lg.Info("Batch navigation: ", pickId))

			ds := lo.Filter(d, func(n order.BatchForDriverApp, i int) bool {
				return strings.EqualFold(n.Type, "DROPOFF") && strings.EqualFold(n.PickupId, ns[0].Id)
			})

			for _, ord := range ds {

				for j, od := range ord.Orders {

					log.Println(lg.Debug(reqId, ": "), lg.Info("==============", od.Status, "=================="))
					ind := StatusMapping(od.Status)
					log.Println(lg.Debug(reqId, ": "), lg.Info("===========current code before check: ", ind, "=========current status: ", ord.Orders[0].Status, "============"))
					if ind > stCode && ind != 8 && ind != 12 {
						log.Println(lg.Debug(reqId, ": "), lg.Info("===========current code after check: ", ind, "=========current status: ", ord.Orders[0].Status, "============"))
						stCode = ind
						st = od.Status
						track = od.TrackingId
						pickId = ord.PickupId
						dropId = ord.Id
					}
					log.Println(lg.Debug(reqId, ": "), lg.Info("==============", j, "=================="))
				}

			}

			log.Println(lg.Debug(reqId, ": "), lg.Yellow("=====Pickup ID: =====", pickId, "========Status code: ====", st, "===dropID===", dropId))

		}
	}

	bn := BatchNavigation{}
	bn.BatchId = batchId
	bn.DropId = dropId
	bn.PickId = pickId
	bn.Status = st
	bn.TrackingIds = []string{track}

	bnByte, _ := json.Marshal(&bn)
	log.Println(lg.Green(string(bnByte)))

	return bn, nil
}
