import { generateUUID } from "@/helpers/uuidGeneration.ts";

class UUIDInput extends HTMLInputElement {
  connectedCallback(): void {
    console.log(this.value)
    if (this.value || this.value !== "") return

    this.value = generateUUID(8)

    console.log(this.value)
  }
}

customElements.define("url-uuid-input", UUIDInput, { extends: "input" });
