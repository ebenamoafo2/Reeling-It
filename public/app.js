import HomePage from "./components/HomePage.js";
import { API } from "./services/API.js";
import "./components/AnimatedLoading.js";
import MovieDetailsPage from "./components/MovieDetailsPage.js";
import "./components/YouTubeEmbed.js";
import Router from "./services/Router.js";
import { MoviesPage } from "./components/MoviesPage.js";

window.addEventListener("DOMContentLoaded", event => {
  app.Router.init();
});

window.app = {
  Router,
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
  search: event => {
    event.preventDefault();
    const keywords = document.querySelector("input[type=search]").value;
    if (keywords.length > 1) {
      app.Router.go(`/movies?q=${keywords}`);
    }
  },
  searchOrderChange: order => {
    const urlParams = new URLSearchParams(window.location.search);
    const q = urlParams.get("q");
    const genre = urlParams.get("genre") ?? "";
    app.Router.go(`/movies?q=${q}&order=${order}&genre=${genre}`);
  },
  searchFilterChange: genre => {
    const urlParams = new URLSearchParams(window.location.search);
    const q = urlParams.get("q");
    const order = urlParams.get("order") ?? "";
    app.Router.go(`/movies?q=${q}&order=${order}&genre=${genre}`);
  },
  api: API,
};
