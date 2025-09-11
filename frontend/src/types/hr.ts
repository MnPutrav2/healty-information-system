export interface EmployeeData {
  id: string;
  name: string;
  gender: string;
  birth_place: string;
  birth_date: string;
  wedding: string;
  address: string;
  nik: string;
  bpjs: string;
  npwp: string;
  phone_number: string;
  email: string;
  status: boolean
}

export interface UserReq {
  id: string | undefined
  employee_id: string
  username: string
  password: string
}

export interface PathAccess {
  name: string,
  child: PathAccessChild[]
}

export interface PathAccessChild {
  name: string,
}
