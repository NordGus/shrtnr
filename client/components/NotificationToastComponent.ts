export default class NotificationToastComponent extends HTMLElement {
  private timerID!: number

  connectedCallback(): void {
    this.timerID = setTimeout(() => { this.remove() }, 3000)
  }

  disconnectedCallback(): void {
    clearTimeout(this.timerID)
    this.timerID = 0
  }
}

customElements.define("notification-toast", NotificationToastComponent);
