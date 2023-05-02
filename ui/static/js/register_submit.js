function register_submit() {
    var inputForm = document.getElementById("register_form");
    inputForm.addEventListener("submit", async (e)=> {
        e.preventDefault();
        
        const formData = new FormData(inputForm);

        if (formData.get("password") != formData.get("password_again")) {
            document.getElementById("serverMessageBox").innerHTML="Passwords don't match";
            console.log("i am in if");
            e.stopPropagation();
            return false;
        }
        console.log("i am in else");
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
                if (data.status == 403) {
                    document.getElementById("serverMessageBox").innerHTML="Registration failed!"
                } else {
                    document.getElementById("serverMessageBox").innerHTML="Registration succeeded!"
                }
            }
        ).catch(
            error => console.error(error)
        );
        
        return true;
    });
}