import { SEARCH_CLEARED_EVENT } from "@/helpers/constants.ts";

export default class SearchInput extends HTMLInputElement {
  connectedCallback(): void {
    document.getElementById("app")!.addEventListener(SEARCH_CLEARED_EVENT, () => { this.value = "" })
  }

  disconnectedCallback(): void {
    document.getElementById("app")!.removeEventListener(SEARCH_CLEARED_EVENT, () => { this.value = "" })
  }
}

customElements.define("search-input", SearchInput, { extends: "input" });
