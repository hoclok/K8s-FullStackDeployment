import { API_CONFIG } from '@/config/api';

export interface Product {
  id: number;
  name: string;
  description: string;
  price: number;
  stock: number;
  created_at: string;
  updated_at: string;
}

export interface CreateProductRequest {
  name: string;
  description: string;
  price: number;
  stock: number;
}

class ProductService {
  private baseUrl = API_CONFIG.BASE_URL;

  async getAllProducts(): Promise<Product[]> {
    const response = await fetch(`${this.baseUrl}${API_CONFIG.ENDPOINTS.PRODUCTS.GET_ALL}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    });

    if (!response.ok) {
      throw new Error('Failed to fetch products');
    }

    return response.json();
  }

  async createProduct(data: CreateProductRequest): Promise<Product> {
    const response = await fetch(`${this.baseUrl}${API_CONFIG.ENDPOINTS.PRODUCTS.CREATE}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });

    if (!response.ok) {
      throw new Error('Failed to create product');
    }

    return response.json();
  }

  async deleteProduct(id: number): Promise<{ message: string }> {
    const response = await fetch(`${this.baseUrl}${API_CONFIG.ENDPOINTS.PRODUCTS.DELETE(id)}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
      },
    });

    if (!response.ok) {
      throw new Error('Failed to delete product');
    }

    return response.json();
  }
}

export const productService = new ProductService(); 