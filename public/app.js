import { HomePage } from "./components/HomePage.js";
import { API } from "./services/API.js";

window.addEventListener("DOMContentLoaded", event => {
  document.querySelector("main").appendChild(new HomePage());
});

window.app = {
  search: event => {
    event.preventDefault();
    const query = document.querySelector(input[(type = search)]).value;
    //TODO
  },

  api: API,
};
