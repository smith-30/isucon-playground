FROM golang:1.21.1-bullseye

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Install air for hot reload
RUN go install github.com/cosmtrek/air@v1.49.0

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the air tool for hot reloading
CMD ["air", "-c", ".air.toml"]
