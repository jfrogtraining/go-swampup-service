# go-swampup-service

This is a simple go program to demonstrate using the modules feature of golang.


## Prerequisites:

- Latest golang version with support for go modules (1.12.6)

- Access to internet to clone the github project

- Development  machine

## Initial Preparation

1. From your browser, open the url `https://github.com/jfrogtraining/go-swampup-service` and fork the repository and copy the clone url.

2. Change to your working directory and clone using the forked repo url

3. cd go-swampup-service

4. git checkout -b <yourBranchName>

5. go build -v ./...

6. You can manually run the service by typing ./go-swampup-service

7. Test by running a GET on http://localhost:7888

6. Observe the logs

8. Interrupt the program by typing CTRL-C

## Working with Modules

1. Remove any go.mod and go.sum files and ensure GO111MODULE environment variable is set to "on" and set GOPROXY to https://gocenter.io

2. Run go mod init go-swampup-service

3. go mod tidy

4. Edit main.go and uncomment the lines as instructed.

5. Also comment out as instructed.

6 Now test again as before and observe the logs.

7. git add main.go; git add go.mod; git add go.sum

8. git commit -m "go modules test"

9. git push
If you see an error indicating there is no upstream set, just execute the command displayed.
e.g git push --set-upstream origin <yourBranchName>

10. Send PR

11 After PR is merged, follow the jenkins pipelines


## Conclusions

We imported logrus module from GoCenter and used it in our example. We saw the dependency on logrus module in go.mod
