export interface Product {
  id: number;
  name: string;
  brand: string;
  model: string;
  price: number;
  stock: number;
  description: string;
  image: string;
  created_at: string;
  updated_at: string;
}

export interface CartItem {
  product: Product;
  quantity: number;
}