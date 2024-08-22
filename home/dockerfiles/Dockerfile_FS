FROM golang:1.23

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
# for file servere volume mount-
    # requirement - volume mount
    # network connect with DB server
    # network connect with Webapp server


    # what i need to do tomorrow is below
        # Create container of Web app for resumemanagement with public IP so that others can access my web.
        # Create container of DB server with volume mount so that the data can be kept persistantly.
            # Before removing the previous server you need to backup the data saved under guideline officially provided by postgre! 
        # Create container of File server with volumue mount/network info and port since  you need to send img and file for users! 
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]


