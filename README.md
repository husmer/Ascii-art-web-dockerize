# Dockerized-ASCII-web

## Instructions
1. Make sure you have the latest Docker installed
Official Docker [link](https://docs.docker.com/get-docker/)

3. Run these commands:
To start the program:
```bash
docker compose up
```
go to http://localhost:8080

To shut down the server:
```bash
docker compose down
```


To delete an image from your computer, you have to first delete the container.
If the container is running, you have to first stop the container.


### Example: 
To delete an image that has an active running container you would need to run.
```bash
docker container stop <container id>
docker container rm <container id>
docker image rm <image id>
```
or if you are aloof and cool who doesn't care about your other docker containers, images or files in your computer, then run:
```bash
docker container stop <container id>
docker system prune -a
```
