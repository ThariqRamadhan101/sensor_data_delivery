package main

import (
	"encoding/json"
	"net/http"
	"time"
	"strconv"
	"math/rand"
	"io/ioutil"
	"fmt"

	"github.com/gorilla/mux"
)

func getAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(clients)
	json.NewEncoder(w).Encode(dataClients)
	
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
    return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
    bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        bytes[i] = byte(randomInt(65, 90))
    }
    return string(bytes)
}

func getRegistration(w http.ResponseWriter, r *http.Request) {
	idNum := strconv.Itoa(len(clients)+1)
	id := idNum + "_" + randomString(4)
	time := time.Now().Format("02-01-2006 15:04:05")
	var newClient = client{
		ID:          		id,
		TIME_REGISTRATION:	time,
		STATUS:				"on",
	}

	clients = append(clients, newClient)

	response := registration{
		ID: newClient.ID,
		TIME: newClient.TIME_REGISTRATION,
	}
	json.NewEncoder(w).Encode(response)

}

func getQuestion(w http.ResponseWriter, r *http.Request) {
	clientID := mux.Vars(r)["id"]

	for _, client := range clients {
		if client.ID == clientID {
			status := client.STATUS
			response := question{
				STATUS: status,
			}
			json.NewEncoder(w).Encode(response)
		}
	}
}



func getData(w http.ResponseWriter, r *http.Request) {
	var newData dataClient
	reqBody, err := ioutil.ReadAll(r.Body)
			
	if err != nil {
		fmt.Fprintf(w, "ERROR: DATA TYPE ENTERED NOT MATCH")
	}
	json.Unmarshal(reqBody, &newData)

	for _, client := range clients {
		if client.ID == newData.ID{
			dataClients = append(dataClients, newData)
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "The data client with ID %v has been added successfully", newData.ID)
		}
	}
	
}

func getReset(w http.ResponseWriter, r *http.Request) {
	clientID := mux.Vars(r)["id"]

	for i, client := range clients {
		if client.ID == clientID {
			clients = append(clients[:i], clients[i+1:]...)
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "The client with ID %v has been deleted successfully", clientID)
		}
	}

	for i, dataClient := range dataClients {
		if dataClient.ID == clientID {
			dataClients = append(dataClients[:i], dataClients[i+1:]...)
			// fmt.Fprintf(w, "Data client with ID %v has been deleted successfully", clientID)
		}
	}
}

func getSync(w http.ResponseWriter, r *http.Request) {
	clientID := mux.Vars(r)["id"]

	for _, client := range clients {
		if client.ID == clientID {
			time := time.Now().Format("02-01-2006 15:04:05")
			total := 0
			for _, dataClient := range dataClients {
				if dataClient.ID == clientID {
					total = total + 1
				}
			}
			response := sync{
				TIME: time,
				TOTAL: total,
			}
			json.NewEncoder(w).Encode(response)
		}
	}
}

func getHistory (w http.ResponseWriter, r *http.Request) {
	clientID := mux.Vars(r)["id"]
	nString := mux.Vars(r)["n"]
	n, _ := strconv.Atoi(nString)
	for _, client := range clients {
		if client.ID == clientID {
			flag := 0
			for _, dataClient := range dataClients {
				if dataClient.ID == clientID {
					flag = flag + 1
					if flag == n {
						sensor := dataClient.SENSOR
						value := dataClient.VALUE
						time := dataClient.TIME
						response := history{
							SENSOR: sensor,
							VALUE: value,
							TIME: time,
						}
						json.NewEncoder(w).Encode(response)
					}
				}
			}
			
		}
	}
}