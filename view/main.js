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
    var url = "/login";
    xhr.open("PATCH", url, true);
    xhr.setRequestHeader("Content-type", "application/json");
    xhr.onreadystatechange = function () {
        if (this.readyState === 4 && this.status === 200) {
            var position = JSON.parse(this.responseText);
            console.log("Position:", position);
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
        window.location.href = "student/index.html"
    } else if (position === "instructor") {
        console.log("This is instructor");
        window.location.href = "instructor/index.html"
    } else {
        console.log("Unknown position!")
    }
}

function getSignUpInfo() {
    var first_name = document.getElementById("first_name").value;
    var last_name = document.getElementById("last_name").value;
    var email = document.getElementById("sign_up_email").value;
    var password = document.getElementById("sign_up_password").value;
    var position = getRadioButtonValue( document.getElementById('positionForm'), 'position');

    sendNewUserInfo(first_name, last_name, email, password, position);
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

function sendNewUserInfo(first_name, last_name, email, password, position) {
    var xhr = new XMLHttpRequest();
    requestSignUp(xhr, first_name);

    var new_user_info = JSON.stringify(
        {
            "first_name": first_name,
            "last_name": last_name,
            "email": email,
            "password": password,
            "position": position
        }
    );
    xhr.send(new_user_info);
    console.log("New User Info to be sent:", new_user_info);
}

function requestSignUp(xhr, first_name) {
    var url = "/signup";
    xhr.open("PATCH", url, true);
    xhr.setRequestHeader("Content-type", "application/json");
    xhr.onreadystatechange = function () {
        if (this.readyState === 4 && this.status === 200) {
            alert(first_name+", you have been successfully signed up!")
        }
    };
}