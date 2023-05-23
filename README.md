PREREQUISISTES

In order to run, this project requires the following:
1. GO installed
2. A MySQL Database server running locally


RUNNING THE APPLICATION

1. Start MySQL
2. Run the MySQL commands contained in “Database commands.txt” to create the test database
3. Start the application in interpreted mode by running “go run cushonTest.go” or compile the application (“go build cushonTest.go”) and run the generated application directly.

The application is now listening for incoming http requests on localhost:8080. It can be tested by sending requests to :8080/funds and :8080/deposits with the appropriate methods and payloads (see DESCRPTION OF SERVICE for a more detailed explanation) using any suitable utility (e.g., Postman). Alternatively, the html file frontend.html has been provided as a very basic illustration of how a web front end would interact with the backend service; this file can be opened with a browser and interacted with manually.


DESCRIPTION OF SERVICE

The service consists of a Go application listening for incoming GET requests on :8080/funds and incoming POST requests on :8080/deposits. Additionally, it interacts with a MySQL database with the following tables:

cushon_test.funds
cushon_test.deposits
cushon_test.customers

funds contains a list of all the funds available to the customer, deposits records all the amounts invested into the respective funds, and customers contains a list of the customers along with their authentication keys. 

The SQL statements in “Database commands.txt” will create the test database, the tables and populate funds and customers with test data; it is expected that a “real world” implementation would involve the list of funds being managed by the relevant team in Cushon and that the customers table would be updated dynamically (with new customers signing up on the main website having new entries added to this table, as well as authentication keys being generated and added to this table).

GET requests to :8080/funds triggers the application to respond with a json payload of all the available funds and their IDs.

POST requests to :8080/deposits with a json payload of the following form:

{"customerID": 1, "deposits": [{"fund": 1, "amount": 4000}]}

and presenting the correct authentic token as a “Authorisation” header, triggers the application to insert a new row into the deposits table recording the id of the customer, the id of the fund, the amount invested, and the current datetime. 

NOTE; the application can accommodate an arbitrary number of “deposits” objects in one payload and adds them accordingly, for example, the following json:

{"customerID": 100, "deposits": [{"fund": 1, "amount": 1000}, {"fund": 2, "amount": 3000}]

Would record two entries simultaneously, one recording £1000 entered into fund 1 and the other £3000 into fund 2, both having the same CreatedAt value and CustomerID.

FRONTEND

The file frontend.html has been provided to illustrate the manner in which the application front end would interact with this service; it must be stressed that it is illustrative only as the file as presented contains serious security flaws (hard-coded and user-visible customer ids and authentication tokens), has no style information and uses primitive html and JavaScript features (including the visually ghastly html button).

It is expected that the “Real” frontend would have a fully-featured UI developed in React as an extension to Cushon’s existing estate and would be powered by a web application framework (such as NodeJS or Angular); tasks such as login, session management and retrieving customer IDs and authentication keys would be performed by the middleware and segregated from the front-end. HTTPS would also be enabled to ensure client-server encryption.

EXAMPLE INTERACTION
When the page (or component) loads, it sends a GET request to localhost:8080/funds, retrieving an up to date list of the available funds. It uses this data to render the UI (in this simple example, add options to a Select element). The user can now select the appropriate fund from the list, inputs their desired investment amount and submits their information. The frontend sends a POST request to localhost:8080/deposits. The application processes the payload, authenticates it and then attempts to insert new records. It sends a success or failure message back to the frontend accordingly. The frontend handles this message accordingly (in this simple example triggering an alert, but in the “Real” implementation triggering a page redirect or a component reload to indicate success or failure).


FUTURE ENHANCEMENTS
Here is a list of potential enhancements that would be considered in a real-world scenario (this list is not exhaustive but is intended to show some key considerations):

1. Deploying the application behind a webserver (e.g. Nginx) and enabling port-forwarding
2. Enabling https on the web server to ensure encryption
3. Deploying the application within a container (e.g. Dockerising)
4. Using limited-period session tokens (e.g. JWT) for authentication and authorisation rather than static header keys
5. Validating that the fund IDs being deposited into exist and are available, that deposit amounts are within valid bounds and that the customer has sufficient balance
6. Creating an API/extending this one for the user to query and update their deposit records
7. Adding more detail to table schemas as required (e.g. Fund Availability Date, Customer Creation Date, Customer Type, etc.)
8. Adding more data to the API responses (e.g., for the GET requests, Fund Type, creation date, rates, maximum/minimum investment amounts, for the POST requests, number of records, total amount added, total amount currently invested by user, total number of funds the user is invested in etc.)
9. CI/CD pipeline deployment, logging and backups.
