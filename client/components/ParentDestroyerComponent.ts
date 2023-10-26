export default class ParentDestroyerComponent extends HTMLElement {
  connectedCallback(): void {
    this.addEventListener("click", this.destroyParent);
  }

  disconnectedCallback(): void {
    this.removeEventListener("click", this.destroyParent);
  }

  private destroyParent(): void {
    this.parentElement!.remove();
  }
}

customElements.define("parent-destroyer", ParentDestroyerComponent);
