package main

import (
	"encoding/json"
<<<<<<< HEAD
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/nubo/jwt"
)

var secret = "bakuretsu"

func getAll(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	var db_credential credential
	var arr_db_credential []credential
	var db_client client
	var arr_db_client []client
	var db_dataclient dataClient
	var arr_db_dataclient []dataClient
	
	rowsCred, err := db.Query("Select credential_pass, used_id from credentials")
	if err != nil {
		log.Print(err)
	}

	for rowsCred.Next() {
	    err := rowsCred.Scan(&db_credential.CREDENTIAL_PASS, &db_credential.USED_ID); 
		if err != nil {
			log.Fatal(err.Error())
		}else {
			arr_db_credential = append(arr_db_credential, db_credential)
		}
	}

	rowsClient, err := db.Query("Select id, time_registration, status from clients")
	if err != nil {
		log.Print(err)
	}

	for rowsClient.Next() {
	    err := rowsClient.Scan(&db_client.ID, &db_client.TIME_REGISTRATION, &db_client.STATUS); 
		if err != nil {
			log.Fatal(err.Error())
		}else {
			arr_db_client = append(arr_db_client, db_client)
		}
	}
	
	rowsData, err := db.Query("Select id, sensor, value, time from dataclients ")
	if err != nil {
		log.Print(err)
	}

	for rowsData.Next() {
	    err := rowsData.Scan(&db_dataclient.ID, &db_dataclient.SENSOR, &db_dataclient.VALUE, &db_dataclient.TIME); 
		if err != nil {
			log.Fatal(err.Error())
		}else {
			arr_db_dataclient = append(arr_db_dataclient, db_dataclient)
		}
	}
	json.NewEncoder(w).Encode(arr_db_credential)
	json.NewEncoder(w).Encode(arr_db_client)
	json.NewEncoder(w).Encode(arr_db_dataclient)
	
	log.Print("Request all data.")
=======
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
	
>>>>>>> 1faa28365bc6d90da37a612597569c569b53f5a3
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
<<<<<<< HEAD
	return min + rand.Intn(max-min)
=======
    return min + rand.Intn(max-min)
>>>>>>> 1faa28365bc6d90da37a612597569c569b53f5a3
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
<<<<<<< HEAD
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

func getRegistration(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	var db_credential credential
	var arr_db_credential []credential
	
	rowsCred, err := db.Query("Select credential_pass, used_id from credentials")
	if err != nil {
		log.Print(err)
	}

	for rowsCred.Next() {
	    err := rowsCred.Scan(&db_credential.CREDENTIAL_PASS, &db_credential.USED_ID); 
		if err != nil {
			log.Fatal(err.Error())
		}else {
			arr_db_credential = append(arr_db_credential, db_credential)
		}
	}

	var db_client client
	var arr_db_client []client	
	
	rows, err := db.Query("Select id, time_registration, status from clients")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
	    err := rows.Scan(&db_client.ID, &db_client.TIME_REGISTRATION, &db_client.STATUS); 
		if err != nil {
			log.Fatal(err.Error())
		}else {
			arr_db_client = append(arr_db_client, db_client)
		}
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "ERROR: DATA TYPE ENTERED NOT MATCH")
	}
	token, ok := jwt.ParseAndVerify(string(reqBody), secret)
	if ok {
		var cred_pass = token.ClaimSet["credential_pass"].(string)
		for _, credential := range arr_db_credential {
			if ((credential.CREDENTIAL_PASS == cred_pass) && (credential.USED_ID == "none")) {

				idNum := strconv.Itoa(len(arr_db_client) + 1)
				id := idNum + "_" + randomString(4)
				time := time.Now().Format("02-01-2006 15:04:05")
				var newClient = client{
					ID:                id,
					TIME_REGISTRATION: time,
					STATUS:            "on",
				}

				_, err = db.Exec("UPDATE credentials set used_id = ? where credential_pass = ?",
					id,
					cred_pass,
				)

				if err != nil {
					log.Print(err)
				}

				_, err = db.Exec("INSERT INTO clients (id, time_registration, status) values (?,?,?)",
					newClient.ID,
					newClient.TIME_REGISTRATION,
					newClient.STATUS,
				)
				if err != nil {
					log.Print(err)
				}

				claims := jwt.ClaimSet{
					"id":   newClient.ID,
					"time": newClient.TIME_REGISTRATION,
				}
				token, err := claims.Sign(secret)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Fprint(w, token)
				// fmt.Println(token)

				log.Print("Client with ID " + newClient.ID + " created.")
							
			}
		}	
			
	} else {
		log.Fatal("Invalid token")
	}

=======
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
>>>>>>> 1faa28365bc6d90da37a612597569c569b53f5a3

}

func getQuestion(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	db := connect()
	defer db.Close()
	
	var db_client client
	var arr_db_client []client
	
	rows, err := db.Query("Select id, time_registration, status from clients")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
	    err := rows.Scan(&db_client.ID, &db_client.TIME_REGISTRATION, &db_client.STATUS); 
		if err != nil {
			log.Fatal(err.Error())
		}else {
			arr_db_client = append(arr_db_client, db_client)
		}
	}
	
	clientID := mux.Vars(r)["id"]
	for _, client := range arr_db_client {
		if client.ID == clientID {
			status := client.STATUS

			claims := jwt.ClaimSet{
				"status": status,
			}
			token, err := claims.Sign(secret)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprint(w, token)
			// fmt.Println(token)
		}
	}
	log.Print("Client with ID " + clientID + " ask question.")
}

func getData(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	
	var db_client client
	var arr_db_client []client
	var db_dataclient dataClient
	var arr_db_dataclient []dataClient
	
	rows, err := db.Query("Select id, time_registration, status from clients")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
	    err := rows.Scan(&db_client.ID, &db_client.TIME_REGISTRATION, &db_client.STATUS); 
		if err != nil {
			log.Fatal(err.Error())
		}else {
			arr_db_client = append(arr_db_client, db_client)
		}
	}
	
	rows1, err := db.Query("Select id, sensor, value, time from dataclients ")
	if err != nil {
		log.Print(err)
	}

	for rows1.Next() {
	    err := rows1.Scan(&db_dataclient.ID, &db_dataclient.SENSOR, &db_dataclient.VALUE, &db_dataclient.TIME); 
		if err != nil {
			log.Fatal(err.Error())
		}else {
			arr_db_dataclient = append(arr_db_dataclient, db_dataclient)
		}
	}
	
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "ERROR: DATA TYPE ENTERED NOT MATCH")
	}

	token, ok := jwt.ParseAndVerify(string(reqBody), secret)
	if ok {
		var newData dataClient

		newData.ID = token.ClaimSet["id"].(string)
		newData.SENSOR = token.ClaimSet["sensor"].(string)
		newData.VALUE = token.ClaimSet["value"].(string)
		newData.TIME = token.ClaimSet["time"].(string)

		for _, client := range arr_db_client {
			if client.ID == newData.ID {
				_, err = db.Exec("INSERT INTO dataclients (id, sensor, value, time) values (?,?,?,?)",
					newData.ID,
					newData.SENSOR,
					newData.VALUE,
					newData.TIME,
				)
				if err != nil {
					log.Print(err)
				}
		
				w.WriteHeader(http.StatusCreated)
				msg := "The data client with ID " + newData.ID + " has been added successfully"
				claims := jwt.ClaimSet{
					"message": msg,
				}
				token, _ := claims.Sign(secret)
				
				fmt.Fprint(w, token)
				// fmt.Println(token)

				log.Print("Client with ID " + newData.ID + " send new data with sensor name:" + newData.SENSOR + ", sensor value: " + newData.VALUE + ".")
			}
		}		
	} else {
		log.Fatal("Invalid token")
	}

=======
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
>>>>>>> 1faa28365bc6d90da37a612597569c569b53f5a3
	
}

func getReset(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	db := connect()
	defer db.Close()

	var db_credential credential
	var arr_db_credential []credential
	var db_client client
	var arr_db_client []client
	var db_dataclient dataClient
	var arr_db_dataclient []dataClient
	
	rowsCred, err := db.Query("Select credential_pass, used_id from credentials")
	if err != nil {
		log.Print(err)
	}

	for rowsCred.Next() {
	    err := rowsCred.Scan(&db_credential.CREDENTIAL_PASS, &db_credential.USED_ID); 
		if err != nil {
			log.Fatal(err.Error())
		}else {
			arr_db_credential = append(arr_db_credential, db_credential)
		}
	}
	
	rows, err := db.Query("Select id, time_registration, status from clients")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
	    err := rows.Scan(&db_client.ID, &db_client.TIME_REGISTRATION, &db_client.STATUS); 
		if err != nil {
			log.Fatal(err.Error())
		}else {
			arr_db_client = append(arr_db_client, db_client)
		}
	}
	
	rows1, err := db.Query("Select id, sensor, value, time from dataclients ")
	if err != nil {
		log.Print(err)
	}

	for rows1.Next() {
	    err := rows1.Scan(&db_dataclient.ID, &db_dataclient.SENSOR, &db_dataclient.VALUE, &db_dataclient.TIME); 
		if err != nil {
			log.Fatal(err.Error())
		}else {
			arr_db_dataclient = append(arr_db_dataclient, db_dataclient)
		}
	}
	
	clientID := mux.Vars(r)["id"]

	for _, client := range arr_db_client{
		if client.ID == clientID {
			_, err = db.Exec("DELETE from clients where id = ?", 
				clientID,
			)
			if err != nil {
			log.Print(err)
			}
			_, err = db.Exec("UPDATE credentials set used_id = ? where used_id = ?",
				"none",
				client.ID,
			)
			if err != nil {
				log.Print(err)
			}

			w.WriteHeader(http.StatusCreated)
			msg := "The data client with ID " + clientID + " has been deleted successfully"
			claims := jwt.ClaimSet{
				"message": msg,
			}
			token, err := claims.Sign(secret)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprint(w, token)
			// fmt.Println(token)
		}
	}

	for _, dataClient := range arr_db_dataclient {
		if dataClient.ID == clientID {
			_, err = db.Exec("DELETE from dataclients where id = ?", 
				clientID,
			)
			if err != nil {
			log.Print(err)
			}
		}
	}
	log.Print("Client with ID " + clientID + " deleted.")
}

func getSync(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	
	var db_client client
	var arr_db_client []client
	var db_dataclient dataClient
	var arr_db_dataclient []dataClient
	
	rows, err := db.Query("Select id, time_registration, status from clients")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
	    err := rows.Scan(&db_client.ID, &db_client.TIME_REGISTRATION, &db_client.STATUS); 
		if err != nil {
			log.Fatal(err.Error())
		}else {
			arr_db_client = append(arr_db_client, db_client)
		}
	}
	
	rows1, err := db.Query("Select id, sensor, value, time from dataclients ")
	if err != nil {
		log.Print(err)
	}

	for rows1.Next() {
	    err := rows1.Scan(&db_dataclient.ID, &db_dataclient.SENSOR, &db_dataclient.VALUE, &db_dataclient.TIME); 
		if err != nil {
			log.Fatal(err.Error())
		}else {
			arr_db_dataclient = append(arr_db_dataclient, db_dataclient)
		}
	}
	
	clientID := mux.Vars(r)["id"]

	for _, client := range arr_db_client {
		if client.ID == clientID {
			time := time.Now().Format("02-01-2006 15:04:05")
			total := 0
			for _, dataClient := range arr_db_dataclient {
=======
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
>>>>>>> 1faa28365bc6d90da37a612597569c569b53f5a3
				if dataClient.ID == clientID {
					total = total + 1
				}
			}
<<<<<<< HEAD
			claims := jwt.ClaimSet{
				"time":  time,
				"total": total,
			}
			token, err := claims.Sign(secret)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprint(w, token)
			// fmt.Println(token)

		}
	}
	log.Print("Client with ID " + clientID + " request syncronization.")
}

func getHistory(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	
	var db_client client
	var arr_db_client []client
	var db_dataclient dataClient
	var arr_db_dataclient []dataClient
	
	rows, err := db.Query("Select id, time_registration, status from clients")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
	    err := rows.Scan(&db_client.ID, &db_client.TIME_REGISTRATION, &db_client.STATUS); 
		if err != nil {
			log.Fatal(err.Error())
		}else {
			arr_db_client = append(arr_db_client, db_client)
		}
	}
	
	rows1, err := db.Query("Select id, sensor, value, time from dataclients ")
	if err != nil {
		log.Print(err)
	}

	for rows1.Next() {
	    err := rows1.Scan(&db_dataclient.ID, &db_dataclient.SENSOR, &db_dataclient.VALUE, &db_dataclient.TIME); 
		if err != nil {
			log.Fatal(err.Error())
		}else {
			arr_db_dataclient = append(arr_db_dataclient, db_dataclient)
		}
	}
	clientID := mux.Vars(r)["id"]
	nString := mux.Vars(r)["n"]
	n, _ := strconv.Atoi(nString)
	for _, client := range arr_db_client {
		if client.ID == clientID {
			flag := 0
			for _, dataClient := range arr_db_dataclient {
=======
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
>>>>>>> 1faa28365bc6d90da37a612597569c569b53f5a3
				if dataClient.ID == clientID {
					flag = flag + 1
					if flag == n {
						sensor := dataClient.SENSOR
						value := dataClient.VALUE
						time := dataClient.TIME
<<<<<<< HEAD
						claims := jwt.ClaimSet{
							"sensor": sensor,
							"value":  value,
							"time":   time,
						}
						token, err := claims.Sign(secret)
						if err != nil {
							log.Fatal(err)
						}

						fmt.Fprint(w, token)
						// fmt.Println(token)
					}
				}
			}

		}
	}
	log.Print("Client with ID " + clientID + " request history data-" + nString + ".")
}
=======
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
>>>>>>> 1faa28365bc6d90da37a612597569c569b53f5a3
