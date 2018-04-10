# Testing locally

this folder contains scripts and kubernetes resources to run. It will setup and
tear down all kubernetes resources required for the test. The different tests are triggered via shell scripts


## Prerequisites
To run the tests the following tools must be installed and in the path of the user executing the script.
Also a conncetion to the internet is required for helm and docker.

### Minikube and kubectl

Minikube must be installed and ready to run. Kubectl must be installed

### Helm and Tiller

Helm client must be installed and tiller pod must be running when spinning up the cluster

### vault

Local vault client must be installed

### ack

`ack` is used to parse output from kubernetes. 
For most linux distros, `grep` can be used with the additional 
`-P` flag

### base64

The tests run under osx, most linux distros use the lower case
decode flag `-d` instead of osx `-D`



## Running tests
The tests are done by shell scripts to run all tests run all.sh
in the top level test folder, which will setup and tear down
all integration tests. If a single test should be run, the setup
has also be called.