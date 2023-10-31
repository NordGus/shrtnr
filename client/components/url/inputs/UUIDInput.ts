import { generateUUID } from "@/helpers/uuidGeneration.ts";

export default class UUIDInput extends HTMLInputElement {
  connectedCallback(): void {
    if (this.value || this.value !== "") return

    this.value = generateUUID(8)
  }
}

customElements.define("url-uuid-input", UUIDInput, { extends: "input" });
