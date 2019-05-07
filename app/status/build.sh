sudo docker build -t status:latest .
sudo docker tag status localhost:5000/status
sudo docker push localhost:5000/status
