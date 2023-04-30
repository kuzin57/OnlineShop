function auth_submit() {
    // let form = document.forms["auth_form"];
    const url = "/127.0.0.1:7000/auth"

    var inputForm = document.getElementById("auth_form");
    inputForm.addEventListener("submit", getValues);
}

function getValues(event) {
    event.preventDefault();
    post_email_and_password(this.email.value, this.password.value);
    // this.http.request(this.server)
}

function post_email_and_password(email, password) {
    fetch("/auth", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
            email: email,
            password: password
        }),
    }).then(
        response => {console.log(response.text());document.getElementById("status").innerText = response.status;}
    ).catch(
        error => console.error(error)
    )
    // document.getElementById("status").innerHTML = JSON.stringify(response.text());



    // var xhr = new XMLHttpRequest();
    // xhr.open("POST", "/auth", true);
    // xhr.setRequestHeader('Content-Type', 'application/json');
    // xhr.send(JSON.stringify({
    //     email: email,
    //     password: password
    // }));
    // document.getElementById("status").innerHTML = xhr.status;
    // console.log(xhr.status);
}