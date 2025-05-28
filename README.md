# quote-book
Simple REST-api with stroing and managing quotes.

# API endpoints

Get -
    Receiving all possible quotes
    localhost:8080/quotes
    test:  Invoke-RestMethod -Uri "http://localhost:8080/quotes" -Method Get

Post request
    Creating a new quote
    localhost:8080/quotes
    test: Invoke-RestMethod -Uri "http://localhost:8080/quotes" -Method POST -Body '{"author":"Eleanor Roosevelt","quote":"The future belongs to those who believe in the beauty of their dreams."}' -ContentType "application/json"

Get - 
    Receiving quotes filtered by author
    localhost:8080/quotes?author=Confucius
    test: Invoke-RestMethod -Uri "http://localhost:8080/quotes?author=Confucius" -Method Get

Get - 
    Receiving random quote 
    localhost:8080/quotes/random
    test: Invoke-RestMethod -Uri "http://localhost:8080/quotes/random" -Method Get


Delete -
    Removing a quote with id
    localhost:8080/quotes/{id}
    test: Invoke-RestMethod -Uri "http://localhost:8080/quotes/2" -Method Delete

# Starting application 

There is used migration due to save and store records in the db. Also there is used 2 external libraries 'lib/pq' and 'gorilla/mux'. 
Before starting any application pay attention on config.json, enter postgres password and username, There is blank space due to security reasons.
To use this application you need to have installed docker. We can use Makefile, if it is not working in your machine, enter commands manually:
"docker-compose up -d" to start docker compose. 
After that download external libraries with the commands:
"go install github.com/lib/pq"
"go install github.com/gorilla/mux". 
Then go to the project "./quotes" and start the application with the command "go run cmd/main.go". 
All information about the ports and setup in the config file.

I wish everything is gonna work, thank you! 

# Working reports
Also you can pay attention in working flow of the application, which was represented in my localmachine 
in the form of screenshots.
