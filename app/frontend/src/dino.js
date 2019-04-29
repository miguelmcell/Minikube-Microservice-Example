var refreshInterval = 60;
var maxScore = 1000;
var curDistance = 0;
var context;
var img;
var c;
const DINOHEIGHT = 13;

var Http = new XMLHttpRequest();
const url = 'https://jsonplaceholder.typicode.com/todos/1';

function init(){
	Http.open("GET", url);
	Http.send();
	Http.onreadystatechange=function(){
		if(this.readyState==4 && this.status==200){
			console.log(Http.responseText);
		}
	}
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
}
function getScore(){
	return sessionStorage.score;
}
function draw(){
	context.clearRect(0,0,c.width,c.height);
	context.beginPath();
	context.moveTo(0, 30);
        context.lineTo(600, 30);
        context.stroke();

        context.drawImage(img, curDistance, DINOHEIGHT,20,20);
}
function update(){
        var score = getScore();
        curDistance = (score/maxScore)*580;

        console.log(score,curDistance,maxScore);
}
function refresh(){
        update();
        draw();
        setTimeout("refresh()",refreshInterval);
}
init();
refresh();

