# Containerized Producer-Consumer using volume persistent storage
This is a simple producer/consumer utility
One, the producer produces text to a certain file and
the consumer consumes the text from the same file.

# Producer Section

# Build 'producer' image
docker build --rm -t golang-producer:v1 .

# Run 'producer' container
docker run -dit golang-producer:v1

# Consumer Section

# Build 'consumer' image
docker build --rm -t golang-consumer:v1 .

# Run 'consumer' container
docker run -dit golang-consumer:v1
