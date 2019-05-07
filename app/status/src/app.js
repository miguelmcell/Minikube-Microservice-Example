const express = require('express');
const app = express();
const path = require('path');
const cors = require('cors');
const router = express.Router();
const axios = require('axios');

router.get('/',function(req,res){
  res.sendFile(path.join(__dirname+'/home.html'));
});
router.get('/getBack',function(req,res){
  axios.get("http://backend/getStatus")
        .then(function (response) {
		res.send('success'); 
        })
        .catch(function (error) {
        	res.send('fail');
	});
});
router.get('/getFront',function(req,res){
	axios.get("http://frontend/getStatus")
        .then(function (response) {
            res.send('success');
	})
        .catch(function (error) {
            res.send('fail');
	});
});

router.get('/getDatabase',function(req,res){
	axios.get("http://mongodb-service.default.svc.cluster.local:27017")
        .then(function (response) {
            res.send('success');
	})
        .catch(function (error) {
            res.send('fail');
	});
});

app.use(cors());
app.use('/', router);
app.use(express.static(__dirname ));
app.listen(process.env.port || 3000);

console.log('Running at Port 3000');
