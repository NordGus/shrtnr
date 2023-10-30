export default class NotificationTimerVisualizerComponent extends HTMLElement {
  private timerID!: number

  connectedCallback(): void {
    this.timerID = setTimeout(() => {
      const timerLine = this.firstElementChild!

      timerLine.classList.remove("w-0")
      timerLine.classList.add("w-full")
    }, 50)
  }

  disconnectedCallback(): void {
    clearTimeout(this.timerID)
    this.timerID = 0
  }
}

customElements.define("notification-timer-visualizer", NotificationTimerVisualizerComponent);
