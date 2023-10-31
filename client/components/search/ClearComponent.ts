import { SEARCH_CLEARED_EVENT } from "@/helpers/constants.ts";

export default class ClearComponent extends HTMLButtonElement {
  connectedCallback(): void {
    this.addEventListener("click", this.onClick);
  }

  disconnectedCallback(): void {
    this.removeEventListener("click", this.onClick);
  }

  private onClick(): void {
    document.getElementById("app")!.dispatchEvent(new Event(SEARCH_CLEARED_EVENT));
  }
}

customElements.define("search-clear", ClearComponent, { extends: "button" });
