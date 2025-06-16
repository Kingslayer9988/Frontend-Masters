export class HomePage extends HTMLElement { //homepage
    constructor() {
        super();
    }
    connectedCallback() {
        const template = document.getElementById("template-home")
        const content = template.content.cloneNode(true);
        this.appendChild(content);
    }
}

customElements.define("home-page", HomePage);