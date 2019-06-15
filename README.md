# go-swampup-service

This is a simple go program to demonstrate using the modules feature of golang.


## Prerequisites:

- latest golang version with support for go modules (1.12.6)

- Access to internet to clone the github project

- Development  machine

## Initial Preparation

1. Run `git clone https://github.com/jfrogtraining/go-swampup-service` and change to the go-swampup-service directory

2. git checkout -b <yourBranchName>

3. go build -v ./...

4. You can manually run the service by typing ./go-swampup-service

5. Test by running a GET on http://localhost:7888

6. Observe the logs

7. Interrupt the program by typing CTRL-C

## Working with Modules

1. Remove any go.mod and go.sum files

2. Run go mod init go-swampup-service

3. go mod tidy

4. Edit main.go and uncomment the lines as instructed.

5. Also comment out as instructed.

6 Now test again as before and observe the logs.

7. git add main.go

8. git commit -m "go modules test"

9. git push

10. Send PR

11 After PR is merged, follow the jenkins pipelines


## Conclusions
