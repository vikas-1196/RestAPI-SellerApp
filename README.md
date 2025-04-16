#RestFul Api to Get details of Employee
# For Database connection :
* first create a database in mysql with name "EmployeeDatabase"
* tUser Name password test1 of MySQL should be root:root 
Database connection Established

# Api to fetch Employee Details 
Get Request : http://localhost:9000/test/v1/a/getEmployee 
POST : http://localhost:9000/test/v1/a/createEmp
PUT: http://localhost:9000/test/v1/a/updateEmp/{id}
DELETE: http://localhost:9000/test/v1/a/deleteEmp/{id}

* For post request :
sample json data:
{
    "Id":"1",
    "Name":"test1"
}

# DockerFile to containerized RESTAPIs
* First build this dockerfile: docker build -t restapi .
