export function labNumber(date: Date, length: number): string {
  const dt: string = String(date.getDate()).padStart(2, '0')
  const month: string  = String(date.getMonth() + 1).padStart(2, '0')

  return `LAB${date.getFullYear()}${month}${dt}${String(1 + length).padStart(3, '0')}`
}
