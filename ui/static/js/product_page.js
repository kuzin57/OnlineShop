import { getCookie } from "./cookies.js";
import { setDropdown } from "./dropdown.js";
import { Product, addProductToList, removeProductFromList } from "./product.js";

let index = document.URL.lastIndexOf("/");
let length = document.URL.length;
let productId = parseInt(document.URL.substring(index + 1, length));

function setBuyButton(product) {
    var footer = document.getElementsByTagName("footer").item(0);
    var buyButton = document.createElement("button");
    buyButton.setAttribute("class", "buy-button");
    buyButton.setAttribute("id", "buy-button");
    buyButton.setAttribute("value", "0");
    buyButton.innerHTML=`<div class="text-inside-buy-button">
                        Buy!
                        </div>`;
    footer.appendChild(buyButton);

    document.getElementById("buy-button").addEventListener('click', function() {
        if (sessionStorage.getItem("authorized") == "false") {
            window.location.replace("/auth");
            return; 
        }

        var button = document.getElementById("buy-button");

        if (button.getAttribute("value") == "0") {
        var newProduct = new Product(
            productId,
            product.name,
            product.brand,
            product.price);
        addProductToList(newProduct);

        var productsAmount = document.createElement("input");
        productsAmount.setAttribute("id", "product_amount " + productId.toString());
        productsAmount.setAttribute("type", "number");
        productsAmount.setAttribute("value", 1);
        productsAmount.setAttribute("min", 1);
        productsAmount.setAttribute("max", 10);
        productsAmount.setAttribute("style", "font-size: 18px;");

        var footer = document.getElementsByTagName("footer").item(0);
        footer.appendChild(productsAmount);

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

function setDetailedInfo(event, product) {
    var description = document.getElementById("description");
    description.style.display = "none";

    var title = document.createElement("div");
    title.setAttribute("style", "font-size: 20pt;");
    title.innerText = product.name;
    product_details.appendChild(title);

    var price = document.createElement("div");
    price.setAttribute("style", "font-size: 30pt; color: red;");
    price.innerText = product.price + "₽";
    product_details.appendChild(price);
}

fetch('/products/', {
    method: 'POST',
    headers: new Headers({
      'Content-Type': 'application/json',
      'Token': getCookie("token"),
      'Product-ID': productId
    })
  })
  .then(response => {
    let json_data = response.json();
    let authorized = response.headers.get('Authorized');
    if (authorized == "true") {
      sessionStorage.setItem("authorized", "true");
      setDropdown();
      var dropdown = document.getElementsByClassName("dropdown").item(0);
      dropdown.setAttribute("style", "margin-left: 90%;");
    } else {
      sessionStorage.setItem("authorized", "false");
      document.getElementById("login-username").innerHTML=`
      <a href="/auth">  
       <button style="margin-left: 90%;">Login</button>  
      </a>
      `; 
    }
    return json_data;
  })
  .then(
    (data) => {
        let product = data.products[0];
        console.log(product);

        var image_div = document.getElementById("img-div");
        image_div.setAttribute("id", "product-container__image");

        var image = document.createElement("img");
        image.setAttribute("src", product.path_to_image);
        image.setAttribute("style", "width: 50%; height: 90%; margin-top: 2%;");

        image_div.appendChild(image);

        var detailedInfoArea = document.getElementById("detailed_info");

        var title = document.createElement("div");
        title.setAttribute("style", "font-size: 20pt;");
        title.innerText = product.name;
        detailedInfoArea.appendChild(title);   

        var price = document.createElement("div");
        price.setAttribute("style", "color: red; font-size: 20pt;");
        price.innerText = product.price + "₽";
        detailedInfoArea.appendChild(price);

        // list
        var list = document.createElement("ul");
        var brand = document.createElement("li");
        brand.setAttribute("style", "margin-left: 2%;");
        brand.innerText = "Производитель: " + product.brand;
        var kcal = document.createElement("li");
        kcal.innerText = "Энергетическая ценность: " + product.kcal + "ккал";
        kcal.setAttribute("style", "margin-left: 2%;");
        var proteins = document.createElement("li");
        proteins.innerText = "Белки: " + product.proteins + "г";
        proteins.setAttribute("style", "margin-left: 2%;");
        var fats = document.createElement("li");
        fats.innerText = "Жиры: " + product.fats + "г";
        fats.setAttribute("style", "margin-left: 2%;");
        var carbohydrates = document.createElement("li");
        carbohydrates.innerText = "Углеводы: " + product.carbohydrates + "г"; 
        carbohydrates.setAttribute("style", "margin-left: 2%;");
        var country = document.createElement("li");
        country.setAttribute("style", "margin-left: 2%;");
        country.innerText = "Страна производителя: " + product.country;
        list.appendChild(brand);
        list.appendChild(kcal);
        list.appendChild(proteins);
        list.appendChild(proteins);
        list.appendChild(fats);
        list.appendChild(carbohydrates);
        list.appendChild(country);

        detailedInfoArea.appendChild(list);

        var descriptionArea = document.getElementById("description");
        var text = document.createElement("div");
        text.setAttribute("style", "font-size: 20pt;")
        // text.innerText = product.description;
        descriptionArea.appendChild(text);

        descriptionArea.style.display = "none";
        detailedInfoArea.style.display = "none";

        var detailsButton = document.getElementById("details_button");
        var descButton = document.getElementById("desc_button");
        
        detailsButton.addEventListener("click", function() {
            console.log("details button clicked!");
            var tabcontent = document.getElementsByClassName("tabcontent");
            for (var i = 0; i < tabcontent.length; i++) {
                tabcontent[i].style.display = "none";
            }

            var detailedInfo  = document.getElementById("detailed_info");
            detailedInfo.style.display = "block";         
        });

        descButton.addEventListener("click", function() {
            console.log("description button clicked!");

            var tabcontent = document.getElementsByClassName("tabcontent");
            for (var i = 0; i < tabcontent.length; i++) {
                tabcontent[i].style.display = "none";
            }

            var description = document.getElementById("description");
            description.style.display = "block";
        });

        // buy-button
        setBuyButton(product);
    }
  );