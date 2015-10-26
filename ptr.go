package plaintextrefresher

import (
	"fmt"
	"net/http"
)

func Handle(url string) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "text/html")
		fmt.Fprint(rw, fmt.Sprintf(html, url))
	}
}

const html = `<!DOCTYPE html>
<html>
<body>
	<div id="title">
		<input type="button" id="playpause"></input>
		<output id="amount" for="slider">0</output>
		<input type="range" id="slider" oninput="amount.value=slider.value+' ms'"/>
	</div>
	<div>
		<pre id="block">
		</pre>
	</div>
</body>
<script>
	var button = document.getElementById('playpause');
	var slider = document.getElementById('slider');
	var amount = document.getElementById('amount');
	slider.max = 5000;
	slider.min = 100;
	slider.value = 500;
	amount.value = slider.value+" ms";
	var playing = false;
	var update = function(){
		var oReq = new XMLHttpRequest();
		oReq.addEventListener("load", function() {
			document.getElementById('block').innerHTML = this.responseText;
		});
		oReq.open("GET", "%s");
		oReq.send();
		if(playing){
			setTimeout(function(){
				update();
			},slider.value);
		}
	}
	var stopPlaying = function(){
		playing =  false;
		button.value = "play every";
	};
	var startPlaying = function(){
		button.value = "pause";
		playing = true;
		update();
	}
	button.addEventListener('click', function() {
		if(playing){
			stopPlaying();
		} else {
			startPlaying();
		}
	}, false);
	startPlaying();
</script>
</html>
`
