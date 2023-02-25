Simple go REST API app with Go and Redis

<b>Run the app on linux machine</b>
Use bash comand:
docker compose up

<b>Run tests on linux machine</b>
Use bash comand:
docker-compose -f docker-compose-test.yml



!!! The app launh fiber server and preprocessing the same time. Api returns error while prepricessing not done. 
It could cause some issuse with tests. Just change 11 line in main.go 	
- go handlers.Preproces() --> handlers.Preproces()


