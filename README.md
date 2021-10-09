# G0-restapi-appointy
# GO-rest-Api from scratch

This Task is done for appointy intern the aim of this project is to create RestApi for minimal instagram clone 

## Database Design

![alt text](https://github.com/rohanailoni/G0-restapi-appointy/blob/main/assets/Screen%20Shot%202021-10-09%20at%209.31.24%20PM.png?raw=true=200x200)






## Running this projcet locally

### To deploy this project run


```bash
mkdir -p $HOME/go
export GOPATH="$HOME/go"

# Folder contains your golang source codes
mkdir -p $GOPATH/src

# Folder contains the binaries when you install an go based executable
mkdir -p $GOPATH/bin

# Folder contains the Go packages you install
mkdir -p $GOPATH/pkg
```
  
```bash
  mkdir -p $GOPATH/src/github.com/appiorty/
cd $GOPATH/src/github.com/appiorty
```
```bash
git clone https://github.com/rohanailoni/G0-restapi-appointy.git
```

####  Installing go packages
```bash
go.mongodb.org/mongo-driver
```

#### Compiling the porogram
```bash
go run server.go
```

## API Reference

#### Get items

```http
  GET /users/?key={$userid}
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `key`     | `JSON` | We should give **Key=userid** as the param  |


#### Get all item

```http
  GET /users/
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `user` | `JSON` | **Required**. It will return all the users in the database |

#### Posting an user details
```http
  POST /users/
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `user` | `JSON` | **Required**.if we post then we will have to post through body value|

#### Getting a post of user
```http
  `GET` /post/?userid=%{$_Id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `userid` | `string`(Id of user object) | **Required**.we can get all the post posted by the user with the id|

#### Posting raw data of postmodel to the backend(API)
```http
  `POST` /post/
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `post_data` | `JSON` | **Required**.we have to give raw data of all the post |



#### 


  
## Concurrency

concurrency is used in this project to prevent unncessary shared variable when multiple request comes to the server

at every point in the database SYNC.Mutex locks have been used


### Examples
#### userget function
![alt text](https://github.com/rohanailoni/G0-restapi-appointy/blob/main/assets/Screen%20Shot%202021-10-09%20at%209.18.45%20PM.png?raw=true)
#### userpost function
![alt text](https://github.com/rohanailoni/G0-restapi-appointy/blob/main/assets/Screen%20Shot%202021-10-09%20at%209.20.05%20PM.png?raw=true)



## ERROR Handling

At every set of point in the program error has bee set with status of the server is returend


### Examples
#### When the user is not available in the collections
![alt text](https://github.com/rohanailoni/G0-restapi-appointy/blob/main/assets/Screen%20Shot%202021-10-09%20at%209.38.46%20PM.png?raw=true)



#### when user id is not given in the GET /post/ a request is sent to user for getting the ID or parms missed
![alt text](https://github.com/rohanailoni/G0-restapi-appointy/blob/main/assets/Screen%20Shot%202021-10-09%20at%209.41.13%20PM.png?raw=true)



### TESTING
#### Get items

```http
  GET /users/?key={$userid}
```
![alt text](https://github.com/rohanailoni/G0-restapi-appointy/blob/main/assets/Screen%20Shot%202021-10-09%20at%209.46.36%20PM.png?raw=true)


#### Get all item

```http
  GET /users/
```
![alt text](https://github.com/rohanailoni/G0-restapi-appointy/blob/main/assets/Screen%20Shot%202021-10-09%20at%209.50.37%20PM.png?raw=true)

#### Posting an user details
```http
  POST /users/
```
![alt text](https://github.com/rohanailoni/G0-restapi-appointy/blob/main/assets/Screen%20Shot%202021-10-09%20at%209.49.36%20PM.png?raw=true)
#### Getting a post of user
```http
  `GET` /post/?userid=%{$_Id}
```
![alt text](https://github.com/rohanailoni/G0-restapi-appointy/blob/main/assets/Screen%20Shot%202021-10-09%20at%209.54.07%20PM.png?raw=true)
#### Posting raw data of postmodel to the backend(API)
```http
  `POST` /post/
```
![alt text](https://github.com/rohanailoni/G0-restapi-appointy/blob/main/assets/Screen%20Shot%202021-10-09%20at%209.53.03%20PM.png?raw=true)



### PASSWORD ENCRYPTION

for the encryption of password we used MD5 alogrithm which is one way algo  of 128-bit hash done by using golang default library `crypto\md5`

![alt text](https://github.com/rohanailoni/G0-restapi-appointy/blob/main/assets/Screen%20Shot%202021-10-09%20at%209.58.29%20PM.png?raw=true)
