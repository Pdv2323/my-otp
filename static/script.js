function redirectToLink(event) {
    event.preventDefault();  // Prevents the form from submitting the traditional way
    window.location.href = "./verify.page.html";  // Replace with your desired link
}


function notify(msg, msgType) {
    notie.alert({
        type: msgType,
        text: msg,
    })
}