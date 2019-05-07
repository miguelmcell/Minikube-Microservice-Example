var refreshInterval = 60;
var maxScore;
var curDistance = 0;
var context;
var img;
var c;
const DINOHEIGHT = 13;
var xhttp = new XMLHttpRequest();
var players;
var score = 0;

function init(){
	c = document.getElementById("myCanvas");
	context = c.getContext("2d");
	context.translate(0.5,0.5);
	img = new Image();
	img.onload = function () {
		context.drawImage(img, curDistance, DINOHEIGHT,20,20);
	}
	img.src = "assets/trex.png";
	if(sessionStorage.score != 0){
		sessionStorage.score = 0;
	}
	xhttp.onreadystatechange = function() {
		if (this.readyState == 4 && this.status == 200) {
			players = JSON.parse(xhttp.responseText);
			console.log(players);
			maxScore = players[0].score;
		}
	};
	xhttp.open("GET", "/getPlayers", true);
	xhttp.send();
}
function getScore(){
	return sessionStorage.score;
}
function draw(){
	if(typeof maxScore !== 'undefined'){
		context.clearRect(0,0,c.width,c.height);
		context.beginPath();
		context.moveTo(0, 30);
		context.lineTo(600, 30);
		context.stroke();

		if(typeof players !== 'undefined'){
			players.forEach(function (p){
					context.drawImage(img,(p.score/maxScore)*580, DINOHEIGHT, 20,20);
					});
		}

		context.drawImage(img, curDistance, DINOHEIGHT,20,20);
	}
}
function update(){
	score = getScore();
	curDistance = (score/maxScore)*580;
	if(score > maxScore){
		maxScore = score;
	}
	
}
function refresh(){
	update();
	draw();
	setTimeout("refresh()",refreshInterval);
}
init();
refresh();

