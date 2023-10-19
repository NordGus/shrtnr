export function generateUUID(size: number): string {
  const arr = new Uint8Array(size)
  const sets = [
    ...generateSet(48, 57), // numbers
    ...generateSet(97, 122), // lower alphabetic values
    ...generateSet(65, 90), // upper alphabetic values
  ]

  window.crypto.getRandomValues(arr)

  return Array.from(arr, (num) => { return sets[num % sets.length] }).join("")
}

function generateSet(min: number, max: number): string[] {
  return Array.from(
    new Array<string>(max-min),
    (_, idx) => { return String.fromCharCode(min + idx) }
  )
}
