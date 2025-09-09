export function formatDate(date: Date): string {
  const defaultDate: Date = date
  const datez: string = defaultDate.getDate().toString()
  const year: string = defaultDate.getFullYear().toString()
  const month: number = defaultDate.getMonth() + 1

  let formatMonth: string = ""
  let formatDate: string = ""

  if (month.toString().length == 1) {
    formatMonth = "0" + month.toString()
  } else {
    formatMonth = month.toString()
  }

  if (datez.length == 1) {
    formatDate = "0" + datez
  } else {
    formatDate = datez
  }

  return `${year}-${formatMonth}-${formatDate}`
}

export function formatDatetime(date: Date, time: string | null): string {
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  const hours = String(date.getHours()).padStart(2, '0')
  const minutes = String(date.getMinutes()).padStart(2, '0')
  const seconds = String(date.getSeconds()).padStart(2, '0')

  let timeFormatted = `${hours}:${minutes}:${seconds}`

  if (time) {
    // Jika time diberikan, pastikan dalam format HH:mm:ss
    const parts = time.split(':')
    if (parts.length === 2) {
      timeFormatted = `${parts[0]}:${parts[1]}:00` // tambahkan detik jika tidak ada
    } else if (parts.length === 3) {
      timeFormatted = time // sudah lengkap
    } else {
      console.warn('formatDatetime: time format tidak valid, gunakan HH:mm atau HH:mm:ss')
    }
  }

  return `${year}-${month}-${day}T${timeFormatted}`
}


export function viewedDateTime(date: string): string {
  const cutDate = date.split('T')
  const cutTime = cutDate[1].split('Z')

  return `${cutDate[0]} ${cutTime[0]}`
}

export function viewedDate(date: string): string {
  const cutDate = date.split('T')

  return `${cutDate[0]}`
}
