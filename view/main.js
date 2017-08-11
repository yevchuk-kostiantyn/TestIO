function getCredentials() {
    var userName = document.getElementById("username").value;
    var password = document.getElementById("password").value;

    console.log("Welcome to TestIO Javascript Logs!");
    console.log("Entered Username:", userName);
    console.log("Entered Password:", password);

    sendCredentials(userName, password);
}

function sendCredentials(username, password) {
    var xhr = new XMLHttpRequest();
    requestHandler(xhr);

    var credentials = JSON.stringify(
        {
            "username": username,
            "password": password
        }
    );
    xhr.send(credentials);
    console.log("Credentials to be sent:", credentials);
}

function requestHandler(xhr) {
    var url = "/patch";
    xhr.open("PATCH", url, true);
    xhr.setRequestHeader("Content-type", "application/json");
    // TODO
}