export interface RecipeForRequest {
  care_number: string;
  recipe_number: string;
  date: string;
  validate: string;
  handover: string;
  type: string;
  drug: RecipesForRequest[]
}

export interface RecipeCompoundForRequest {
  care_number: string;
  recipe_number: string;
  date: string;
  validate: string;
  handover: string;
  type: string;
  recipes: RecipeCompound[]
}

export interface Recipe {
  id: string;
  name: string;
  value: number;
  use:string;
  embalming:number;
  tuslah:number;
  price:number;
}

export interface RecipesForRequest {
  drug_id: string;
  value: number;
  use:string;
  embalming:number;
  tuslah:number;
  total_price:number;
}

export interface Drug {
  id: string;
  name: string;
  distributor_id: string;
  distributor: string;
  capacity: number;
  fill: number;
  unit: string;
  composition: string;
  category: string;
  price: number;
  expired_date: string;
}

export interface DrugRecipeCompound {
  name: string;
  drug_id:  string;
  value: number;
  embalming: number;
  tuslah: number;
  price: number;
}

export interface RecipeCompound {
  recipe_name: string;
  value: number;
  use: string;
  drug: DrugRecipeCompound[]
}

export interface RecipeRequest {
  care_number: string;
  recipe_number: string;
  drug: Recipe[]
}

export interface DrugData {
  id: string;
  name: string;
  distributor: string;
  capacity: number;
  fill: number;
  unit: string;
  composition: string;
  category: string;
  price: number;
  expired_date: string;
}

export interface RequestBodyDrugDataUpdate {
  id: string;
  data: DrugData;
}

export interface ResponseDrugData {
  id: string;
  name: string;
  distributor_id: string;
  distributor: string;
  capacity: number;
  fill: number;
  unit: string;
  composition: string;
  category: string;
  price: number;
  expired_date: string;
}

export interface Distributor {
  id: string;
  name: string;
  address: string;
}

export interface RecipeData {
  recipe_id: string;
  care_number: string;
  name: string;
  date: string;
  validate: string;
  validate_status:  boolean;
  handover: string;
}

export interface RecipeType {
	recipe_type: string;
	compound_name: string;
	data: DetailRecipe[]
}

export interface DetailRecipe {
  recipe_id: string;
  drug_id: string;
  drug_name: string;
  validate_status: boolean;
  compound_name: string;
  compound_value: number;
  recipe_type: string;
  value: number;
  use: string;
  embalming: number;
  tuslah: number;
  total_price: number;
}

export interface ValidateTypeRecipe {
  validate_status: boolean;
  validate_date: string;
}

export interface HandoverTypeRecipe {
  status: boolean;
  date: string;
}
