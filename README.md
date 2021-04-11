# Containerized Producer-Consumer using volume persistent storage
This is a simple producer/consumer utility
One, the producer produces text to a certain file and
the consumer consumes the text from the same file.

# Create a volume
```docker volume create messagesVolume```
Where `messageVolume` is the name of our volume

# Producer Section

# Build 'producer' image
```docker build --rm -f Dockerfile.producer -t golang-producer:v1 .```

# Run 'producer' container with mounted volume
```docker run -dit --name="golang-producer" --mount source=messagesVolume,destination=/app golang-producer:v1```
Where
 `messageVolume` is the name of our volume
 `/app` is the folder inside the container we want to mount onto

# Enter container and produce messages using the producer utility
```docker exec -it golang-producer /bin/sh```

Execute the producer
```./producer```

# Consumer Section

# Build 'consumer' image
```docker build --rm -f Dockerfile.consumer -t golang-consumer:v1 .```

# Run 'consumer' container with mounted volume
```docker run -dit --name="golang-consumer" --mount source=messagesVolume,destination=/app golang-consumer:v1```

# Enter container and read messages using the consumer utility
```docker exec -it golang-consumer /bin/sh```

Execute the consumer
```./consumer```

# Volume
----------------------
The volume content can be seen in the following directory:
`/var/lib/docker/volumes/messagesVolume/_data/`

## List volumes
```docker volume ls```

## Inspect volume
```docker volume inspect <volumename>```

## Removing volume
```docker volume rm <volumename>```


* Note that the producer/consumer utilities were not the main objective of this example.
* Hence, they can probably be optimized
