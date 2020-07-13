# !/bin/bash

function startMicroservice {
sudo docker build -t go-docker .
sudo docker run -d -p 5000:5000 go-docker
echo "microservice started succesfully"
}

function stopMicroservice { 
x=`sudo docker container ls|sed -n '2p'|{ read first rest ; echo $first ; }`
sudo docker container stop $x
echo "microservice stopped succesfully"
}
function statusMicroservice {
sudo docker container ls
}

echo "Enter Choice :"
echo "1. Start microservice"
echo "2. Stop microservice"
echo "3. check Status of microservice"
read ch 
  
# Switch Case to perform 
# calulator operations 
case $ch in
  1)startMicroservice
  ;; 
  2)stopMicroservice 
  ;; 
  3)statusMicroservice 
  ;; 
  *)
    echo -n "wrong choice"
esac
