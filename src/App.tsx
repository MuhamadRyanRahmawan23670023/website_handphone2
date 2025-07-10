import React from 'react';
import { useState, useEffect } from 'react';
import { Header } from './components/Header';
import { ProductCard } from './components/ProductCard';
import { Cart } from './components/Cart';
import { useCart } from './hooks/useCart';
import { Product } from './types/product';

function App() {
  const [products, setProducts] = useState<Product[]>([]);
  const [filteredProducts, setFilteredProducts] = useState<Product[]>([]);
  const [searchQuery, setSearchQuery] = useState('');
  const [isCartOpen, setIsCartOpen] = useState(false);
  const [loading, setLoading] = useState(true);
  
  const { cartItems, addToCart, updateQuantity, removeFromCart, getTotalItems } = useCart();

  // Sample data (since we can't connect to Go backend in this environment)
  useEffect(() => {
    const sampleProducts: Product[] = [
      {
        id: 1,
        name: "iPhone 15 Pro",
        brand: "Apple",
        model: "A3108",
        price: 15999000,
        stock: 25,
        description: "Latest iPhone with titanium design and A17 Pro chip",
        image: "https://images.pexels.com/photos/788946/pexels-photo-788946.jpeg",
        created_at: "2024-01-01T00:00:00Z",
        updated_at: "2024-01-01T00:00:00Z"
      },
      {
        id: 2,
        name: "Samsung Galaxy S24 Ultra",
        brand: "Samsung",
        model: "SM-S928B",
        price: 18999000,
        stock: 30,
        description: "Premium Android flagship with S Pen and 200MP camera",
        image: "https://images.pexels.com/photos/1092644/pexels-photo-1092644.jpeg",
        created_at: "2024-01-01T00:00:00Z",
        updated_at: "2024-01-01T00:00:00Z"
      },
      {
        id: 3,
        name: "Xiaomi 14 Ultra",
        brand: "Xiaomi",
        model: "2405CPX3DG",
        price: 12999000,
        stock: 20,
        description: "Photography-focused flagship with Leica cameras",
        image: "https://images.pexels.com/photos/1275229/pexels-photo-1275229.jpeg",
        created_at: "2024-01-01T00:00:00Z",
        updated_at: "2024-01-01T00:00:00Z"
      },
      {
        id: 4,
        name: "Google Pixel 8 Pro",
        brand: "Google",
        model: "GC3VE",
        price: 13999000,
        stock: 15,
        description: "AI-powered smartphone with pure Android experience",
        image: "https://images.pexels.com/photos/1207583/pexels-photo-1207583.jpeg",
        created_at: "2024-01-01T00:00:00Z",
        updated_at: "2024-01-01T00:00:00Z"
      },
      {
        id: 5,
        name: "OnePlus 12",
        brand: "OnePlus",
        model: "CPH2573",
        price: 11999000,
        stock: 18,
        description: "Fast charging flagship with OxygenOS",
        image: "https://images.pexels.com/photos/1279107/pexels-photo-1279107.jpeg",
        created_at: "2024-01-01T00:00:00Z",
        updated_at: "2024-01-01T00:00:00Z"
      }
    ];
    
    setTimeout(() => {
      setProducts(sampleProducts);
      setFilteredProducts(sampleProducts);
      setLoading(false);
    }, 1000);
  }, []);

  // Filter products based on search query
  useEffect(() => {
    if (!searchQuery.trim()) {
      setFilteredProducts(products);
    } else {
      const filtered = products.filter(product =>
        product.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
        product.brand.toLowerCase().includes(searchQuery.toLowerCase()) ||
        product.model.toLowerCase().includes(searchQuery.toLowerCase())
      );
      setFilteredProducts(filtered);
    }
  }, [searchQuery, products]);

  if (loading) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600 mx-auto mb-4"></div>
          <p className="text-gray-600">Loading products...</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <Header
        cartItemsCount={getTotalItems()}
        onCartClick={() => setIsCartOpen(true)}
        searchQuery={searchQuery}
        onSearchChange={setSearchQuery}
      />

      {/* Hero Section */}
      <div className="bg-gradient-to-r from-blue-600 to-purple-600 text-white py-16">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h1 className="text-4xl md:text-6xl font-bold mb-4">
            Premium Phone Collection
          </h1>
          <p className="text-xl md:text-2xl mb-8 opacity-90">
            Discover the latest smartphones from top brands
          </p>
        </div>
      </div>

      {/* Products Section */}
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <div className="flex items-center justify-between mb-8">
          <h2 className="text-3xl font-bold text-gray-900">
            {searchQuery ? `Search Results (${filteredProducts.length})` : 'Featured Products'}
          </h2>
          {searchQuery && (
            <button
              onClick={() => setSearchQuery('')}
              className="text-blue-600 hover:text-blue-800 font-medium"
            >
              Clear Search
            </button>
          )}
        </div>
        
        {filteredProducts.length === 0 ? (
          <div className="text-center py-12">
            <p className="text-gray-500 text-lg">No products found</p>
            <p className="text-gray-400">Try adjusting your search terms</p>
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-8">
            {filteredProducts.map((product) => (
              <ProductCard
                key={product.id}
                product={product}
                onAddToCart={addToCart}
              />
            ))}
          </div>
        )}
      </div>

      {/* Cart */}
      <Cart
        isOpen={isCartOpen}
        onClose={() => setIsCartOpen(false)}
        cartItems={cartItems}
        onUpdateQuantity={updateQuantity}
        onRemoveItem={removeFromCart}
      />
    </div>
  );
}

export default App;
