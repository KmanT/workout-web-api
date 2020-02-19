# Model Package
This package contains the models that the database entities are based off of as well as what our API functions are wrapped around. Inside there are four types: user, exercise, exception and token. The user and exercise are the actual models for the database and the API while the other two help the API out.

## Exception
This is just for when the API encounters an error, and it will return the error as a message.

## Exercise & User
These two structures are annotated with 'gorm.Model.' This means that when GORM is updating the schema, it will look at these structures and their gorm annotations to make appropriate changes. Just a note, GORM will *not* delete data if an attribute is removed from any of the structures.

Inside of the backtick annotations, there are two labels: 'gorm' and 'json'. The gorm annotation determines what type of field the attribute will be inside of the database, while the json tag annotates how the structure will be when the structure as JSON will be returned.

## Token

This is what is used for authentication during runtime. When the user logins or signs up for the service, a token will be returned for the user to use for further API calls. The token must be stored as a cookie in the frontend somewhere in order for the authentication to work.