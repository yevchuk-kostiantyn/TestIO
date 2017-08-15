// Get the modal
var modal = document.getElementById('id01');

// When the user clicks anywhere outside of the modal, close it
window.onclick = function(event) {
        if (event.target === modal) {
            modal.style.display = "none";
        }
};

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
        window.location.href = "admin/index.html"
    } else if (position === "student") {
        console.log("This is student");
    } else if (position === "instructor") {
        console.log("This is instructor");
    }
}

function getSignUpInfo() {
    var first_name = document.getElementById("first_name").value;
    var last_name = document.getElementById("last_name").value;
    var email = document.getElementById("sign_up_email").value;
    var password = document.getElementById("sign_up_password").value;
    var position = getRadioButtonValue( document.getElementById('positionForm'), 'position');
    console.log(first_name, last_name, email, password, position);
}

function getRadioButtonValue(form, name) {
    var value;
    var radios = form.elements[name];

    for (var i=0, len=radios.length; i<len; i++) {
        if(radios[i].checked) {
            value = radios[i].value;
            break;
        }
    }
    return value;
}