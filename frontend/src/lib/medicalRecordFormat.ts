export function medicalRecord(mr: string): string {
  const medicalRecord = String(mr).padStart(6, "0")

  return medicalRecord
}
