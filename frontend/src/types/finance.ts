export interface ExaminationCostRequest {
  examination_name: string
  payment_method: string
  doctor_cost: number
  nurse_cost: number
  management_cost: number
}

export interface ExaminationCost {
  id: string
  examination_name: string
  payment_method: string
  doctor_cost: number
  nurse_cost: number
  management_cost: number
}
