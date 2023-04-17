function auth_submit() {
    let form = document.forms["auth_form"];
    form.addEventListener("submit", getValues);
}

function getValues(event) {
    event.preventDefault();
    post_email_and_password(this.email.value, this.password.value);
}

function post_email_and_password(form, password) {
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/auth", true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify({
        email: form,
        password: password
    }));
}