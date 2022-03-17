# Mock
```
git clone https://github.com/prasanjeevi/mock.git
cd mock

docker login

# Change prasanjeevi to your docker username

# Build & Run 1
docker build --build-arg port=8081 -t prasanjeevi/mock:0.1 .
docker run -p 8081:8081 -d prasanjeevi/mock:0.1
# curl http://127.0.0.1:8081 -v

# Build & Run 2
docker build --build-arg port=8082 --build-arg status=400 -t prasanjeevi/mock:0.2 .
docker run -p 8082:8082 -d prasanjeevi/mock:0.2
# curl http://127.0.0.1:8082 -v
```

Images available @ https://hub.docker.com/r/prasanjeevi/mock/tags
