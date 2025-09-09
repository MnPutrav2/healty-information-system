export interface SearchLimit {
  search: string
  limit: number
}

export interface ResponseError {
  status: string
  errors: string
}

export interface ResponseSuccess {
  status: string
  response: string
}

export interface ResponseToken {
  status: string
  token: string
}
