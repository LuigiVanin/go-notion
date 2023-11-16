// Core
import { createApp } from "vue";
import { createPinia } from "pinia";

// Libraries
import Toast, { type PluginOptions, POSITION } from "vue-toastification";

// Components
import App from "./App.vue";

// Helpers and Services
import { router } from "./router";

// Assets
import "./assets/styles/_root.scss";
import "vue-toastification/dist/index.css";

const pinia = createPinia();

const app = createApp(App);

app.use(pinia);
app.use(router);
app.use(Toast, {
    draggable: true,
    position: POSITION.TOP_RIGHT,
    timeout: 3500,
} satisfies PluginOptions);
app.mount("#app");
