document.addEventListener("DOMContentLoaded", function() {
    var container = document.getElementById("App");
    var message = document.createElement("p");
    message.appendChild(document.createTextNode("Hello, Caddy World!"));
    container.appendChild(message);
});
