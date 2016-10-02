# Welcome to User Services
This is a open source has built-in Golang that provided some basics of user management such as name, address, authenticated by Facebook.

# Reprequisite
1. Golang
2. Gin(web framework) [https://github.com/gin-gonic/gin]
3. mgo(mongodb driver) [https://labix.org/mgo]
4. mongodb
5. oauth2 [https://github.com/golang/oauth2]
6. Go cache [https://github.com/robfig/go-cache]

# Tools (optional)
1. Atom with go-plus package
2. Mongodb management studio [http://mms.litixsoft.de/index.php?lang=de/]

# Architecture
The application has been built-in Onion Architecture with:
1. API: exposes the application services that provide all functionalities of user management.
2. Domain: contains the models of application, application services.
3. Infrastructure: contains all common functions, repositories
