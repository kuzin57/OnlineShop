function register_submit() {
    let form = document.forms["register_form"];
    form.addEventListener("submit", getValues);
}

function getValues(event) {
    event.preventDefault();
    post_user_parameters(this.username.value, this.email.value, this.password.value);
}

function post_user_parameters(username, email, password) {
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/registration", true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify({
        username: username,
        email: email,
        password: password
    }));
    // xhr.responseText
}