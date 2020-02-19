# API Package
This package contains functions used to make CRUD operations to the database. 

## Helper files
In order for the api to work correctly, I created two different helper files in order to help the api functions work

### db.go
Simply makes the database available to everything inside of this package. The instance of the database connection is held within the variable ```db```, and is instantiated using the function ```ConnectDB``` inside of the ```utils``` package. To understand the different functions that ```db``` is capable of, familiarize yourself with GORM.

### errorResponse.go
This defines two different types, ```ErrorResponse``` and ```error``` to be used by the api functions.

## API Functions
These files, ```exercise.go``` and ```user.go```, are where the actual api functions are happening. They take advantage of the ```db``` variable to update the database. Each function takes in an ```http.ResponseWriter``` as ```w``` and a pointer to the ```http.Request``` as ```r``` for parameters. Inside each function, I set the HTTP header to JSON, and allow CORS from anywhere (Once we're in real development and out of the beginning pahses, we will need to avoid doing this).

The patter that each of these different functions follow is that the entity that function is dealing with is declared with var and a reference to the entity used inside of the db function. Once the result is found, then the function will return the entity encoded with JSON.

