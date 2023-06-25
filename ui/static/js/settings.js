import { getCookie } from "./cookies.js";
import { setDropdown } from "./dropdown.js";

function send_new_user_data() {
    var inputForm = document.getElementById("settings_form");
    const formData = new FormData(inputForm);
    fetch('/settings', {
        method: 'POST',
        headers: new Headers({
          'Content-Type': 'application/json',
          'Token': getCookie("token"),
          'Email': getCookie("email")
        }),
        body: JSON.stringify({
            email: formData.get("email"),
            firstname: formData.get("firstname"),
            surname: formData.get("surname"),
            birthday: formData.get("birthday"),
            phone: formData.get("phonenumber")
        }),
      })
      .then(response => response.json())
      .then(data => {
        if (data.status == 200) {
            window.location.replace("/");
        } else {
            var messagePlace = document.getElementById("serverMessageBox");
            messagePlace.innerText = data.description;
        }
    });
}


fetch('/settings', {
    method: 'GET',
    headers: new Headers({
      'Content-Type': 'application/json',
      'Token': getCookie("token"),
      'Email': getCookie("email")
    })
  })
  .then(response => {
    let authorized = response.headers.get('Authorized');
    let userEmail = response.headers.get('Email');
    let userFirstname = response.headers.get('Firstname');
    let userSurname = response.headers.get('Surname');
    let birthdate = response.headers.get('Birthdate');
    let phoneNumber = response.headers.get('Phone-Number');
    if (authorized == "true") {
        setDropdown();
       
        var form = document.createElement("form");
        form.setAttribute("id", "settings_form");

        var emailLabel = document.createElement("label")
        emailLabel.setAttribute("for", "email");
        emailLabel.innerHTML = "Email";

        var email = document.createElement("input");
        email.setAttribute("type", "text");
        email.setAttribute("name", "email");
        email.setAttribute("id", "email");
        email.setAttribute("value", userEmail);

        var firstnameLabel = document.createElement("label")
        firstnameLabel.setAttribute("for", "firstname");
        firstnameLabel.innerHTML = "Firstname";

        var firstname = document.createElement("input");
        firstname.setAttribute("type", "text");
        firstname.setAttribute("name", "firstname");
        firstname.setAttribute("id", "firstname");
        firstname.setAttribute("value", userFirstname);

        var surnameLabel = document.createElement("label")
        surnameLabel.setAttribute("for", "surname");
        surnameLabel.innerHTML = "Surname";

        var surname = document.createElement("input");
        surname.setAttribute("type", "text");
        surname.setAttribute("name", "surname");
        surname.setAttribute("id", "surname");
        surname.setAttribute("value", userSurname);

        var birthdayLabel = document.createElement("label")
        birthdayLabel.setAttribute("for", "birthday");
        birthdayLabel.innerHTML = "Birthday";

        var birthday = document.createElement("input");
        birthday.setAttribute("type", "text");
        birthday.setAttribute("name", "birthday");
        birthday.setAttribute("id", "birthday");
        birthday.setAttribute("value", birthdate);

        var phonenumberLabel = document.createElement("label")
        phonenumberLabel.setAttribute("for", "phonenumber");
        phonenumberLabel.innerHTML = "Phonenumber";

        var phonenumber = document.createElement("input");
        phonenumber.setAttribute("type", "text");
        phonenumber.setAttribute("name", "phonenumber");
        phonenumber.setAttribute("id", "phonenumber");
        phonenumber.setAttribute("value", phoneNumber);

        var s = document.createElement("input");
        s.setAttribute("type", "submit");
        s.setAttribute("value", "Save");
        s.setAttribute("id", "send_new_user_data");

        s.addEventListener('click', send_new_user_data);

        form.appendChild(emailLabel);
        form.appendChild(email);

        form.appendChild(firstnameLabel);
        form.appendChild(firstname);

        form.appendChild(surnameLabel)
        form.appendChild(surname)

        form.appendChild(birthdayLabel)
        form.appendChild(birthday)

        form.appendChild(phonenumberLabel)
        form.appendChild(phonenumber)

        form.appendChild(s);
        
        document.getElementsByTagName("main")[0]
        .appendChild(form);
    } else {
      document.getElementById("login-username").innerHTML=`
      <a href="/auth">  
       <button>Login</button>  
      </a>
      `;
    }
});