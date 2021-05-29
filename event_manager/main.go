package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/gocql/gocql"
	"time"
)



func main(){

	type Event struct {


		Eventname  string
		Event_creationTimestamp string
		Kind string
		Namespace string
		Name string
		Message string
		Reason string 
		FirstTimestamp string 
		LastTimestamp string
		Typename string
	
	
	}


   cluster := gocql.NewCluster("10.157.50.15:9042") //replace PublicIP with the IP addresses used by your cluster.
   cluster.Consistency = gocql.Quorum
   cluster.ProtoVersion = 4
   cluster.ConnectTimeout = time.Second * 10
   //cluster.Authenticator = gocql.PasswordAuthenticator{Username: "Username", Password: "Password"} //replace the username and password fields with their real settings.
   session, err := cluster.CreateSession()

   if err != nil {
	fmt.Println(err)
	return
   }
   defer session.Close()
   
   // create keyspaces
   err = session.Query("CREATE KEYSPACE IF NOT EXISTS events_keyspace WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1};").Exec()

   if err != nil {
   fmt.Println(err)
   return
   }
   
   // create table
   err = session.Query("CREATE TABLE IF NOT EXISTS events_keyspace.events_table (eventname text, event_creationTimestamp date, kind text,namespace text,name text,message text,reason text,firstTimestamp date,lastTimestamp date,type text,  PRIMARY KEY (event_creationTimestamp, eventname));").Exec()
   if err != nil {
   fmt.Println(err)
   return
   }

  
   resp, err := http.Get("http://127.0.0.1:5000/events")
   if err != nil {
	fmt.Println(err)
   }
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
	fmt.Println(err)
   }
   var events [] Event
   json.Unmarshal(body,&events)

   for _, event := range events {


	DBquery := fmt.Sprintf("INSERT INTO events_keyspace.events_table (eventname,event_creationTimestamp,kind,namespace,name,message,reason,firstTimestamp,lastTimestamp,type) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s');", event.Eventname, event.Event_creationTimestamp, event.Kind,event.Namespace, event.Name, event.Message, event.Reason, event.FirstTimestamp, event.LastTimestamp, event.Typename)
	fmt.Println(DBquery)
	err = session.Query(DBquery).Exec()
	if err != nil {
		fmt.Println(err)
	   }
    fmt.Println("******************************************")
   
}

}