# Employee REST API
This program has 5 endpoints that interact with the employee table (migrated in this program, defined in database\migrations\create_table.up.sql ), and contains the following endpoints:

1. GET /employees:
- This function selects all the data in the employees table, and returns it in the form of a JSON array.

2. GET /employees/{id}:
- This endpoint retrieves the employee data by its ID, which is specified in the request URL.

3. POST /employees:
- This endpoint creates a new employee record in the database.

4. PUT /employees/{id}:
- This endpoint updates an existing employee record based on its ID

5. DELETE /employees/{id}:
- Deletes an employee record based on its ID

## Executing the program
1. Clone the repository
2. execute the following script : docker-compose -f docker_compose.yml up
    - The script above would start the docker_compose.yml file, set up the DB and API, and using the Dockerfile, run the program by executing the main.go file
    - The current setup uses default values for the DB credentials and port. Update them as per your system requirements

## Documentation
- A simple documentation for this API can be found in the OpenAPI_Documentation.yaml file. Do note that the url used in the documentation is still in the localhost:8080 environment. Adjust if necessary