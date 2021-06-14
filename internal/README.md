# internal

The internal directory mainly contains codes used to implement basic logic in the service, such as some CRUD interfaces for operating databases. The server and client of the microservice are 

placed in pkg/microservice and do not directly perform database operations, so you need to call the internal package to perform database operations.

