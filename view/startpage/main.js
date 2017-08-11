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
    xhr.onreadystatechange = function () {
        if (this.readyState === 4 && this.status === 200) {
            var decoded_response = JSON.parse(this.responseText);
            var position = decoded_response["Position"];
            definePosition(position);
        }
    };
}

function definePosition(position) {
    if (position === "admin") {
        console.log("This is admin");
    } else if (position === "student") {
        console.log("This is student");
    } else if (position === "instructor") {
        console.log("This is instructor");
    }
}