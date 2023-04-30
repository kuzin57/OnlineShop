function auth_submit() {
    // let form = document.forms["auth_form"];
    const url = "/127.0.0.1:7000/auth"

    var inputForm = document.getElementById("auth_form");
    inputForm.addEventListener("submit", (e)=> {
        e.preventDefault()
        
        const formData = new FormData(inputForm)

        fetch("/auth", {
            method: "POST",
            body: formData,
        }).then(
            response => response.text()
        ).then(
            (data) => {console.log(data); document.getElementById("serverMessageBox").innerHTML = data}
        ).catch(
            error => console.error(error)
        )
    });
}