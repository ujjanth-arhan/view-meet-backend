# view-meet-backend

 View Meet is a communication applicaiton. This this is the backend part of the code. Current Features:
 1. Accept connections from clients (including localhosts) and sends chat to all clients

### Run and Test Application

Install docker https://docs.docker.com/get-docker/<br>
(tested on Docker Desktop 4.27.2 (137060))

Clone the repository<br>
**OR**<br>
Pull the docker image
```docker
docker pull ujjanth/view-meet-backend:latest
```

Execute the following in `view-meet-backend/`
```docker
docker build -t view-meet-backend .
```

```docker
docker run --name view-meet-backend-container --rm -p 127.0.0.1:8080:8080 view-meet-backend
```

Now you have access to web socket and can test it by making appropriate request to /<br>
**OR**<br>
Download and run view-meet-frontend

Happy chatting away!
