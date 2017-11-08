var list = document.getElementById("conversation");
var btn = document.getElementById("btn");

btn.addEventListener("click", function(){    
    const userInput = document.getElementById("user-input");
    const question = userInput.value;
    document.getElementById("user-input").value = "";
    const request = new XMLHttpRequest();
    const params = "question=" + question;
    request.open("POST", "http://localhost:8080/ask", true);
    request.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    request.onreadystatechange = function(){
        if(request.readyState === XMLHttpRequest.DONE) {
            var response = request.responseText;
            addListItem("user_message", question);
            setTimeout(function(){
                addListItem("eliza_message", response);
            }, 2000); // add a 2 second delay to make it look like Eliza is "typing" her message. instead of instantly displaying it.
        }
    }
    request.send(params);
});

function addListItem(speaker, text){
    const htmlString = "<li class=\"list-group-item " + speaker + "\"><p align=\"left\">" + text + "</p></li>"
    list.insertAdjacentHTML("beforeend", htmlString);
}

