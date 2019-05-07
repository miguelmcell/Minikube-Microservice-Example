const express = require('express');
const app = express();
const path = require('path');
const cors = require('cors');
const router = express.Router();
const axios = require('axios');

var scores;
/*
axios.get("http://backend/getPlayers")
	.then(function (response) {
	    scores = response.data;
	    //console.log(response.data);
  	})
  	.catch(function (error) {
    	    console.log(error);
	});*/
router.get('/',function(req,res){
  res.sendFile(path.join(__dirname+'/index.html'));
});
router.get('/getPlayers',function(req,res){
  axios.get("http://backend/getPlayers")
        .then(function (response) {
            scores = response.data;
	    res.send(scores);
            //console.log(response.data);
        })
        .catch(function (error) {
            console.log(error);
        });
});
router.get('/getStatus', function(req,res){
  res.send('success'); 
});
router.get('/addPlayer',function(req,res){
  console.log("name: " + req.param('name'));
  console.log("score " + req.param('score'));
  axios.post("http://backend/postPlayer/"+req.param('name')+"_"+req.param('score'))
        .then(function (response) {
            console.log(response.data);
        })
        .catch(function (error) {
            console.log(error);
        });
  res.send('success');
});

app.use(cors());
app.use('/', router);
app.use(express.static(__dirname ));
app.listen(process.env.port || 3000);

console.log('Running at Port 3000');
