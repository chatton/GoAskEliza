const keyCodes = {
    ENTER : 13
}

$('#user-input').on('keyup keypress', function(e) {
    // found method to supress the default behaviour of the enter key here.
    // https://stackoverflow.com/questions/11235622/jquery-disable-form-submit-on-enter
    var keyCode = e.keyCode;
    if(keyCode !== keyCodes.ENTER){
        return; // ignore any other keypress.
    }

    e.preventDefault(); // default behaviour is refreshing the page, which will reset the list and lose the converstaion.
    
    const userInput = document.getElementById("user-input");
    const question = userInput.value;
    
    userInput.value = ""; // wipe the text box clean.
    if(question.trim() === ""){
        return; // user doesn't actually have a question, don't send anything.
    }

   

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

function addListItem(speaker, text){
    const list = document.getElementById("conversation");
    const htmlString = "<li class=\"list-group-item " + speaker + "\"><p align=\"left\">" + text + "</p></li>"
    list.insertAdjacentHTML("beforeend", htmlString);
}

