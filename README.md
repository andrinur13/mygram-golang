# Project 2 - Golang Hacktiv8 X Kampus Merdeka

Scalable Web Services with Golang

Kelompok : 
- Andri Nur Hidayatulloh
- Andri Kuwito

### List library : 
- [Gin Gionic](https://github.com/gin-gonic/gin) - Web Framework


### Cara Penggunaan : 

* ##### URL : https://murmuring-savannah-72759.herokuapp.com


* #### Endpoint List : 
    * ##### User : 
        * #### Register
            
            [POST]```https://murmuring-savannah-72759.herokuapp.com/users/register```
            
            body :

            ```json
            {
                "age": 21,
                "email": "andrikuwito@gmail.com",
                "password": "Bismillah",
                "username": "andrikuwito"
            }
            ```

            response
            ```json
            {
                "status": "created",
                "data": {
                    "id": 34,
                    "age": 21,
                    "email": "andrikuwito2@gmail.com",
                    "password": "$2a$04$7HbHtngP7SjMDqCaV6SFiesLdD0QLxse7HpX5a9LiiHqY9Nj2Ihyq",
                    "username": "andrikuwito2",
                    "date": "2021-12-23T13:06:37.558+07:00"
                }
            }
            ```
            

        * #### Login
            [POST]```https://murmuring-savannah-72759.herokuapp.com/users/login```
            
            body :

            ```json
            {
                "email": "andribis13@gmail.com",
                "password": "Bismillah"
            }
            ```

            response
            ```json
            {
                "status": "ok",
                "data": {
                    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZF91c2VyIjoyNX0.HWPYtCbnydJ5TXxHUc-A9my_L1NbQb1-rSlI5tPsSOc"
                }
            }
            ```

        * #### Upate User
            [PUT]```https://murmuring-savannah-72759.herokuapp.com/users```
            
            headers :

            ```
            {
                Authorization: "Bearer {{token}}"
            }
            ```

            token pada header dapat didapatkan ketika melalui proses login

            contoh :
            ```json
            {
                Authorization: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZF91c2VyIjoyNH0.XpOW4v9hpneBw9gnsVAGli_zDqj7VmMLW6ZHL80MauQ"
            }
            ```

            body :

            ```json
            {
                "age": 22
            }
            ```

            response
            ```json
            {
                "status": "ok",
                "data": {
                    "id": 24,
                    "username": "andrinur13",
                    "email": "andribis12@gmail.com",
                    "age": 90
                }
            }
            ```

        * #### Delete User
            [DELETE]```https://murmuring-savannah-72759.herokuapp.com/users```
            
            headers :

            ```
            {
                Authorization: "Bearer {{token}}"
            }
            ```

            token pada header dapat didapatkan ketika melalui proses login

            contoh :
            ```
            {
                Authorization: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZF91c2VyIjoyNH0.XpOW4v9hpneBw9gnsVAGli_zDqj7VmMLW6ZHL80MauQ"
            }
            ```            

            response
            ```json
            {
                "status": "ok",
                "data": {
                    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZF91c2VyIjoyNX0.HWPYtCbnydJ5TXxHUc-A9my_L1NbQb1-rSlI5tPsSOc"
                }
            }
            ```
    * ##### Photo : 
        * #### Add Photo
            
            [POST]```https://murmuring-savannah-72759.herokuapp.com/photos```
            
            body :

            ```json
            {
                "title": "Hacktiv",
                "caption": "Hello Hacktiv Photos ",
                "photo_url": "https://photos.hacktiv.com"
            }
            ```

            response
            ```json
            {
                "status": "created",
                "data": {
                    "id": 10,
                    "title": "Hacktiv",
                    "caption": "Hello Hacktiv Photos ",
                    "photo_url": "https://photos.hacktiv.com",
                    "user_id": 25,
                    "created_at": "0001-01-01T00:00:00Z"
                }
            }
            ```
        * #### Delete Photo
            
            [Delete]```https://murmuring-savannah-72759.herokuapp.com/photos/:id```

            example : 
            ```
            https://murmuring-savannah-72759.herokuapp.com/photos/8
            ```

            headers :

            ```
            {
                Authorization: "Bearer {{token}}"
            }
            ```

            token pada header dapat didapatkan ketika melalui proses login

            contoh :
            ```
            {
                Authorization: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZF91c2VyIjoyNH0.XpOW4v9hpneBw9gnsVAGli_zDqj7VmMLW6ZHL80MauQ"
            }
            ```            

            response
            ```json
            {
                "status": "ok",
                "data": "success deleted photo!"
            }
            ```
          
        * #### Get Photo
            
            [GET]```https://murmuring-savannah-72759.herokuapp.com/photos```

            headers :

            ```
            {
                Authorization: "Bearer {{token}}"
            }
            ```

            token pada header dapat didapatkan ketika melalui proses login

            contoh :
            ```
            {
                Authorization: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZF91c2VyIjoyNH0.XpOW4v9hpneBw9gnsVAGli_zDqj7VmMLW6ZHL80MauQ"
            }
            ```            

            response
            ```json
            {
                "status": "ok",
                "data": [
                    {
                        "id": 9,
                        "title": "Judul Apa ya 3",
                        "caption": "Hello 3",
                        "photo_url": "https://facebook.com",
                        "created_at": "2021-12-16T22:34:17+07:00",
                        "comments": [
                            {
                                "id": 2,
                                "message": "Update comments update",
                                "photo_id": 9,
                                "user_id": 25
                            },
                            {
                                "id": 3,
                                "message": "kira2 bagus gaa yaa ?",
                                "photo_id": 9,
                                "user_id": 25
                            }
                        ]
                    },
                    {
                        "id": 10,
                        "title": "Hacktiv",
                        "caption": "Hello Hacktiv Photos ",
                        "photo_url": "https://photos.hacktiv.com",
                        "created_at": "2021-12-23T06:31:39+07:00",
                        "comments": []
                    }
                ]
            }                           
            ```
          
        * #### Get Photo ID
            
            [GET]```https://murmuring-savannah-72759.herokuapp.com/photos/:id```

            example : 
            ```
            https://murmuring-savannah-72759.herokuapp.com/photos/9
            ```
           

            response
            ```json
            {
                "status": "ok",
                "data": {
                    "id": 9,
                    "title": "Judul Apa ya 3",
                    "caption": "Hello 3",
                    "photo_url": "https://facebook.com",
                    "created_at": "2021-12-16T22:34:17+07:00",
                    "user": {
                        "id": 25,
                        "username": "andrinur19",
                        "email": "andribis13@gmail.com"
                    },
                    "comments": [
                        {
                            "id": 2,
                            "message": "Update comments update",
                            "photo_id": 9,
                            "user_id": 25
                        },
                        {
                            "id": 3,
                            "message": "kira2 bagus gaa yaa ?",
                            "photo_id": 9,
                            "user_id": 25
                        }
                    ]
                }
            }
            ```
          

        * #### Update Photo
            
            [GET]```https://murmuring-savannah-72759.herokuapp.com/photos```

            headers :

            ```
            {
                Authorization: "Bearer {{token}}"
            }
            ```

            token pada header dapat didapatkan ketika melalui proses login

            contoh :
            ```
            {
                Authorization: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZF91c2VyIjoyNH0.XpOW4v9hpneBw9gnsVAGli_zDqj7VmMLW6ZHL80MauQ"
            }
            ```            

            body
            ```json
            {
                "title": "caption update",
                "caption": "caption update",
                "photo_url": "https://images.google.com"
            }
            ```

            response
            ```json
            {
                "status": "ok",
                "data": {
                    "id": 9,
                    "title": "caption update",
                    "caption": "caption update",
                    "photo_url": "https://images.google.com",
                    "user_id": 25
                }
            }
            ```
          