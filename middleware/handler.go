package middleware

import (
    "database/sql"
    "encoding/json" // package to encode and decode the json into struct and vice versa
    "fmt"
    "log"
    "net/http" // used to access the request and response object of the api
    "os"       // used to read the environment variable
    "carParking/models"
    "github.com/joho/godotenv" // package used to read the .env file
    _ "github.com/lib/pq"      // postgres golang driver
    "github.com/gorilla/mux" // used to get the params from the route
)

// response format
type response struct {
    ID      int64  `json:"id,omitempty"`
    Message string `json:"message,omitempty"`
}

// create connection with postgres db
func createConnection() *sql.DB {
    // load .env file
    err := godotenv.Load(".env")

    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Open the connection
    db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

    if err != nil {
        panic(err)
    }

    // check the connection
    err = db.Ping()

    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully connected!")
    // return the connection
    return db
}

// GetAllUser will return all the users
func GetFreePakingSlot(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "PUT")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    // get all the users in the db

    params := mux.Vars(r)
    log.Println(params["carId"])
    log.Println(params["ownerName"])
    users, err := getFreePaking(params["carId"],params["ownerName"])
    log.Println("okkkk bhai")
    log.Println(users)
    log.Println("okkkk bhai")

    if err != nil {
        log.Fatalf("Unable to get Free Parking Slot. %v", err)
    }

    // send all the users as response
    json.NewEncoder(w).Encode(users)
}

// get one user from the DB by its userid
func getFreePaking(carId string, ownerName string) (string, error) {
    // create the postgres db connection
    db := createConnection()
    slotStatus := ""
    // close the db connection
    defer db.Close()

    var users []models.ParkingDetails

    // create the select sql query
    sqlStatement := `SELECT "PakingId" FROM public."Parking" where "CarId"='' LIMIT 1 ;`

    // execute the sql statement
    rows, err := db.Query(sqlStatement)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }


    // close the statement
    defer rows.Close()
    // iterate over the rows
    for rows.Next() {
        var user models.ParkingDetails

        // unmarshal the row object to user
       err = rows.Scan(&user.ParkingId)

        if err != nil {
            log.Fatalf("Unable to sc an the row. %v", err)
        }

        
        // append the user in the users slice
        users = append(users, user)
        log.Println(user.ParkingId)
        UpdateParkingSlot(user.ParkingId,carId,ownerName)
    }

    if users == nil {
        log.Println("No slot available")
        slotStatus = "No slot available"
    } else {
        slotStatus = "slot allocated"
    }
 
    // return empty user on error
    return slotStatus, err
}


func UpdateParkingSlot(id int64 , carId string, ownerName string ){
    
db := createConnection()

defer db.Close()
log.Println(carId)
log.Println(ownerName)
log.Println("updating ")
allocateParkingSlotSql:=`UPDATE public."Parking" SET "CarId"=$1, "OwnerName"=$2 WHERE "PakingId"=$3;`
res, err := db.Exec(allocateParkingSlotSql, carId, ownerName,id )
if err != nil {
    log.Fatalf("Unable to execute the query. %v", err)
}

// check how many rows affected
rowsAffected, err := res.RowsAffected()

if err != nil {
    log.Fatalf("Error while checking the affected rows. %v", err)
}

fmt.Printf("Total rows/record affected %v", rowsAffected)

}


func DeallocateParking(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "PUT")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // get the userid from the request params, key is "id"
    params := mux.Vars(r)
    users,_:=DeallocateCarSpace(params["id"])
    // send all the users as response
    json.NewEncoder(w).Encode(users)
}



func DeallocateCarSpace(carId string)(string,error){
    db := createConnection()
    deallocateStatus := ""
    defer db.Close()
    allocateParkingSlotSql:=`UPDATE public."Parking" SET "CarId"='', "OwnerName"='' WHERE "CarId"=$1;`
    res,err := db.Exec(allocateParkingSlotSql, carId )

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }
    
    // check how many rows affected
    rowsAffected, err := res.RowsAffected()
    
    if err != nil {
        log.Fatalf("Error while checking the affected rows. %v", err)
    }
    if rowsAffected == 0 {
        log.Println("No car With the carId :" + carId)
        deallocateStatus="No car With the carId :" + carId
    } else {
        deallocateStatus="carId :" + carId + " left the parking slot"
    }
    fmt.Printf("Total rows/record affected %v", rowsAffected)
    return deallocateStatus , err
}