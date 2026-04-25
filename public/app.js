import HomePage from "./components/HomePage.js";
import { API } from "./services/API.js";
import "./components/AnimatedLoading.js";
import MovieDetailsPage from "./components/MovieDetailsPage.js";
import "./components/YouTubeEmbed.js";
import Router from "./services/Router.js";
import { MoviePage } from "./components/MoviePage.js";

window.addEventListener("DOMContentLoaded", event => {
  app.Router.init();
});

window.app = {
  Router,
  search: event => {
    event.preventDefault();
    const query = document.querySelector("input[type='search']").value;
    //TODO
  },
  showError: (
    message = "There was an error loading the page",
    goToHome = true,
  ) => {
    document.querySelector("#alert-modal").showModal();
    document.querySelector("#alert-modal p").textContents = message;
    if (goToHome) app.Router.go("/");
    return;
  },
  closeError: () => {
    document.getElementById("alert-modal").close();
  },

  api: API,
};
