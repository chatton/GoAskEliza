


var list = document.getElementById("conversation");
var btn = document.getElementById("btn");


btn.addEventListener("click", function(){
    var userInput = document.getElementById("user-input");
    var question = userInput.value;
    document.getElementById("user-input").value = "";
    var request = new XMLHttpRequest();
    var params = "question=" + question;
    request.open("POST", "http://localhost:8080/ask", true);
    request.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    request.onreadystatechange = function(){
        if(request.readyState === XMLHttpRequest.DONE) {
            var response = request.responseText;
            addListItem("user_message", question);
            addListItem("eliza_message", response);
        }
    }
    request.send(params);
});


function addListItem(speaker, text){
    var htmlString = "<li class=\"list-group-item " + speaker + "\"><p align=\"left\">" + text + "</p></li>"
    list.insertAdjacentHTML("beforeend", htmlString);
}

