# AnyPortInAStorm

A coding exercise to look at RPC & Golang MicroServices. 


To Run the Unit Tests

```./test.sh```

To run through docker, first build all the images. 

```./buildall.sh```

Once you have built the images 

```docker compose up```

To run locally, there are 2 scripts to run using "go runs". This uses a local database not 
the one defined in the docker compose file.  To run the code like this you'll need the database which is all defined in the /db/init.sql.

```runServer.sh```

and 

```runClient.sh```


### To access the database,  

```
docker exec -it anyportinastorm_mysql_1 mysql -u sqluser -psqlpass
```

### To upload a file,  you can use browser to http://localhost:8080 and use the form.  

Or use curl

``` 
curl -F "portsFile=@ports.json" http://localhost:8080/upload
```

## Random Commands to make a Note off 

Generate the RPC Code.  
```
protoc --go_out=. --go-grpc_out=. PortController.proto
```

Log into the Docker Compose Database
```
docker exec -it anyportinastorm_mysql_1 mysql -u sqluser -psqlpass
```


### Issues or things I am not happy with.

* Memory Management is really poor
* The JSON for Cordinates, since slice orders are not guaranteed in Go, the Lat Long could be the wrong way around. 
* Limited Database Connection pooling.  (Open and Close for Each Save Operation)
* Seperate this project into 3 repo's   PortController Defination, Client and Server
* Better Networking with a service mesh.   (K8's is better here)
* Need to Use something like Vault for the Creds
* if there are No Cords,  the database stores -1.

### TODO's 

* Use proper make files for building and running
* Implement Signals for the OS and Clean up nicely.
* Clean up the JSON import implement some streaming/chunking of the file. Fix the memory has well.  (better use of pointers)