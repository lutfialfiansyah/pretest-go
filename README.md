## How to Install

Do the following command
```
mkdir -p yourpath/pretest
cd yourpath/pretest
git clone https://github.com/lutfialfiansyah/pretest.git
```

## Go to your project folder

Follow the following command via your cmd / terminal

```
cd pretest/
```
And you should now find the source code under `pretest` directory.

```
  + yourpath/
  |
  +--+ src/
  |  |
  |  +--+ github.com/
  |     |
  |     +--+ pretest/
  |         |
  |            + app/
  |            + controller/
  |            + database/
  |         +--+ main.go

```

## How to Run
```
go run main.go
```

## How to Test APP
There are 3 end points :
- Post/Send a message
```
curl -X POST http://localhost:8080/post/yourmessagesend
```
- Get messages
```
curl -X GET http://localhost:8080/get-message
```
- Get messages with long life connection

You can use this [tool](https://chrome.google.com/webstore/detail/websocket-test-client/fgponpodhbmadfljofbimhhlengambbn) for test

## Project Status
For end point ```/all-message```