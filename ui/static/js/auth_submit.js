import { setCookie } from "./cookies.js";

export function auth_submit() {
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
                document.getElementById("serverMessageBox").innerHTML=data.description;
                if (data.status != 403) {
                    setCookie("username", data.userName);
                    setCookie("token", data.token);
                    console.log(document.cookie);
                    window.location.replace("/");
                }
            }
        ).catch(
            error => console.error(error)
        )
    });
}