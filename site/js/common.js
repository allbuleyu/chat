let socketUrl = "ws://127.0.0.1:8882/ws";
let apiUrl = "http://127.0.0.1:8881";

function getLocalStorage(name) {
    let value = localStorage.getItem(name);
    if (value == null) {
        return "";
    }
    return value;
}

function setLocalStorage(name, value) {
    localStorage.setItem(name, value)
}
