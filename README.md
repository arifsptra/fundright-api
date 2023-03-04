
# Fundright App API

This API Project is a server application created using the Golang programming language, utilizing the Gin framework to implement the API, the Gorm framework to implement Database (MySQL), and integrate with the Midtrans online payment platform.

## How to Run

To run this project, make sure you have Golang installed on your computer. Then do the following steps:

    # clone repository
    $ git clone https://github.com/arifsptra/fundright-api.git
    $ cd fundright-api
    
    # run project
    $ go run main.go

The API server will run on [http://localhost:8080](http://localhost:8080/)

## Database Structure
![fundright](https://user-images.githubusercontent.com/91882024/222872345-f93cf3b4-632b-47aa-bb06-c2f337170323.png)

## API Endpoint

Following are some of the API endpoints that can be accessed from this application:

1. POST /users - for register
2. POST /session - for login
3. POST /email_checkers - for email checker
4. POST /avatars - for upload avatar
5. POST /users/fetch - for fetch user
6. GET /campaigns - for get all campaign
7. GET /campaigns/:id - for get campaign by id
8. POST /campaigns - for create new campaign
9. PUT /campaigns/:id - for update campaign
10. POST /campaign-images - for post campaign image
11. GET /campaigns/:id/transactions - for get transaction campaign
12. GET /transactions - for get user transaction
13. POST /transactions - for create transaction
14. POST /transaction/notification - for create transaction notification

## Contribution

If you want to contribute to this project, please do a pull request on this repository. I would really appreciate your help in developing this app.

## Licence

This project is licensed under the MIT license. Please see the LICENSE file for more information.
