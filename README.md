Simple go REST API app with Go and Redis

To run the app in linux machine use bash comand:
docker compose up

Docker should be installed


!!! The app launh fiber server and preprocessing the same time. Api returns error while prepricessing not done. 
It could cause some issuse with tests. Just change 11 line in main.go 	
- go handlers.Preproces() --> handlers.Preproces()


