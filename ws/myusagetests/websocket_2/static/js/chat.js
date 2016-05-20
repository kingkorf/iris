var connection;
var messageTxt;
var messages;

$(function(){

	if (!window["WebSocket"]) {
		window.alert("Your browser doesn't support websockets!!!")
		return
	}

	messageTxt = $("#messageTxt");
	messages = $("#messages");


	connection = new WebSocket("ws://"+HOST+"/ws");
	connection.onclose = function(evt){
		appendMessage($("<div><center><h3>Disconnected</h3></center></div>"));
	}

	connection.onmessage = function(evt) {
		appendMessage($("<div>"+evt.data+"</div>"));
	}


	$("#sendBtn").click(function(){
		connection.send(messageTxt.val());
		messageTxt.val("");
	})

})


function appendMessage(messageDiv) {
    var theDiv = messages[0]
    var doScroll = theDiv.scrollTop == theDiv.scrollHeight - theDiv.clientHeight;
    messageDiv.appendTo(messages)
    if (doScroll) {
        theDiv.scrollTop = theDiv.scrollHeight - theDiv.clientHeight;
    }
}
