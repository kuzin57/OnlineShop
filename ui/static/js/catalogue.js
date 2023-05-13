import { getCookie, setCookie } from "./cookies.js";
import { setDropdown } from "./dropdown.js";
import { makeOrder } from "./make_order.js";
import { checkNavigationPanel } from "./navigation_panel.js";
import { Product, addProductToList, removeProductFromList } from "./product.js";

checkNavigationPanel();
document.getElementById("messageBox").innerHTML = "";
fetch('/catalogue', {
    method: 'POST',
    headers: new Headers({
      'Content-Type': 'application/json',
      'Token': getCookie("token")
    })
  })
  .then(response => {
    let json_data = response.json();
    let authorized = response.headers.get('Authorized');
    console.log(authorized);
    if (authorized == "true") {
      setDropdown();

    } else {
      document.getElementById("login-username").innerHTML=`
      <a href="/auth">  
       <button>Login</button>  
      </a>
      `; 
    }

    return json_data;
  }).then(
    (data) => {
      console.log(data.products.length)
      var tbl = document.getElementById("container");

      for (let i = 0; i < data.products.length; i++) {
        var card = document.createElement("div");
        card.setAttribute("class", "card");
        card.setAttribute("id", "card " + i.toString());

        var title = document.createElement("div");
        title.setAttribute("class", "text");
        title.innerHTML = `<a href='/products/` +
                data.products[i].id + `'>` +
                data.products[i].name + `</a>`;

        card.appendChild(title);

        var divImage = document.createElement("div");
        divImage.setAttribute("class", "image");

        var img = document.createElement("img");
        img.setAttribute("src", data.products[i].path_to_image);

        divImage.appendChild(img);
        card.appendChild(divImage);

        var description = document.createElement("div");
        description.setAttribute("class", "text");
        description.innerHTML = data.products[i].price + "₽";

        card.appendChild(description);

        var buyButton = document.createElement("button");
        buyButton.setAttribute("class", "buy-button");
        buyButton.setAttribute("id", "buy-button " + i.toString());
        buyButton.setAttribute("value", "0");
        buyButton.innerHTML=`<div class="text-inside-buy-button">
                              Buy!
                            </div>`;

        card.appendChild(buyButton);


        tbl.appendChild(card);

        document.getElementById("buy-button " + i.toString()).addEventListener('click', function() {
          var button = document.getElementById("buy-button " + i.toString());

          if (button.getAttribute("value") == "0") {
            var newProduct = new Product(
              data.products[i].id,
              data.products[i].name,
              data.products[i].brand,
              data.products[i].price);
            addProductToList(newProduct);

            var productsAmount = document.createElement("input");
            productsAmount.setAttribute("id", "product_amount " + data.products[i].id.toString());
            productsAmount.setAttribute("type", "number");
            productsAmount.setAttribute("value", 1);
            productsAmount.setAttribute("min", 1);
            productsAmount.setAttribute("max", 10);
            productsAmount.setAttribute("style", "font-size: 18px;");

            var card = document.getElementById("card " + i.toString());
            card.appendChild(productsAmount);

            button.style.backgroundColor = "red";
            button.innerHTML = `<div class="text-inside-buy-button">
                                                  Remove!
                                                </div>`;
            button.setAttribute("value", "1");
            return;
          }

          removeProductFromList(data.products[i].id);
          document.getElementById("product_amount " + data.products[i].id.toString()).remove();
          button.setAttribute("value", "0");
          button.style.backgroundColor = "#3bf12a";
          button.innerHTML = `<div class="text-inside-buy-button">
                                Buy!
                              </div>`;
        });
      }

      document.getElementById("make-order").onclick = function() {
        var chosenProductsJSON = sessionStorage.getItem("chosen_products");
        if (chosenProductsJSON == null) {
          document.getElementById("messageBox").innerHTML = "You need to choose at least one product to make order";
          return;
        }

        var chosenProducts = JSON.parse(chosenProductsJSON);
        for (var i = 0; i < chosenProducts.length; i++) {
          chosenProducts[i].amount = document.getElementById(
            "product_amount " + chosenProducts[i].id.toString()).value;
        }

        var chosenProductsArray = JSON.stringify(chosenProducts);
        sessionStorage.setItem("chosen_products", chosenProductsArray);
        makeOrder();
      }
    }
  )
