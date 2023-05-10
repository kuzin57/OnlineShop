import { delete_cookies, getCookie, setCookie } from "./cookies.js";

fetch('/', {
    method: 'GET',
    headers: new Headers({
      'Content-Type': 'application/json',
      'Token': getCookie("token")
    })
  })
  .then(response => {
    let authorized = response.headers.get('Authorized');
    let profile = response.headers.get('Profile');
    console.log(authorized);
    if (authorized == "true") {
      document.getElementById("login-username").innerHTML=`
      <div class="dropdown">
        <button>` + getCookie("username") + `</button>
        <div class="dropdown-options">
          <button>Orders</button>
          <button>Settings</button>
          <button id="logout-button">Logout</button>
        </div>
      </div>
      `;

      const logout = document.getElementById("logout-button");
      logout.addEventListener("click", function() {
        setCookie("token", "");
        window.location.reload();
      });
       
      if (document.URL == "http://localhost:7000/") {
        document.getElementById("serverMessageBox").innerHTML = "Hello, " + getCookie("username");  
      }
    } else {
      document.getElementById("login-username").innerHTML=`
      <a href="/auth">  
       <button>Login</button>  
      </a>
      `;
    }
  })
