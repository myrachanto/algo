package main

func main(){
	
}


// docker network create ubuntu-test

// docker run -td --name ubuntu1 --network ubuntu-test ubuntu:20.04
// docker run -td --name ubuntu2 --network ubuntu-test ubuntu:20.04
// docker exec ubuntu1 bash -c "apt update; apt install -y iputils-ping; ping -c 5 ubuntu2"

// docker rm -f ubuntu1
// docker rm -f ubuntu2

// docker run -td --name ubuntu1 ubuntu:20.04
// docker run -td --name ubuntu2 ubuntu:20.04
// docker exec ubuntu1 bash -c "apt update; apt install -y iputils-ping; ping -c 5 ubuntu2"

// docker rm -f ubuntu1
// docker rm -f ubuntu2

// docker run -td --name ubuntu1 --network host ubuntu:20.04
// docker run -td --name ubuntu2 --network host ubuntu:20.04
// docker exec ubuntu1 bash -c "apt update; apt install -y iputils-ping; ping -c 5 ubuntu2"

> How to bind multiple ports from a container onto the host?

> How to bind multiple consecutive ports from a container onto the host?

> Can you bind both TCP and UDP ports? If yes, how do you bind TCP and UDP ports?

> How to bind on all interfaces or only on some of them?

> How do you check which ports are bindable in a Docker Image?