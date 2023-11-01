export default class CardComponent extends HTMLElement {
  connectedCallback(): void {
    const searchInput = document.querySelector<HTMLInputElement>("#search-url-form input")!;

    if (searchInput.value === "") return;

    const target = this.querySelector<HTMLAnchorElement>(".entry-to")!;
    const from = this.querySelector<HTMLAnchorElement>(".entry-from")!;

    if (!target.href.includes(searchInput.value) && !from.href.includes(searchInput.value)) { this.remove(); return; }

    if (!this.parentElement!.querySelector<HTMLButtonElement>("button")) return;

    const index = Array.from(this.parentElement!.children).indexOf(this);

    if (index !== 0) return;

    this.parentElement!.insertBefore(this.parentElement!.children.item(1)!, this)
  }
}

customElements.define("url-card", CardComponent);
