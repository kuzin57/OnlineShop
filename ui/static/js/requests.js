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
      document.getElementById("Dropdown").innerHTML=`
      <div class="dropdown">
        <button>` + getCookie("username") + `</button>
        <div class="dropdown-options">
          <a href="/orders">Orders</a>
          <a href="/settings">Settings</a>
          <a href="/" id="logout">Logout</a>
        </div>
      </div>
      `;

      const logout = document.getElementById("logout");
      logout.addEventListener("click", function() {
        setCookie("token", "");
        window.location.reload();
      });
       
      if (document.URL == "http://localhost:7000/") {
        document.getElementById("serverMessageBox").innerHTML = "Hello, " + getCookie("username");  
      }
    } else {
      document.getElementById("login").innerHTML=`
      <form id="login button" action="/auth">
        <input type="submit" style="float: right;" value="Login"/>
      </form>
      `; 
    }
  })
