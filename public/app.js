import { API } from "./services/API.js";

window.app = {
  search: event => {
    event.preventDefault();
    const query = document.querySelector(input[(type = search)]).value;
    //TODO
  },

  api: API,
};
