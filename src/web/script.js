var list = document.getElementById("conversation");
var btn = document.getElementById("btn");

keyCodes = {
    ENTER : 13
}

document.addEventListener("keypress", function(e){
    if(e.keyCode != keyCodes.ENTER){
        return; // we want to ignore all keypresses other than enter.
    }

    const userInput = document.getElementById("user-input");
    const question = userInput.value;

    document.getElementById("user-input").value = ""; // wipe the text box clean.

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
            }, 1500); // add a delay to make it look like Eliza is "typing" her message. instead of instantly displaying it.
        }
    }
    request.send(params);
});

// found method to supress the default behaviour of the enter key here.
// https://stackoverflow.com/questions/11235622/jquery-disable-form-submit-on-enter
$('#user-input').on('keyup keypress', function(e) {
    var keyCode = e.keyCode || e.which;
    if (keyCode === keyCodes.ENTER) { 
        e.preventDefault(); // default behaviour is refreshing the page, which will reset the list and lose the converstaion.
    }
});

function addListItem(speaker, text){
    const htmlString = "<li class=\"list-group-item " + speaker + "\"><p align=\"left\">" + text + "</p></li>"
    list.insertAdjacentHTML("beforeend", htmlString);
}

