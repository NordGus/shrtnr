import { v4 } from "uuid"

class IDInput extends HTMLInputElement {
  connectedCallback(): void {
    if (this.value || this.value !== "") return

    this.value = v4()
  }
}

customElements.define("url-id-input", IDInput, { extends: "input" });
