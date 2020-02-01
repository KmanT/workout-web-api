# Workout Web Api
**READ THE ENTIRE THING!**

Hello, everyone, this is the source code to the API I build for my work out app. I'll be walking through each package and defining what they're used for and why they're there. Before we get started, feel more than welcome to get a hold of me if you have any sort of questions if you need any sort of clarification.

## Getting started

There will be a few things you will need to have and understand before you run the program

### What you will need

Before you run the program, you're going to need:

- The Go run time, obviously
- I suggest Visual Studio Code, but you are welcome to use which ever IDE you prefer. I like VSCode for it's integrated terminal and the customization through the extensions. (To open, just press ``` ` + CTRL ```)
- PostgreSQL running on your computer (I also reccomend pgAdmin to view whether or not the data is being updated).
- Dep, which is used to manage the *dep*endencies.
- Postman, which is an HTTP Client that can interact with the API. I will provide possible inputs in order to get what we want.
- This repository must be in your GOPATH! It will not run otherwise. So place your clone of this repository inside:

```C:\Users\[YOU]\go\src```

- Outlined later in this README, proper configurations. Go to the "Configuration" section, then come back here.

Once you have everything all set, try to run this inside your integrated terminal:

```dep init```

then:

```go build main.go```

If there is anything wrong, it will let you know. Otherwise, the program is ready to run! Before that, you will need to understand what each of the packages are for.


### The packages

In go, source code is organized into folders, or packages. Each time you import. For instance, if you import the ```models``` package, you will be importing the contents of exception.go, exercise.go, token.go, and then user.go. Anyways, I'll be writing README's for each of these different packages.

The packages that I have made are:
- ```api```: Everything that pertains with communication between the database and the api.
- ```config```: Sets up the variables for the server to run and the database to be connected
- ```model```: A programatic representation of each of the entities that will be inside the database (as well as database authentication through JWT, more on that later).
- ```routes```: Defines the http routes in order to run the api funcitons properly.
- ```utils```: Has a function to connect to the database, then a subpackage for authentication.
- vendor: These contain external packages such as gorilla/mux, jinzhu/gorm, and dgrijalva/jwt-go.


### Configuration

Before the program will actually start, you will need to go to your ```config``` package to match it with your system.

So go your ```config``` file and create ```dev.config.json```

Inside there please enter the following:

```json
{
    "server": {
        "host": "localhost",
        "port": "8000"
    },
    "database": {
        "host": "127.0.0.1",
        "port": "5432",
        "username": "postgres",
        "dbname": "WorkoutApp",
        "password": "[YOURDATABASEPASSWORD]"
    }
}
```

This is to allow the server to run locally on your computer. If you would like to try to deploy this yourself with a free Heroku Dyno create this inside of ```heroku.config.json```:

```json
{
    "server": {
        "host": "localhost",
        "port": "8000"
    },
    "database": {
        "host": "[heroku's provided host]",
        "port": "5432",
        "username": "[heroku's username]",
        "dbname": "[the name of your heroku postgres db]",
        "password": "[the password to the heroku postgres db]"
    }
}
```

That information is provided for you once you set up an instance of Heroku Postgres. You **do not** have to do this in order to run this locally. Once you have everything set up correctly go ahead and run this inside your intergrated terminal:

```go run main.go```

or this if you have the ```main.exe``` built:

```./main```