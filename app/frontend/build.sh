sudo docker build -t frontend:latest .
sudo docker tag frontend localhost:5000/frontend
sudo docker push localhost:5000/frontend
