```
docker container run --name fluentdserver --rm -it \ 
  -p 24224:24224 \
  -v $(pwd)/docker.conf:/fluentd/etc/docker.conf \
  -e FLUENTD_CONF=docker.conf \
  fluent/fluentd:latest
  ```

  ```
  docker build -t go-logger:0.1 .
  
   docker run -it -d -p 8080:8080 --log-driver=fluentd --log-opt tag="go-logger:0.1" --name gogo go-logger:0.1
  ```