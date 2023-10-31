export default class CardComponent extends HTMLElement {
  connectedCallback(): void {
    const searchInput = document.querySelector<HTMLInputElement>("#search-url-form input")!;

    if (searchInput.value === "") return;

    const target = this.querySelector<HTMLAnchorElement>(".entry-to")!;

    if (!target.href.includes(searchInput.value)) { this.remove(); return; }

    if (!this.parentElement!.querySelector<HTMLButtonElement>("button")) return;

    const index = Array.prototype.indexOf.call(this.parentElement!.children, this);

    if (index !== 0) return;

    this.parentElement!.insertBefore(this.parentElement!.children.item(1)!, this)
  }
}

customElements.define("url-card", CardComponent);
