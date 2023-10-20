export function generateUUID(size: number): string {
  const arr = new Uint8Array(size)
  const sets = [
    ...generateSet(48, 57), // numbers
    ...generateSet(97, 122), // lower alphabetic values
    ...generateSet(65, 90), // upper alphabetic values
  ]

  window.crypto.getRandomValues(arr)

  for (let i = sets.length - 1; i > 0; i--) {
    const j: number = Math.floor(Math.random() * (i + 1))
    const tmp = sets[i]
    sets[i] = sets[j]
    sets[j] = tmp
  }

  return Array.from(arr, (num) => { return sets[num % sets.length] }).join("")
}

function generateSet(min: number, max: number): string[] {
  return Array.from(
    new Array<string>(max-min),
    (_, idx) => { return String.fromCharCode(min + idx) }
  )
}
