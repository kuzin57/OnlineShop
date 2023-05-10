import { getCookie, setCookie } from "./cookies.js";

export function setDropdown() {
    document.getElementById("login-username").innerHTML=`
      <div class="dropdown">
        <button>` + getCookie("username") + `</button>
        <div class="dropdown-options">
          <button>Orders</button>
          <button id="settings-button">Settings</button>
          <button id="logout-button">Logout</button>
        </div>
      </div>
      `;

    const logout = document.getElementById("logout-button");
    logout.addEventListener("click", function() {
    setCookie("token", "");
    window.location.reload();
    });

    const settings = document.getElementById("settings-button");
    settings.addEventListener("click", function() {
    window.location.replace("/settings");
    });
}