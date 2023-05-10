import { getCookie, setCookie } from "./cookies.js";

function send_new_password_to_server() {
    document.getElementById("serverMessageBox").innerHTML = "";
    var inputForm = document.getElementById("new_password_form");
    inputForm.addEventListener("submit", async (e)=> {
        e.preventDefault();
        
        const formData = new FormData(inputForm);

        if (formData.get("password") != formData.get("password_again")) {
            document.getElementById("serverMessageBox").innerHTML = "Passwords do not match!"
            return false;
        }

        fetch("/recovery", {
            method: "POST",
            headers: new Headers({
                'Content-Type': 'application/json',
                'Email': getCookie("email"),
                'New-Password': formData.get("password")
            })
        }).then(
            response => response.json()
        ).then(
            (data) => {
                console.log("status ", data.status)
                if (data.status == 200) {
                    window.location.replace("/auth");
                } else {
                    document.getElementById("serverMessageBox").innerHTML = data.description
                }
            }
        ).catch(
            error => console.error(error)
        );
        
        return true;
    });
}

function set_new_password() {
    document.getElementById("serverMessageBox").innerHTML = "";
    console.log("I am in set new password!");
    var inputForm = document.getElementById("code_form");
    inputForm.addEventListener("submit", async (e)=> {
        e.preventDefault();
        
        const formData = new FormData(inputForm);

        fetch("/recovery", {
            method: "POST",
            headers: new Headers({
                'Content-Type': 'application/json',
                'Email': getCookie("email"),
                'Code': formData.get("code")
            })
        }).then(
            response => response.json()
        ).then(
            (data) => {
                if (data.status == 200) {
                    console.log("oops");
                    document.getElementById("code_form").remove();

                    var form = document.createElement("form");
                    form.setAttribute("id", "new_password_form");

                    var passwordLabel = document.createElement("label")
                    passwordLabel.setAttribute("for", "password");
                    passwordLabel.innerHTML = "Enter new password";

                    var password = document.createElement("input");
                    password.setAttribute("type", "password");
                    password.setAttribute("name", "password");
                    password.setAttribute("id", "password");

                    var passwordAgainLabel = document.createElement("label");
                    passwordAgainLabel.setAttribute("for", "password_again");
                    passwordAgainLabel.innerHTML = "Repeat your new password again";

                    var passwordAgain = document.createElement("input");
                    passwordAgain.setAttribute("type", "password");
                    passwordAgain.setAttribute("name", "password_again");
                    passwordAgain.setAttribute("id", "password_again");

                    var s = document.createElement("input");
                    s.setAttribute("type", "submit");
                    s.setAttribute("value", "Confirm");
                    s.setAttribute("id", "send_new_password");

                    s.addEventListener('click', send_new_password_to_server);

                    form.appendChild(passwordLabel);
                    var br1 = document.createElement('br');
                    form.appendChild(br1);

                    form.appendChild(password);
                    var br2 = document.createElement('br');
                    form.appendChild(br2);

                    form.appendChild(passwordAgainLabel);
                    var br3 = document.createElement('br');
                    form.appendChild(br3);

                    form.appendChild(passwordAgain);
                    var br4 = document.createElement('br');
                    form.appendChild(br4);

                    form.appendChild(s);
                    
                    document.getElementsByTagName("main")[0]
                    .appendChild(form);
                }
            }
        ).catch(
            error => console.error(error)
        );
        
        return true;
    });
}

export function send_code_to_email() {
    document.getElementById("serverMessageBox").innerHTML = "";
    console.log("sending code...");
    var inputForm = document.getElementById("access_form");
    inputForm.addEventListener("submit", async (e)=> {
        e.preventDefault();
        
        const formData = new FormData(inputForm);
        setCookie("email", formData.get("email"))

        fetch("/recovery", {
            method: "POST",
            headers: new Headers({
                'Content-Type': 'application/json',
                'Email': formData.get("email")
            })
        }).then(
            response => response.json()
        ).then(
            (data) => {
                if (data.status == 200) {
                    document.getElementById("access_form").remove();

                    var form = document.createElement("form");
                    form.setAttribute("id", "code_form");

                    var code = document.createElement("input");
                    code.setAttribute("type", "text");
                    code.setAttribute("name", "code");
                    code.setAttribute("placeholder", "enter your code here");

                    var s = document.createElement("input");
                    s.setAttribute("type", "submit");
                    s.setAttribute("value", "Submit");
                    // s.setAttribute("onclick", "return set_new_password();");
                    s.setAttribute("id", "set_password")
                    s.addEventListener("click", set_new_password);

                    form.append(code);

                    form.append(s);

                    document.getElementsByTagName("main")[0]
                    .appendChild(form);

                } else {
                    document.getElementById("serverMessageBox").innerHTML = data.description
                }
            }
        ).catch(
            error => console.error(error)
        );
        
        return true;
    });
}
