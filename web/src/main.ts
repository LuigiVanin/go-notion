// Core
import { createApp } from "vue";

// Components
import App from "./App.vue";

// Helpers and Services
import { router } from "./router";

// Assets
import "./assets/styles/_root.scss";

const app = createApp(App);
app.use(router);
app.mount("#app");
