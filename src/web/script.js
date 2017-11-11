const keyCodes = {
    ENTER : 13
}

const BASE_URL = "http://localhost:8080/"

function endPoint(end){
    return BASE_URL + end;
}

$(document).ready( function(){
    // send GET request using jQuery and ajax.
    $.get(endPoint("history"), function(data){
        const history = JSON.parse(data); 
        for(var i = 0; i < history.Questions.length; i++){ // add all the past questions to maintain the state of the conversation.
            addListItem("user_message", history.Questions[i]);
            addListItem("eliza_message", history.Answers[i]);
        }    
    })
});

$('#user-input').on('keyup keypress', function(e) {
    // found method to supress the default behaviour of the enter key here.
    // https://stackoverflow.com/questions/11235622/jquery-disable-form-submit-on-enter
    var keyCode = e.keyCode;
    if(keyCode !== keyCodes.ENTER){
        return; // ignore any other keypress.
    }

    e.preventDefault(); // default behaviour is refreshing the page, which will reset the list and lose the converstaion.
    
    const question = $("#user-input").val().trim(); // remove all white space from either side of string.
    $("#user-input").val("") // wipe the text box clean.

    if(!question){
        return; // user doesn't actually have a question, don't send or add anything.
    }

    addListItem("user_message", question); // add the question the user entered.

    // jQuery docs https://api.jquery.com/jquery.get/
    // use jQuery to send POST request
    $.post(endPoint("ask"), {
        question : question // the question is a query parameter.
    }).done(function(data){ // this function gets called when the response is received.
        setTimeout(function(){ // wait a little bit before displaying elizas answer to simulate a person typing
            addListItem("eliza_message", data); 
        }, 1500); 
    }).fail(function(){
        // if there was a network issue, display a message indicating so.
        addListItem("eliza_message", "Sorry, the doctor is out, please check your connection and try again."); 
    });
});

function addListItem(speaker, text){
    const htmlString = "<li class=\"list-group-item " + speaker + "\"><p align=\"left\">" + text + "</p></li>"
    $("#conversation").append(htmlString);

    // scroll down to see the newest messages any time the list is added.
    // found this solition here https://stackoverflow.com/questions/11715646/scroll-automatically-to-the-bottom-of-the-page
    window.scrollTo(0, document.body.scrollHeight);
}

