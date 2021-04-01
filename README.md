# Producer Section
# Writes text lines to /app/messages.txt file
cd producer

# Build 'producer' image
docker build --rm -t golang-producer:v1 .

# Run 'producer' container
docker run -dit golang-producer:v1
