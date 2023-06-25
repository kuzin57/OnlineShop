var inputForm = document.getElementById("register_form");
inputForm.addEventListener("submit", async (e)=> {
    e.preventDefault();
    
    const formData = new FormData(inputForm);

    if (formData.get("password") != formData.get("password_again")) {
        document.getElementById("serverMessageBox").innerHTML="Passwords don't match";
        e.stopPropagation();
        return false;
    }

    if (formData.get("password") == "") {
        document.getElementById("serverMessageBox").innerHTML="Password can't be empty!";
        e.stopPropagation();
        return false;
    }

    fetch("/registration", {
        method: "POST",
        body: JSON.stringify({
            email: formData.get("email"),
            password: formData.get("password"),
            firstname: formData.get("firstname"),
            surname: formData.get("surname"),
            birthday: formData.get("birthday"),
            phone: formData.get("phone number")
        }),
    }).then(
        response => response.json()
    ).then(
        (data) => {
            console.log(data)
            document.getElementById("serverMessageBox").innerHTML=data.description
            if (data.status != 403) {
                document.getElementById("register_form").remove()
                document.getElementById("serverMessageBox").innerHTML="<h2>Thank you for registration!</h2><h3>Now you can <a href='/auth'>log in</a></h3>"
            }
        }
    ).catch(
        error => console.error(error)
    );
    
    return true;
});

