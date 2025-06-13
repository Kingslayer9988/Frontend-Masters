import { API } from "./services/API.js";

windows.app = {
    search: (event) => {
        event.preventDefault();
        const q = document.querySelector('input[type=search]').value;
        // TODO API CALL
    },
    api: API
} 