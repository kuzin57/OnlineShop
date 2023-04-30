function auth_submit() {
    var inputForm = document.getElementById("auth_form");
    inputForm.addEventListener("submit", async (e)=> {
        e.preventDefault()
        
        const formData = new FormData(inputForm)
        console.log(formData.get("email"))

        fetch("/auth", {
            method: "POST",
            body: JSON.stringify({email: formData.get("email"), password: formData.get("password")}),
        }).then(
            response => response.json()
        ).then(
            (data) => {
                console.log(data)
                if (data.status == 403) {
                    document.getElementById("serverMessageBox").innerHTML="Login failed!"
                } else {
                    document.getElementById("serverMessageBox").innerHTML="Login succeeded!"
                }
            }
        ).catch(
            error => console.error(error)
        )
        
    });
}