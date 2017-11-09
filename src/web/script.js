const keyCodes = {
    ENTER : 13
}


$(document).ready( function(){
    
    const request = new XMLHttpRequest();
    request.open("GET", "http://localhost:8080/history", true)
    request.onreadystatechange = function(){
        if(request.readyState === XMLHttpRequest.DONE) {
            var response = request.responseText; // the history is returned in json format.
            const history = JSON.parse(response); 
            for(var i = 0; i < history.questions.length; i++){ // add all the past questions to maintain the state of the conversation.
                addListItem("user_message", history.questions[i]);
                addListItem("eliza_message", history.answers[i]);
            }
        }
    }
    request.send(null);
});

$('#user-input').on('keyup keypress', function(e) {
    // found method to supress the default behaviour of the enter key here.
    // https://stackoverflow.com/questions/11235622/jquery-disable-form-submit-on-enter
    var keyCode = e.keyCode;
    if(keyCode !== keyCodes.ENTER){
        return; // ignore any other keypress.
    }

    e.preventDefault(); // default behaviour is refreshing the page, which will reset the list and lose the converstaion.
    
    const userInput = document.getElementById("user-input");
    const question = userInput.value.trim(); // remove all white space from either side of string.

    userInput.value = ""; // wipe the text box clean.
    if(question === ""){
        return; // user doesn't actually have a question, don't send anything.
    }

   
    // there's actually a question to send to Eliza.
    const request = new XMLHttpRequest();
    
    request.open("POST", "http://localhost:8080/ask", true); // open the connection
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

    // add the question as a POST parameter.
    const params = "question=" + question; 
    // send the actual request.
    request.send(params);
});

function addListItem(speaker, text){
    const list = document.getElementById("conversation");
    const htmlString = "<li class=\"list-group-item " + speaker + "\"><p align=\"left\">" + text + "</p></li>"
    list.insertAdjacentHTML("beforeend", htmlString);
    
    // scroll down to see the newest messages any time the list is added.
    // found this solition here https://stackoverflow.com/questions/11715646/scroll-automatically-to-the-bottom-of-the-page
    window.scrollTo(0, document.body.scrollHeight);
}

