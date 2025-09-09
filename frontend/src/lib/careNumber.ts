export function careNumber(date: Date, length: number): string {
  const dates: string = `${date.getFullYear()}${String(date.getMonth() + 1).padStart(2, '0')}${String(date.getDate()).padStart(2, '0')}`
  const lt: string = String(length + 1).padStart(6, "0")

  return `${dates}${lt}`
}

export function regNumber(length: number): string {
  const reg = String(length + 1).padStart(3, '0')

  return reg
}

export function recipeNumber(date: Date, length: number): string {
  const dates: string = `R${date.getFullYear()}${String(date.getMonth() + 1).padStart(2, '0')}${String(date.getDate()).padStart(2, '0')}`
  const lt: string = String(length + 1).padStart(5, "0")

  return `${dates}${lt}`
}
